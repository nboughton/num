package num

import (
	"math"
	"strconv"
	"strings"
)

// Int is the default value used in Sets and Matrices in this package
type Int int64

func init() {
	// populate atoi map for Abc func
	for i, v := range abc {
		atoi[v] = Int(i + 1)
	}
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
		if v.Is(PRIME) {
			pf = append(pf, v)
		}
	}

	return pf
}

// Factorial returns n!
func (n Int) Factorial() Int {
	if n == 0 {
		return Int(1)
	}

	var res Set
	for i := n; i > 0; i-- {
		res = append(res, i)
	}

	return res.Product()
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

// Rotations returns a sequence of rotations of n.
// I.e Rotations(123) = 123 -> 312 -> 231
func (n Int) Rotations() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		s := n.ToSet()
		for i := 0; i < len(s); i++ {
			rt := append(Set{s[len(s)-1]}, s[:len(s)-1]...)
			c <- rt.ToInt()
			s = rt
		}
	}()

	return c
}

// Truncate returns a channel of Int slices that contain the
// truncation sequence of n from the left and the right simultaneously.
// I.e Truncate(123) = [123, 123] -> [23, 12] -> [3, 1]
func (n Int) Truncate() chan Set {
	c := make(chan Set, 1)

	go func() {
		defer close(c)

		s := n.ToSet()
		for i := range s {
			c <- Set{s[i:].ToInt(), s[:len(s)-i].ToInt()}
		}
	}()

	return c
}

// Partition returns the number of partitions of n with m parts. See https://en.wikipedia.org/wiki/Partition_(8number_theory)
func Partition(n, m Int) Int {
	if m < 2 {
		return m
	}

	if n < m {
		return 0
	}

	var memo Matrix
	for i := Int(0); i <= n-m; i++ {
		memo = append(memo, make(Set, m))
	}

	var p func(n, m Int) Int
	p = func(n, m Int) Int {
		if n <= m+1 {
			return 1
		}

		if memo[n-m][m-2] != 0 {
			return memo[n-m][m-2]
		}

		max := n / m
		if m == 2 {
			return max
		}

		count := Int(0)
		for ; max > 0; max, n = max-1, n-m {
			memo[n-m][m-3] = p(n-1, m-1)
			count += memo[n-m][m-3]
		}

		return count
	}

	return p(n, m)
}

var abc = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var atoi = make(map[byte]Int)

// AbcToInt returns the a = 1, b = 2 etc score for string s
func AbcToInt(s string) Int {
	s = strings.ToUpper(s)

	score := Int(0)
	for _, v := range []byte(s) {
		score += atoi[v]
	}

	return score
}
