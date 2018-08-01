package num

import (
	"fmt"
	"math"
	"math/big"
)

// Seq returns a channel of numbers for type t
func Seq(t T) chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		switch t {
		case EVEN:
			for i := Int(2); i < Int(math.MaxInt64); i += 2 {
				c <- i
			}

		case ODD:
			for i := Int(1); i < Int(math.MaxInt64); i += 2 {
				c <- i
			}

		case PRIME:
			c <- 2

			for i := Int(3); i < Int(math.MaxInt64); i += 2 {
				if i.Is(PRIME) {
					c <- i
				}
			}

		case TRIANGLE:
			for i := Int(1); i < Int(math.MaxInt64); i++ {
				c <- i * (i + 1) / 2
			}

		case SQUARE:
			for i := Int(1); i < Int(math.MaxInt64); i++ {
				c <- i * i
			}

		case PENTAGONAL:
			for i := Int(1); i < Int(math.MaxInt64); i++ {
				c <- i * (3*i - 1) / 2
			}

		case HEXAGONAL:
			for i := Int(1); i < Int(math.MaxInt64); i++ {
				c <- i * (2*i - 1)
			}

		case HEPTAGONAL:
			for i := Int(1); i < Int(math.MaxInt64); i++ {
				c <- i * (5*i - 3) / 2
			}

		case OCTAGONAL:
			for i := Int(1); i < Int(math.MaxInt64); i++ {
				c <- i * (3*i - 2)
			}

		case FIBONACCI:
			a, b := Int(0), Int(1)
			for {
				a = a + b
				a, b = b, a
				c <- a
			}

		default:
			fmt.Println("Not implemented for num.Seq()")
		}
	}()

	return c
}

// PellLucas streams n iterations of the Pell/Pell-Lucas sequence. These can
// be used as approximations for the continued fraction of the square root of 2
func PellLucas(n Int) chan [2]*big.Int {
	a, b, r := big.NewInt(0), big.NewInt(1), make(chan [2]*big.Int)

	go func() {
		defer close(r)

		for i := Int(0); i < n; i++ {
			c, _ := big.NewInt(0).SetString(a.String(), 10)

			a.Add(a, big.NewInt(0).Mul(big.NewInt(2), b))
			a, b = b, a

			c.Add(c, a)

			r <- [2]*big.Int{big.NewInt(0).Set(c), big.NewInt(0).Set(a)}
		}
	}()

	return r
}

// BigFib returns a channel of the Fibonacci sequence using big Ints.
func BigFib() chan *big.Int {
	c := make(chan *big.Int, 1)

	go func() {
		defer close(c)

		a, b := big.NewInt(0), big.NewInt(1)

		for {
			a.Add(a, b)
			a, b = b, a
			c <- big.NewInt(0).Set(a)
		}
	}()

	return c
}

// Collatz returns the Collatz sequence starting at n
func Collatz(n Int) Set {
	seq := Set{n}

	for {
		n = seq[len(seq)-1]
		if n <= 1 {
			break
		}

		if n%2 == 0 {
			seq = append(seq, n/2)
		} else {
			seq = append(seq, 3*n+1)
		}
	}

	return seq
}
