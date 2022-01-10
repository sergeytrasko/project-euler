package main

import (
	"fmt"
	"math"
)

func mebius(n int) []int {
	sieve := func(max int) []int {
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
	primes := sieve(n + 1)
	m := make([]int, n+1)
	for i := 0; i < len(m); i++ {
		m[i] = 1
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for pos := p; pos < len(m); pos += p {
			m[pos] = -m[pos]
		}
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for pos := p * p; pos < len(m); pos += p * p {
			m[pos] = 0
		}
	}
	return m
}

func solveWithMebius(N int64) {
	limit := int(math.Sqrt(float64(N)))
	m := mebius(limit)
	cnt := int64(0)
	for i := int64(1); i <= int64(limit); i++ {
		cnt += int64(m[i]) * ((N - 1) / (i * i))
	}
	fmt.Println(cnt)
}

func main() {
	N := math.Pow(2, 50)
	limit := int64(math.Sqrt(N))
	c := make([]int64, limit+2)
	for i := limit; i > 1; i-- {
		c[i] = int64(N) / (i * i)
		for j := int64(2); j <= limit/i; j++ {
			c[i] -= c[i*j]
		}
	}
	cnt := int64(0)
	for i := int64(1); i <= limit; i++ {
		cnt += c[i]
	}
	fmt.Println(int64(N) - cnt)
	solveWithMebius(int64(N))
}
