package main

import "fmt"

// const max = 10_000_000
const max = 10_000_000

var primes = []int{}

func isPrime(p int) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	return true
}

func generatePrimes() {
	primes = append(primes, 2)
	for p := 3; p < max; p += 2 {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	fmt.Printf("Found %d primes under %d\n", len(primes), max)
}

func powerOf(n, k int) int {
	p := 0
	for n%k == 0 {
		n /= k
		p++
	}
	return p
}

func main() {
	generatePrimes()
	divisors := [max]int{}
	for i := 1; i < max; i++ {
		divisors[i] = 1
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		divisors[p] *= 2
		k := p
		for k+p < max {
			k += p
			divisors[k] *= (powerOf(k, p) + 1)
		}
	}
	cnt := 0
	for i := 1; i < max; i++ {
		if divisors[i-1] == divisors[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
