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

	//fmt.Println(fr)
	d := math.Pow10(decimalPlaces(fr))
	n := d * fr

	return NewFrac(Int(n), Int(d)).Reduce(), nil
}

// CF emits the Continued Fraction integer terms of f using Euclid's algorithm. This only works for
// rational numbers as irrational numbers eventually become inaccurate.
func (f *Frac) CF() Set {
	res := Set{}

	var cf func(n, d Int)
	cf = func(n, d Int) {
		if d == 0 {
			return
		}

		r := n % d
		res = append(res, n/d)

		//fmt.Printf("%d; %d/%d\n", n/d, r, d)
		cf(d, r)
	}

	cf(f.Num, f.Den)

	return res
}

// CfSqrtX attempts to calculate the convergents of Sqrt(x) such that they
// satisfy Pell's Equation. It is currently unfinished.
func CfSqrtX(x float64) {
	sqrt := math.Sqrt(x)
	if _, frac := math.Modf(sqrt); frac == 0 {
		return
	}

	sf := sqrt / x
	fmt.Println(sf)
}

// decimalPlaces returns the number of decimal places in f
func decimalPlaces(f float64) int {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	arr := strings.Split(str, ".")
	return len(arr[1])
}
