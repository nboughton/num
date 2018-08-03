package num

import (
	"math"
	"math/big"
	"sort"
	"strconv"
)

// Set is a slice of Int
type Set []Int

func (s Set) Len() int           { return len(s) }
func (s Set) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Set) Less(i, j int) bool { return s[i] < s[j] }

// Range returns a set of start .. end inclusive
func Range(start, end Int) Set {
	res := make(Set, end-start+1)

	for i, n := Int(0), start; i < end-start+1; i, n = i+1, n+1 {
		res[i] = n
	}

	return res
}

// PrimeSieve returns a set of primes below n
func PrimeSieve(n Int) Set {
	var res Set

	if n < 2 {
		return res
	}

	sieve := make([]bool, n)
	for i := Int(2); i <= Int(math.Sqrt(float64(n))); i++ {
		if !sieve[i] {
			for j := i * i; j < n; j += i {
				sieve[j] = true
			}
		}
	}

	for i, val := range sieve {
		if !val {
			res = append(res, Int(i))
		}
	}

	return res[2:]
}

// Contains returns whether or not n exists in set s
func (s Set) Contains(n Int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}

	return false
}

/*
// GRS returns the Greatest Repeating Subset in a set
func (s Set) GRS() Set {
	var res Set

	if len(s) < 10 { // a tad arbitrary
		return s
	}

	i, j, k, h, t1, t2 := 0, 1, 2, 3, Set{}, Set{}
	for ; i < j; i++ {
		for j = i + 1; j < k; j++ {
			t1 = Set(s[i : j+1])

			for k = j + 1; k < h; k++ {
				for h = k + 1; h < len(s); h++ {
					t2 = Set(s[k : h+1])

					if t1.Cmp(t2) && len(t1) > len(res) {
						res = t1
					}
				}
			}
		}
	}

	return res
}
*/

// Cmp compares the contents of two Sets and returns true if they are of
// equal length and contain the same values in the same order
func (s Set) Cmp(t Set) bool {
	if len(s) != len(t) {
		return false
	}

	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}

// GCD returns the Greatest Common Divisor of all items in Set s
func (s Set) GCD() Int {
	if len(s) < 2 {
		return Int(0)
	}

	var gcd func(x, y Int) Int
	gcd = func(x, y Int) Int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}

	res := gcd(s[0], s[1])

	for i := 2; i < len(s); i++ {
		res = gcd(res, s[i])
	}

	return res
}

// Sum returns the sum total of the set
func (s Set) Sum() Int {
	var t Int

	for _, n := range s {
		t += n
	}

	return t
}

// Product returns the product of the set
func (s Set) Product() Int {
	t := s[0]

	for _, n := range s[1:] {
		t *= n
	}

	return t
}

// ToInt returns an Int by contacetenating its elements
func (s Set) ToInt() Int {
	var b []byte

	for _, v := range s {
		b = strconv.AppendInt(b, int64(v), 10)
	}

	n, _ := strconv.ParseInt(string(b), 10, 64)
	return Int(n)
}

// ToBigInt returns a big int as the concatenated elements of Set s
func (s Set) ToBigInt() *big.Int {
	var t string

	for _, n := range s {
		t += n.String()
	}

	o, _ := big.NewInt(0).SetString(t, 10)
	return o
}

// Dedupe returns a sorted set with only unique values
func (s Set) Dedupe() Set {
	var (
		m   = make(map[Int]int)
		res Set
	)

	for _, n := range s {
		m[n]++
		if m[n] == 1 {
			res = append(res, n)
		}
	}

	sort.Sort(res)
	return res
}

// Combinations returns all ln length combinations of set s
func (s Set) Combinations(ln int) chan Set {
	c := make(chan Set)

	go func() {
		defer close(c)

		pool := s
		n := len(pool)

		indices := make(Set, ln)
		for i := range indices {
			indices[i] = Int(i)
		}

		result := make(Set, ln)
		for i, el := range indices {
			result[i] = pool[el]
		}

		c <- result

		for {
			i := ln - 1
			for i >= 0 && indices[i] == Int(i+n-ln) {
				i--
			}

			if i < 0 {
				break
			}

			indices[i]++
			for j := i + 1; j < ln; j++ {
				indices[j] = indices[j-1] + 1
			}

			result := make(Set, ln)
			for i = 0; i < len(indices); i++ {
				result[i] = pool[indices[i]]
			}

			c <- result
		}
	}()

	return c
}

// Permutations returns ln length permutations of Set s in lexicographic order
func (s Set) Permutations(ln int) chan Set {
	c := make(chan Set)
	if ln > len(s) || ln == 0 {
		close(c)
		return c
	}

	go func() {
		defer close(c)

		pool := s
		n := len(pool)

		indices := make(Set, n)
		for i := range indices {
			indices[i] = Int(i)
		}

		cycles := make(Set, ln)
		for i := range cycles {
			cycles[i] = Int(n - i)
		}

		result := make(Set, ln)
		for i, el := range indices[:ln] {
			result[i] = pool[el]
		}
		c <- result

		for n > 0 {
			i := ln - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					index := indices[i]
					for j := i; j < n-1; j++ {
						indices[j] = indices[j+1]
					}
					indices[n-1] = index
					cycles[i] = Int(n - i)
				} else {
					j := int(cycles[i])
					indices[i], indices[n-j] = indices[n-j], indices[i]

					result := make(Set, ln)
					for k := 0; k < ln; k++ {
						result[k] = pool[indices[k]]
					}

					c <- result
					break
				}
			}
			if i < 0 {
				break
			}
		}
	}()

	return c
}

// Convergents treats s as the [n0; n1, n2, n3...] representation of a continued fraction
// and returns a stream of its convergents. If the sequence is recurring then the Channel
// stops when either h or k exceeds math.MaxInt64
func (s Set) Convergents(recurring bool) chan Set {
	c := make(chan Set, 1)

	go func() {
		defer close(c)

		if s == nil || len(s) < 2 {
			return
		}

		a, h, k := Set{0, 0}, Set{0, 1}, Set{1, 0}
		a = append(a, s...)

		for n := 2; n < len(a); n++ {
			hn := a[n]*h[n-1] + h[n-2]
			kn := a[n]*k[n-1] + k[n-2]

			h, k = append(h, hn), append(k, kn)

			// If hn or kn overflow math.MaxInt64 they will become negative.
			if hn < 0 || kn < 0 {
				return
			}

			c <- Set{hn, kn}

			// Grow a as necessary by appending the recurring portion of the
			// continued fraction.
			if recurring && n == len(a)-1 && len(a) < math.MaxInt32-len(s) {
				a = append(a, s[1:]...)
			}
		}
	}()

	return c
}
