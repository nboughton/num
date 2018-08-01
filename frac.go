package num

import (
	"fmt"
	"math"
	"math/big"
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

	fmt.Println(fr)
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

// CF represents both the Integer part and the simplified fractional part of each step in a continued fraction
type CF struct {
	I    int64
	Frac *big.Rat
}

// Continued "should" emit the continued fraction represenations of f and probably doesn't work due to the limitations of
// floating point number accuracy See: https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems
// This code works for the example given in the above wikipedia page under Calculating continued fraction representations.
func (f *Frac) Continued() chan CF {
	c := make(chan CF)

	go func() {
		defer close(c)

		var cf func(r *big.Rat)
		cf = func(r *big.Rat) {
			f, _ := r.Float64()
			i, _ := math.Modf(f)

			// n becomes the simplified fractional part
			n := new(big.Rat).Sub(r, big.NewRat(int64(i), 1))
			c <- CF{I: int64(i), Frac: n}

			if n.Num().Cmp(big.NewInt(0)) == 0 {
				return
			}

			cf(n.Inv(n))
		}

		cf(big.NewRat(int64(f.Num), int64(f.Den)))
	}()

	return c
}
