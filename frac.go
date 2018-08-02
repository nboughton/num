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
	Int  Int
	Frac *big.Rat
}

// ContinuedFraction "should" emit the continued fraction represenations of f as their Integer and simplified Fractional parts.
// Bearing in mind the difficulties inherent in representing fractional values in base 10 floating point numbers. See:
// https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems
func ContinuedFraction(f *big.Rat) chan CF {
	c := make(chan CF, 1)

	go func() {
		defer close(c)

		var cf func(r *big.Rat)
		cf = func(r *big.Rat) {
			f, _ := r.Float64()
			i, _ := math.Modf(f)

			// s becomes the simplified fractional part
			s := new(big.Rat).Sub(r, big.NewRat(int64(i), 1))

			// Return Step values
			c <- CF{Int: Int(i), Frac: new(big.Rat).Set(s)}

			// Stop at 0/1
			if s.IsInt() {
				return
			}

			cf(s.Inv(s))
		}

		// Start with a new copy of f to prevent mutation
		cf(new(big.Rat).Set(f))
	}()

	return c
}
