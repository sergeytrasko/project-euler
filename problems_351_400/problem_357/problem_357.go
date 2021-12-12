package main

import (
	"fmt"
	"sort"
)

const max = 100_000_000

var primes = []int{}

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

func isPrime(p int, primes []int) bool {
	i := sort.SearchInts(primes, p)
	return i < len(primes) && primes[i] == p
}

/*
n is even, because:
if n = 2k + 1 then 2k+1 is a divisor of n. Therefore 2k+1 + (2k+1)/(2k+1) = 2k+1 + 1 = 2(k+1) is not prime
*/
func main() {
	primes = sieve(max * 2)
	sum := 1
	for n := 2; n <= max; n += 4 {
		if !isPrime(n+1, primes) {
			continue
		}
		if !isPrime(2+n/2, primes) {
			continue
		}
		ok := true
		for d := 3; d*d <= n; d++ {
			if n%d == 0 {
				if !isPrime(d+n/d, primes) {
					ok = false
					break
				}
			}
		}
		if ok {
			// fmt.Println(n)
			sum += n
		}
	}
	fmt.Println(sum)
	// generatePrimes()
}
