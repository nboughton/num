package num

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Frac represents a (possibly mixed) fraction by its Int(eger), Num(erator), Den(ominator)
// and Flt (float64) values
type Frac struct {
	Num, Den, Int Int
	Flt           float64
}

// NewFrac returns a new Frac
func NewFrac(n, d Int) *Frac {
	f := &Frac{
		Num: n,
		Den: d,
		Flt: float64(n) / float64(d),
	}

	i, _ := math.Modf(f.Flt)
	f.Int = Int(i)

	return f
}

// GCD returns the Greatest Common Denominator (aka Highest Common Factor) of the numerator and denominator of Frac f
func (f *Frac) GCD() Int {
	return Set{f.Num, f.Den}.GCD()
}

func (f *Frac) String() string {
	return fmt.Sprintf("%v/%v", f.Num, f.Den)
}

// Reduce simplifies f and returns it
func (f *Frac) Reduce() *Frac {
	gcd := f.GCD()

	f.Num /= gcd
	f.Den /= gcd

	return f
}

// Inverse returns the inverse of f as a new Frac
func (f *Frac) Inverse() *Frac {
	return NewFrac(f.Den, f.Num)
}

// FloatToFrac converts a decimal to a reduced fraction
func FloatToFrac(f float64) (*Frac, error) {
	_, fr := math.Modf(f)

	d := math.Pow10(decimalPlaces(fr))
	n := d * fr

	return NewFrac(Int(n), Int(d)).Reduce(), nil
}

// decimalPlaces returns the number of decimal places in f
func decimalPlaces(f float64) int {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	arr := strings.Split(str, ".")
	return len(arr[1])
}

// Continued "should" emit the continued fraction represenations of f and probably doesn't work due to the limitations of
// floating point number accuracy. See: https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems
func (f *Frac) Continued() chan *Frac {
	c := make(chan *Frac)

	go func() {
		defer close(c)

		var cf func(f *Frac)
		cf = func(f *Frac) {
			fr := f.Flt - float64(f.Int)
			if fr == 0 {
				return
			}

			s, err := FloatToFrac(fr)
			if err != nil {
				return
			}

			i := s.Inverse()
			c <- i
			cf(i)
		}

		cf(f)
	}()

	return c
}
