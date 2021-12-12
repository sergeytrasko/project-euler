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

func fillPrimes(max int) {
	primes = append(primes, 2)
	p := 3
	for p <= max {
		if isPrime(p) {
			primes = append(primes, p)
		}
		p += 2
	}
}

func checkIfPrime(p int, pos int) bool {
	for pos < len(primes) {
		if primes[pos] == p {
			return true
		}
		if primes[pos] > p {
			return false
		}
		pos++
	}
	return false
}

func sum(c []int) int {
	s := 0
	for _, r := range c {
		s += r
	}
	return s
}

func main() {
	max := 1000000
	fillPrimes(20 * max)
	fmt.Println(len(primes))

	bestSum := 0
	bestLen := 0
	for i := 0; i < len(primes); i++ {
		sum := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			if sum > max {
				break
			}
			if checkIfPrime(sum, j) && j-i+1 > bestLen {
				bestLen = j - i + 1
				bestSum = sum
				fmt.Printf("%d from %d gives sum %d\n", j-i+1, primes[i], sum)
			}
		}
	}
	fmt.Println(bestSum)
}
