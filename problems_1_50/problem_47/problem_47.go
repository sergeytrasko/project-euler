package main

import "fmt"

var primes = []int{}

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeFactors(n int) int {
	cnt := 0
	for i := 0; i < len(primes) && primes[i] < n; i++ {
		if n%primes[i] == 0 {
			cnt++
		}
	}
	return cnt
}

const n = 4

func main() {
	for i := 2; i < 1_000_000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	k := 2
	for {
		ok := true
		for i := 0; i < n; i++ {
			if countPrimeFactors(k+i) != n {
				ok = false
				break
			}
		}
		if ok {
			break
		}
		k++
	}
	fmt.Println(k)
}
