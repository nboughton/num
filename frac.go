package num

import (
	"fmt"
	"math"
)

// Frac represents a fraction by its Num(erator), Den(ominator) and decimal Dec
type Frac struct {
	Num, Den Int
	Dec      float64
}

// NewFrac returns a new Frac
func NewFrac(n, d Int) *Frac {
	return &Frac{
		Num: n,
		Den: d,
		Dec: float64(n) / float64(d),
	}
}

// GCD returns the Greatest Common Denominator (aka Highest Common Factor) of the numerator and denominator of Frac f
func (f *Frac) GCD() Int {
	return Set{f.Num, f.Den}.GCD()
}

func (f *Frac) String() string {
	return fmt.Sprintf("%v/%v", f.Num, f.Den)
}

// Simplify simplifies f and returns it
func (f *Frac) Simplify() *Frac {
	gcd := f.GCD()

	f.Num /= gcd
	f.Den /= gcd

	return f
}

// DecToFrac converts a decimal to a reduced fraction
func DecToFrac(d float64) (*Frac, error) {
	if d > 1 {
		return nil, fmt.Errorf("d must be < 1")
	}

	var numerator Int
	if _, err := fmt.Sscanf(fmt.Sprint(d), "0.%d", &numerator); err != nil {
		return nil, err
	}

	denominator := Int(math.Pow10(len(numerator.ToSet())))
	return NewFrac(numerator, denominator).Simplify(), nil
}

// WORK IN PROGRESS
// Continued emits the continued fraction represenations of f
/*
func (f *Frac) Continued() chan *Frac {
	c := make(chan *Frac)

	go func() {
		defer close(c)

		var cf func(f *Frac)
		cf = func(f *Frac) {
			i, frac := math.Modf(f.Dec)
			fmt.Println(i, frac)
			if frac == 0 {
				return
			}

			r := NewFrac(Int(i), Int(frac))
			fmt.Println(r.Simplify())
			//r := NewFrac(f.Den, f.Num)
			c <- r
			cf(r)
		}

		cf(f)
	}()

	return c
}
*/
