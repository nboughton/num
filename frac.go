package num

import (
	"fmt"
	"math"
)

// Frac represents a fraction by its Num(erator), Den(ominator) and decimal Dec
type Frac struct {
	Num, Den, Int Int
	Dec           float64
}

// NewFrac returns a new Frac
func NewFrac(i, n, d Int) *Frac {
	return &Frac{
		Int: i,
		Num: n,
		Den: d,
		Dec: float64(i) + float64(n)/float64(d),
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
func DecToFrac(f float64) (*Frac, error) {
	var i, n Int
	if _, err := fmt.Sscanf(fmt.Sprint(f), "%d.%d", &i, &n); err != nil {
		return nil, err
	}

	d := Int(math.Pow10(len(n.ToSet())))
	return NewFrac(i, n, d).Simplify(), nil
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
