package decimal

import "io"

type byteScanner struct {
	r io.ByteScanner
	s string
	b []byte
	i int
}

func (s *byteScanner) ReadByte() (byte, error) {
	if s.r != nil {
		return s.r.ReadByte()
	}
	if s.b != nil {
		if s.i >= len(s.b) {
			return 0, io.EOF
		}
		c := s.b[s.i]
		s.i++
		return c, nil
	}
	if s.i >= len(s.s) {
		return 0, io.EOF
	}
	c := s.s[s.i]
	s.i++
	return c, nil
}

func (s *byteScanner) UnreadByte() error {
	if s.r != nil {
		return s.r.UnreadByte()
	}
	if s.i <= 0 {
		return io.EOF
	}
	s.i--
	return nil
}