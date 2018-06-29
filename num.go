package num

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Int is the default value used in Sets and Matrices in this package
type Int int64

var abc = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var atoi = make(map[byte]Int)

func init() {
	// populate atoi map for Abc func
	for i, v := range abc {
		atoi[v] = Int(i + 1)
	}
}

// AbcToInt returns the a = 1, b = 2 etc score for string s
func AbcToInt(s string) Int {
	s = strings.ToUpper(s)

	score := Int(0)
	for _, v := range []byte(s) {
		score += atoi[v]
	}

	return score
}

// String returns n as a base 10 string and satisfies the stringer interface
func (n Int) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// ToSet returns n as a set of its digits
func (n Int) ToSet() Set {
	var (
		res Set
		s   = strconv.FormatInt(int64(n), 10)
	)

	for _, v := range s {
		i, _ := strconv.ParseInt(string(v), 10, 64)
		res = append(res, Int(i))
	}

	return res
}

// Divisors returns a Set of divisors of n
func (n Int) Divisors() Set {
	var (
		div Set
		lim = Int(math.Sqrt(float64(n)))
	)

	for i := Int(1); i <= lim; i++ {
		if n%i == 0 {
			div = append(div, i)
			if i*i != n {
				div = append(div, n/i)
			}
		}
	}

	return div.Dedupe()
}

// PrimeFactors returns the Set of prime factors of n
func (n Int) PrimeFactors() Set {
	pf := Set{}

	for _, v := range n.Divisors() {
		if v.IsPrime() {
			pf = append(pf, v)
		}
	}

	return pf
}

// Totient returns the result of Eulers Totient or Phi function of value n
func (n Int) Totient() Int {
	pF := n.PrimeFactors()

	ans := n

	for _, prime := range pF {
		ans = ans * (prime - 1) / prime
	}

	return ans
}

// BIG INTS
// Some calculcations are going to produce numbers that overflow Int(int64)

// BigToSet returns a Set from the digits of a big.Int
func BigToSet(n *big.Int) Set {
	var res Set

	for _, v := range n.String() {
		i, _ := strconv.ParseInt(string(v), 10, 64)
		res = append(res, Int(i))
	}

	return res
}

// Factorial returns n! using big.Int as values increase exponentially
func (n Int) Factorial() *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	s := []string{}
	for i := n; i > 0; i-- {
		s = append(s, strconv.FormatInt(int64(i), 10))
	}

	return BigProduct(s)
}

// BigSum takes an array of strings representing big ints and returns
// a big Int value of the sum
func BigSum(n []string) *big.Int {
	a := big.NewInt(0)
	a.SetString(n[0], 10)

	for i := 1; i < len(n); i++ {
		b := big.NewInt(0)
		b.SetString(n[i], 10)

		a.Add(a, b)
	}
	return a
}

// BigProduct takes an array of strings representing numbers and
// and returns a big Int containing their Product
func BigProduct(n []string) *big.Int {
	a := big.NewInt(0)
	a.SetString(n[0], 10)

	for i := 1; i < len(n); i++ {
		b := big.NewInt(0)
		b.SetString(n[i], 10)

		a.Mul(a, b)
	}
	return a
}

// BigPow returns x^y as a big Int
func BigPow(x, y Int) *big.Int {
	n, m := big.NewInt(int64(x)), big.NewInt(int64(x))
	for i := Int(2); i <= y; i++ {
		n.Mul(n, m)
	}
	return n
}
