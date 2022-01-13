package main

import "fmt"

func sieve(max int64) []int64 {
	n := make([]bool, max)
	primes := []int64{}
	for i := int64(2); i < max; i++ {
		if !n[i] {
			primes = append(primes, i)
			for j := int64(1); ; j++ {
				if i*j >= max {
					break
				}
				n[i*j] = true
			}
		}
	}
	return primes
}

func mebius(n int64) []int64 {
	primes := sieve(n + 1)
	m := make([]int64, n+1)
	for i := 0; i < len(m); i++ {
		m[i] = 1
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for pos := p; pos < int64(len(m)); pos += p {
			m[pos] = -m[pos]
		}
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for pos := p * p; pos < int64(len(m)); pos += p * p {
			m[pos] = 0
		}
	}
	return m
}

func phiSum(n int64) int64 {
	m := mebius(n)
	sum := int64(0)
	for k := int64(1); k <= n; k++ {
		f := n / k
		sum += int64(m[k] * (f * f))
	}
	sum++
	sum /= 2
	return sum
}

func H(n int64) int64 {
	return 3*n*(n+1) - 6*phiSum(n)
}

func main() {
	// fmt.Println(H(5))
	// fmt.Println(H(10))
	// fmt.Println(H(1000))
	// fmt.Println(H(1_000_000))
	fmt.Println(H(100_000_000))
}
