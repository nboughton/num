package num

import (
	"math"
	"math/big"
)

// IsPrime returns true if n is prime
func (n Int) IsPrime() bool {
	return big.NewInt(int64(n)).ProbablyPrime(5)
}

// IsTriangular returns true if n is a triangular number
func (n Int) IsTriangular() bool {
	i, f := math.Modf((math.Sqrt(float64(8*n+1)) - 1) / 2)
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// IsSquare returns true if n is a square number
func (n Int) IsSquare() bool {
	i, f := math.Modf(math.Sqrt(float64(n)))
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// IsPentagonal returns true if n is a pentagonal number
func (n Int) IsPentagonal() bool {
	i, f := math.Modf((math.Sqrt(float64(24*n+1)) + 1) / 6)
	if f == 0 && i >= 0 {
		return true
	}
	return false
}

// IsHexagonal returns true if n is a hexagonal number
func (n Int) IsHexagonal() bool {
	i, f := math.Modf((math.Sqrt(float64(8*n+1)) + 1) / 4)
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// IsHeptagonal returns true if n is a heptagonal number
func (n Int) IsHeptagonal() bool {
	i, f := math.Modf((math.Sqrt(float64(40*n+9)) + 3) / 10)
	if f == 0 && i >= 1 {
		return true
	}
	return false
}

// IsOctagonal returns true if n is an octagonal number
func (n Int) IsOctagonal() bool {
	i, f := math.Modf((math.Sqrt(float64(3*n+1)) + 1) / 3)
	if f == 0 && i >= 1 {
		return true
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
			return false
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
