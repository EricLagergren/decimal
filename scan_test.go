package decimal

import (
	"math/rand"
	"strconv"
	"testing"
)

func randString() string {
	const chars = "0123456789"
	n := rand.Intn(5000)
	if n == 0 {
		n = 16
	}
	p := make([]byte, n)
	for i := range p {
		p[i] = chars[rand.Intn(len(chars))]
	}
	switch rand.Intn(3) {
	case 0:
		p = append(p, 'e', '+')
		p = strconv.AppendInt(p, int64(rand.Intn(350)), 10)
	case 1:
		p = append(p, 'e', '-')
		p = strconv.AppendInt(p, int64(rand.Intn(350)), 10)
	case 2:
		if lp := len(p); lp > 0 {
			p[rand.Intn(lp)] = '.'
		}
	}
	return string(p)
}

var benchInput = func() (a [4096]struct {
	b Big
	s string
}) {
	for i := range a {
		a[i].s = randString()
	}
	return a
}()

var globOk bool

func BenchmarkBig_SetString(b *testing.B) {
	var ok bool
	for i := 0; i < b.N; i++ {
		m := &benchInput[i%len(benchInput)]
		_, ok = m.b.SetString(m.s)
	}
	globOk = ok
}

func TestUnmarshalTextTwice(t *testing.T) {
	x := &Big{}

	x.UnmarshalText([]byte(`1`))
	x.UnmarshalText([]byte(`2`))
	if x.String() != "2" {
		t.Errorf("wanted: %q, got %q", "2", x.String())
	}
}

func TestSetStringTruncatedNumeric(t *testing.T) {
	for _, s := range []string{"1e", "1E", "1e+", "1e-", "-1e", "1.2e", "."} {
		x, ok := new(Big).SetString(s)
		if !ok {
			t.Fatalf("%s: SetString returned false", s)
		}
		if !x.IsNaN(0) {
			t.Fatalf("%s = %s, want NaN", s, x)
		}
		if x.Context.Conditions&ConversionSyntax == 0 {
			t.Fatalf("%s: missing ConversionSyntax", s)
		}
	}

	kept := []struct{ in, want string }{
		{"1e2", "1E+2"},
		{"1e+2", "1E+2"},
		{"1e-1", "0.1"},
		{"1.", "1"},
		{".5", "0.5"},
		{"10.00", "10.00"},
	}
	for _, tc := range kept {
		x, ok := new(Big).SetString(tc.in)
		if !ok || x.String() != tc.want {
			t.Fatalf("%s = %s ok=%v, want %s", tc.in, x, ok, tc.want)
		}
	}
}

func TestSetStringTwice(t *testing.T) {
	x := &Big{}

	x.SetString("1")
	x.SetString("2")
	if x.String() != "2" {
		t.Errorf("wanted: %q, got %q", "2", x.String())
	}
}
