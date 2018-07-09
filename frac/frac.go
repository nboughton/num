package frac

import (
	"fmt"
	"github.com/nboughton/num"
)

// Frac represents a fraction by its Num(erator), Den(ominator) and decimal Val(ue)
type Frac struct {
	Num, Den num.Int
	Val      float64
}

// NewFrac returns a new Frac
func NewFrac(n, d num.Int) *Frac {
	return &Frac{
		Num: n,
		Den: d,
		Val: float64(n) / float64(d),
	}
}

// GCD returns the Greatest Common Denominator (aka Highest Common Factor) of the numerator and denominator of Frac f
func (f *Frac) GCD() num.Int {
	return num.Set{f.Num, f.Den}.GCD()
}

func (f *Frac) String() string {
	return fmt.Sprintf("%v/%v", f.Num, f.Den)
}
