package main

import (
	"fmt"
)

func residue(pn int64, n int64) int64 {
	if n%2 == 0 {
		return 2
	} else {
		return (2 * n * pn) % (pn * pn)
	}
}

var primes = []int64{2}

func isPrime(p int64) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := int64(len(primes))
	for p := int64(3); ; p += 2 {
		if isPrime(p) {
			primes = append(primes, p)
			n++
			r := residue(p, n)
			if r > 1e10 {
				break
			}
		}
	}
	fmt.Println(n)
}
