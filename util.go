package decimal

import (
	"math/big"

	"github.com/ericlagergren/decimal/internal/arith"
	"github.com/ericlagergren/decimal/internal/arith/checked"
	"github.com/ericlagergren/decimal/internal/arith/pow"
	"github.com/ericlagergren/decimal/internal/c"
)

func (z *Big) norm() *Big {
	if z.isInflated() && arith.IsUint64(&z.unscaled) {
		if v := z.unscaled.Uint64(); v != c.Inflated {
			z.compact = v
		}
	}
	return z
}

func (z *Big) test() *Big {
	adj := z.adjusted()
	if emax := z.emax(); adj > emax {
		if z.IsFinite() {
			z.exp = emax
			// TODO(eric): mandatory clamping?
		}
		z.Context.Conditions |= Clamped
	} else if adj < z.emin() {
		tiny := z.etiny()
		if z.IsFinite() {
			if z.exp < tiny {
				z.exp = tiny
				z.Context.Conditions |= Clamped
			}
			return z
		}

		z.Context.Conditions |= Subnormal
		if z.exp < tiny {
			z.Round(tiny - z.exp)
			z.exp = tiny
		}
	}
	return z
}

// alias returns z if z != x, otherwise a newly-allocated big.Int.
func alias(z, x *big.Int) *big.Int {
	if z != x {
		// We have to check the first element of their internal slices since
		// Big doesn't store a pointer to a big.Int.
		zb, xb := z.Bits(), x.Bits()
		if cap(zb) > 0 && cap(xb) > 0 && &zb[0:cap(zb)][cap(zb)-1] != &xb[0:cap(xb)][cap(xb)-1] {
			return z
		}
	}
	return new(big.Int)
}

func precision(z *Big) (p int) {
	p = z.Context.Precision
	if p > 0 {
		return p
	}
	if p == 0 {
		z.Context.Precision = DefaultPrecision
	} else {
		z.Context.Conditions |= InvalidContext
	}
	return DefaultPrecision
}

func mode(x *Big) OperatingMode { return x.Context.OperatingMode }

// copybits can be useful when we want to allocate a big.Int without calling
// new or big.Int.Set. For example:
//
//   var x big.Int
//   if foo {
//       x.SetBits(copybits(y.Bits()))
//   }
//   ...
//
func copybits(x []big.Word) []big.Word {
	z := make([]big.Word, len(x))
	copy(z, x)
	return z
}

// cmpNorm compares x and y in the range [0.1, 0.999...] and returns true if x
// > y.
func cmpNorm(x uint64, xs int, y uint64, ys int) (ok bool) {
	goodx, goody := true, true

	// xs, ys > 0, so no overflow
	if diff := xs - ys; diff != 0 {
		if diff < 0 {
			x, goodx = checked.MulPow10(x, -uint64(diff))
		} else {
			y, goody = checked.MulPow10(y, uint64(diff))
		}
	}
	if goodx {
		if goody {
			return arith.Cmp(x, y) > 0
		}
		return false
	}
	return true
}

// cmpNormBig compares x and y in the range [0.1, 0.999...] and returns true if
// x > y. It uses z as backing storage, provided it does not alias x or y.
func cmpNormBig(z, x *big.Int, xs int, y *big.Int, ys int) (ok bool) {
	if xs != ys {
		z = alias(alias(z, x), y)
		if xs < ys {
			x = checked.MulBigPow10(z, x, uint64(ys-xs))
		} else {
			y = checked.MulBigPow10(z, y, uint64(xs-ys))
		}
	}
	// x and y are non-negative
	return x.Cmp(y) > 0
}

// scalex adjusts x by scale. If scale > 0, x = x * 10^scale, otherwise
// x = x / 10^-scale.
func scalex(x uint64, scale int) (sx uint64, ok bool) {
	if scale > 0 {
		sx, ok = checked.MulPow10(x, uint64(scale))
		if !ok {
			return 0, false
		}
		return sx, true
	}
	p, ok := pow.Ten(uint64(-scale))
	if !ok {
		return 0, false
	}
	return x / p, true
}

// bigScalex sets z to the big.Int equivalient of scalex.
func bigScalex(z, x *big.Int, scale int) *big.Int {
	if scale > 0 {
		return checked.MulBigPow10(z, x, uint64(scale))
	}
	return z.Quo(x, pow.BigTen(uint64(-scale)))
}
