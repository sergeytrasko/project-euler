package main

import (
	"fmt"
	"strconv"
)

var primes = []int{}

func isPrime(p int) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	return true
}

func isCyclePrime(p int) bool {
	s := strconv.Itoa(p)
	for i := 0; i < len(s); i++ {
		s = s[1:] + s[0:1]
		// fmt.Printf("%d, %d, %s\n", p, i, s)
		pp, _ := strconv.Atoi(s)
		if !isPrime(pp) {
			return false
		}
	}
	return true
}

func main() {
	max := 1000000
	cnt := 0
	for p := 2; p < max; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	for p := 0; p < len(primes); p++ {
		if isCyclePrime(primes[p]) {
			// fmt.Println(primes[p])
			cnt++
		}
	}
	fmt.Println(cnt)
}
