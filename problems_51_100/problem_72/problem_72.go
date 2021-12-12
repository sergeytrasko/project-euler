package main

import (
	"fmt"
	"math"
)

const max = 1_000_000

func totient(n int, primes []int) int {
	nn := n
	factors := []int{}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if n%p == 0 {
			factors = append(factors, p)
			for n%p == 0 {
				n /= p
			}
		}
		if n == 1 {
			break
		}
	}
	t := nn
	for i := 0; i < len(factors); i++ {
		p := factors[i]
		t = t * (p - 1) / p
	}
	return t
}

func sieve(max int) []int {
	n := make([]bool, max)
	primes := []int{}
	for i := 2; i < max; i++ {
		if !n[i] {
			primes = append(primes, i)
			for j := 1; ; j++ {
				if i*j >= max {
					break
				}
				n[i*j] = true
			}
		}
	}
	return primes
}

func naive() int64 {
	//https://en.wikipedia.org/wiki/Farey_sequence#Next_term
	f := int64(250)
	n := int64(max)
	a, b := int64(0), int64(1)
	c, d := int64(1), int64(n)
	cnt := int64(0)
	for {
		// fmt.Printf("%d/%d\n", c, d)
		p := int64(math.Floor(float64(n+b)/float64(d)))*c - a
		q := int64(math.Floor(float64(n+b)/float64(d)))*d - b
		a, b, c, d = c, d, p, q
		cnt++
		if a == 1 && b == f {
			break
		}
	}
	return cnt*f + 1
}

func main() {
	primes := sieve(max)
	// https://en.wikipedia.org/wiki/Farey_sequence#Sequence_length_and_index_of_a_fraction
	// F(n) = F(n-1) + totient(n), F(1) = 2
	f := 1
	for i := 1; i <= 1_000_000; i++ {
		f = f + totient(i, primes)
	}
	fmt.Println(f - 2)
}
