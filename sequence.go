package num

import (
	"math"
	"math/big"
)

// SeqPrimes returns a continuous channel of Int Primes
func SeqPrimes() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		c <- 2

		for i := Int(3); i < Int(math.MaxInt64); i += 2 {
			if i.IsPrime() {
				c <- i
			}
		}
	}()

	return c
}

// SeqPrimesBetween returns a channel with all primes between start and finish
func SeqPrimesBetween(start, finish Int) chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for _, i := range PrimeSieve(finish) {
			if i > start {
				c <- i
			}
		}
	}()

	return c
}

// SeqPrimesFrom returns a channel of primes from start
func SeqPrimesFrom(start Int) chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		if start == 2 {
			c <- start
		}

		if start%2 == 0 {
			start++
		}

		for i := start; i < Int(math.MaxInt64); i += 2 {
			if i.IsPrime() {
				c <- i
			}
		}
	}()

	return c
}

// SeqNPrimesFrom returns n conescutive primes starting from x
func SeqNPrimesFrom(start, n Int) chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		if start == 2 {
			c <- start
		}

		if start%2 == 0 {
			start++
		}

		count := Int(0)
		for i := start; count < n; i += 2 {
			if i.IsPrime() {
				c <- i
				count++
			}
		}
	}()

	return c
}

// SeqPellLucas streams n iterations of the Pell/Pell-Lucas sequence. These can
// Seqbe used as approximations for the continued fraction of the square root of 2
func SeqPellLucas(n Int) chan [2]*big.Int {
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

// SeqFibonacci returns a channel of the Fibonacci sequence using big Ints.
// SeqBig ints are used because of the exponential growth of Fibonacci numbers.
func SeqFibonacci() chan *big.Int {
	c := make(chan *big.Int, 1)

	go func() {
		defer close(c)

		a, b := big.NewInt(0), big.NewInt(1)

		for true {
			a.Add(a, b)
			a, b = b, a
			c <- big.NewInt(0).Set(a)
		}
	}()

	return c
}

// SeqCollatz returns the Collatz sequence starting at n
func SeqCollatz(n Int) Set {
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

// SeqBigInts returns a continuous stream of big Ints integers from 1
func SeqBigInts() chan *big.Int {
	c := make(chan *big.Int, 1)

	go func() {
		defer close(c)

		for i := big.NewInt(1); true; i.Add(i, big.NewInt(1)) {
			c <- i
		}
	}()

	return c
}

// SeqInts returns a continuous channel of integers from 1
func SeqInts() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i
		}
	}()

	return c
}

// SeqEvens returns a continuous channel of even numbers from 2
func SeqEvens() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(2); i < Int(math.MaxInt64); i += 2 {
			c <- i
		}
	}()

	return c
}

// SeqOdds returns a continuous channel of odd numbers from 1
func SeqOdds() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i += 2 {
			c <- i
		}
	}()

	return c
}

// SeqTriangles returns a channel of the triangle number sequence
func SeqTriangles() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i * (i + 1) / 2
		}
	}()

	return c
}

// SeqSquares returns a channel of square numbers in sequence
func SeqSquares() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i * i
		}
	}()

	return c
}

// SeqPentagonals returns a channel of the pentagonal number sequence
func SeqPentagonals() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i * (3*i - 1) / 2
		}
	}()

	return c
}

// SeqHexagonals returns a channel of the hexagonal number sequence
func SeqHexagonals() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i * (2*i - 1)
		}
	}()

	return c
}

// SeqHeptagonals returns a channel of the heptagonal number sequence
func SeqHeptagonals() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i * (5*i - 3) / 2
		}
	}()

	return c
}

// SeqOctagonals returns a channel of the octagonal number sequence
func SeqOctagonals() chan Int {
	c := make(chan Int, 1)

	go func() {
		defer close(c)

		for i := Int(1); i < Int(math.MaxInt64); i++ {
			c <- i * (3*i - 2)
		}
	}()

	return c
}

// SeqRotations returns a sequence of rotations of n.
// I.e Rotations(123) = 123 -> 312 -> 231
func SeqRotations(n Int) chan Int {
	rts := make(chan Int, 1)

	go func() {
		defer close(rts)

		s := []byte(big.NewInt(int64(n)).String())

		for i := 0; i < len(s); i++ {
			t := []byte{s[len(s)-1]}
			t = append(t, s[:len(s)-1]...)
			m, _ := big.NewInt(0).SetString(string(t), 10)
			rts <- Int(m.Int64())
			s = t
		}
	}()

	return rts
}

// SeqTruncate returns a channel of Int slices that contain the
// truncation sequence of n from the left and the right simultaneously.
// I.e Truncate(123) = [123, 123] -> [23, 12] -> [3, 1]
func SeqTruncate(n Int) chan Set {
	c := make(chan Set, 1)

	go func() {
		d := []byte(big.NewInt(int64(n)).String())

		for i := range d {
			l, _ := big.NewInt(0).SetString(string(d[i:]), 10)
			r, _ := big.NewInt(0).SetString(string(d[:len(d)-i]), 10)
			c <- Set{Int(l.Int64()), Int(r.Int64())}
		}

		close(c)
	}()

	return c
}
