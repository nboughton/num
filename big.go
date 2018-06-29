package num

import (
	"math/big"
	"strconv"
)

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

// BigFactorial returns n! using big.Int
func (n Int) BigFactorial() *big.Int {
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
