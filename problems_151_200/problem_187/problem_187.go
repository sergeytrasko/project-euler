package main

import "fmt"

var primes = []int{}

func isPrime(p int) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	return true
}

const max = 100_000_000

func main() {
	for p := 2; p < max/2; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	fmt.Printf("There are %d primes below %d\n", len(primes), max)
	// fmt.Println(primes)
	cnt := 0
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for j := i; j < len(primes); j++ {
			q := primes[j]
			if p*q > max {
				break
			}
			cnt++
		}
	}
	fmt.Println(cnt)
}
