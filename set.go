package num

import (
	"math"
	"math/big"
	//"math/big"
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
	var res Set

	for i := start; i <= end; i++ {
		res = append(res, i)
	}

	return res
}

// PrimeSieve returns a set of primes below n
func PrimeSieve(n Int) Set {
	var (
		sieve = make([]bool, n)
		res   Set
	)

	if n < 2 {
		return res
	}

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
	}

	for k := range m {
		res = append(res, k)
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
