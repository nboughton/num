package num

import (
	"fmt"
	"math"
	"math/big"
)

// T represents IDs of numerical types
type T int

// Numerical "types" (triangle/square/pentagon etc)
const (
	EVEN T = iota
	ODD
	PRIME
	TRIANGLE
	SQUARE
	PENTAGONAL
	HEXAGONAL
	HEPTAGONAL
	OCTAGONAL
	FIBONACCI
	PANDIGITAL
)

// Is tests n for numerical attribute t
func (n Int) Is(t T) bool {
	switch t {
	case EVEN:
		return n%2 == 0

	case ODD:
		return n%2 != 0

	case PRIME:
		return big.NewInt(int64(n)).ProbablyPrime(5)

	case TRIANGLE:
		i, f := math.Modf((math.Sqrt(float64(8*n+1)) - 1) / 2)
		return f == 0 && i >= 1

	case SQUARE:
		i, f := math.Modf(math.Sqrt(float64(n)))
		return f == 0 && i >= 1

	case PENTAGONAL:
		i, f := math.Modf((math.Sqrt(float64(24*n+1)) + 1) / 6)
		return f == 0 && i >= 0

	case HEXAGONAL:
		i, f := math.Modf((math.Sqrt(float64(8*n+1)) + 1) / 4)
		return f == 0 && i >= 1

	case HEPTAGONAL:
		i, f := math.Modf((math.Sqrt(float64(40*n+9)) + 3) / 10)
		return f == 0 && i >= 1

	case OCTAGONAL:
		i, f := math.Modf((math.Sqrt(float64(3*n+1)) + 1) / 3)
		return f == 0 && i >= 1

	case PANDIGITAL:
		m := make(map[Int]bool)
		for _, i := range n.ToSet() {
			if _, ok := m[i]; !ok {
				m[i] = true
			} else {
				return false
			}
		}

		return true

	default:
		fmt.Println("Not implemented for num.Is()")
	}

	return false
}

// IsPyTriplet returns true if a < b < c and a^2 + b^2 = c^2
func IsPyTriplet(a, b, c Int) bool {
	if a < b && b < c && a*a+b*b == c*c {
		return true
	}
	return false
}

// IsUniqueCharString returns true if a string contains no duplicate
// characters
func IsUniqueCharString(s string) bool {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
		if m[v] > 1 {

		}
	}

	return true
}

// IsPalindrome returns true if b is a palindrome
func IsPalindrome(s string) bool {
	dst := make([]byte, len(s))
	for i := range s {
		dst[i] = s[len(s)-1-i]
	}

	return string(dst) == s
}
