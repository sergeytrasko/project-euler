package main

import (
	"fmt"
	"math"
)

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

func bruteForce() {
	primes := sieve(2_000)
	sum := 0
	for n := 5; n <= 1_000; n++ {
		s := math.Sqrt(float64(n))
		lps, ups := 0, 0
		for i := 0; i < len(primes); i++ {
			if float64(primes[i]) <= s && s < float64(primes[i+1]) {
				lps = primes[i]
				break
			}
		}
		for i := 0; i < len(primes); i++ {
			if float64(primes[i]) < s && s <= float64(primes[i+1]) {
				ups = primes[i+1]
				break
			}
		}
		if (n%lps == 0 && n%ups != 0) || (n%lps != 0 && n%ups == 0) {
			// fmt.Printf("n=%d, lps^2=%d, ups^2=%d\n", n, lps*lps, ups*ups)
			sum += n
		}
	}
	fmt.Println(sum)
}

func main() {
	// bruteForce()
	max := int64(999966663333)
	primes := sieve(int(2 * math.Sqrt(float64(max))))
	sum := int64(0)
	for i := 0; i < len(primes)-1; i++ {
		p := int64(primes[i])
		q := int64(primes[i+1])
		for j := p*p + p; j < q*q; j += p {
			if j <= max {
				sum += j
			}
		}
		for j := q*q - q; j > p*p; j -= q {
			if j <= max {
				sum += j
			}
		}
		if p*q < max {
			sum -= 2 * p * q
		}
		if q*q > max {
			break
		}
	}
	fmt.Println(sum)
}
