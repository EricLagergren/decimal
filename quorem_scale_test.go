package decimal_test

import (
	"math/big"
	"testing"

	"github.com/ericlagergren/decimal"
)

func TestQuoRemScale(t *testing.T) {
	for _, tc := range []struct{ x, y string }{
		{"1.2", "1"}, {"12", "1.1"}, {"-1.2", "1"}, {"1.2", "-1"},
		{"123456789012345678901.2", "10"}, {"12", "100000000000000000000.1"},
		{"1E+30", "3"}, {"1", "3E+30"}, {"1.20", "0.10"}, {"0.00", "1.0"},
	} {
		t.Run(tc.x+"/"+tc.y, func(t *testing.T) {
			for _, alias := range []string{"fresh", "big-method", "reused", "quotient-x", "remainder-x", "quotient-y", "remainder-y"} {
				t.Run(alias, func(t *testing.T) {
					x, _ := new(decimal.Big).SetString(tc.x)
					y, _ := new(decimal.Big).SetString(tc.y)
					xr, yr := x.Rat(nil), y.Rat(nil)
					wantScale := x.Scale()
					if y.Scale() > wantScale {
						wantScale = y.Scale()
					}
					ratio := new(big.Rat).Quo(xr, yr)
					qi := new(big.Int).Quo(ratio.Num(), ratio.Denom())
					wantQ := new(big.Rat).SetInt(qi)
					wantR := new(big.Rat).Sub(xr, new(big.Rat).Mul(wantQ, yr))
					q, r := new(decimal.Big), new(decimal.Big)
					switch alias {
					case "reused":
						q.SetString("0.001")
						r.SetString("0.00001")
					case "quotient-x":
						q = x
					case "remainder-x":
						r = x
					case "quotient-y":
						q = y
					case "remainder-y":
						r = y
					}
					ctx := decimal.Context{Precision: 100}
					if alias == "big-method" {
						q.Context = ctx
						q.QuoRem(x, y, r)
					} else {
						ctx.QuoRem(q, x, y, r)
					}
					if q.Scale() != 0 || r.Scale() != wantScale {
						t.Errorf("scales = (%d, %d), want (0, %d)", q.Scale(), r.Scale(), wantScale)
					}
					if q.Rat(nil).Cmp(wantQ) != 0 || r.Rat(nil).Cmp(wantR) != 0 {
						t.Fatalf("QuoRem = (%s, %s), want (%s, %s)", q, r, wantQ, wantR)
					}
				})
			}
		})
	}
}
