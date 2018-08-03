package num

import (
	"math"
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
func BigFactorial(n Int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	b := big.NewInt(int64(n))
	for i := n - 1; i > 0; i-- {
		b.Mul(b, big.NewInt(int64(i)))
	}

	return b
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

// BigPartition returns the number of parts of n of size m using *big.Int
func BigPartition(n, m Int) *big.Int {
	if m < 2 {
		return big.NewInt(int64(m))
	}

	if n < m {
		return big.NewInt(0)
	}

	var memo [][]*big.Int
	for i := Int(0); i <= n-m; i++ {
		set := make([]*big.Int, m)
		for j := range set {
			set[j] = big.NewInt(0)
		}
		memo = append(memo, set)
	}

	var p func(n, m Int) *big.Int
	p = func(n, m Int) *big.Int {
		if n <= m+1 {
			return big.NewInt(1)
		}

		if memo[n-m][m-2].Cmp(big.NewInt(0)) != 0 {
			return memo[n-m][m-2]
		}

		max := n / m
		if m == 2 {
			return big.NewInt(int64(max))
		}

		count := big.NewInt(0)
		for ; max > 0; max, n = max-1, n-m {
			memo[n-m][m-3] = p(n-1, m-1)
			count.Add(count, memo[n-m][m-3])
		}

		return count
	}

	return p(n, m)
}

// BigConvergents treats s as the [n0; n1, n2, n3...] representation of a continued fraction
// and returns a stream of its convergents.
func BigConvergents(s Set, recurring bool) chan [2]*big.Int {
	c := make(chan [2]*big.Int, 1)

	go func() {
		defer close(c)

		if s == nil || len(s) < 2 {
			return
		}

		a := []*big.Int{big.NewInt(0), big.NewInt(0)}
		h := []*big.Int{big.NewInt(0), big.NewInt(1)}
		k := []*big.Int{big.NewInt(1), big.NewInt(0)}

		for _, n := range s {
			a = append(a, big.NewInt(int64(n)))
		}

		for n := 2; n < len(a); n++ {
			//hn := a[n]*h[n-1] + h[n-2]
			//kn := a[n]*k[n-1] + k[n-2]

			hn := new(big.Int).Add(
				new(big.Int).Mul(a[n], h[n-1]),
				h[n-2],
			)

			kn := new(big.Int).Add(
				new(big.Int).Mul(a[n], k[n-1]),
				k[n-2],
			)

			h, k = append(h, hn), append(k, kn)

			c <- [2]*big.Int{new(big.Int).Set(hn), new(big.Int).Set(kn)}

			// Grow a as necessary by appending the recurring portion of the
			// continued fraction.
			if recurring && n == len(a)-1 && len(a) < math.MaxInt32-len(s) {
				for _, i := range s[1:] {
					a = append(a, big.NewInt(int64(i)))
				}
			}
		}
	}()

	return c
}
