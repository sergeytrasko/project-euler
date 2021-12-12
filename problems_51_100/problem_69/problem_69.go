package main

import "fmt"

const max = 1_000_000

var primes = []int{}

func isPrime(p int) bool {
	for i := 2; i*i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func phi(n int) int {
	primeDivisors := []int{}
	p := n
	for i := 0; i < len(primes) && primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			for p%primes[i] == 0 {
				p = p / primes[i]
			}
			primeDivisors = append(primeDivisors, primes[i])
		}
	}
	r := float64(n)
	for i := 0; i < len(primeDivisors); i++ {
		r *= float64(1.0 - 1.0/float64(primeDivisors[i]))
	}

	return int(r)
}

func main() {
	for p := 2; p < max; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	best := float64(0)
	bestN := 0
	for n := 2; n < max; n += 2 {
		ratio := float64(n) / float64(phi(n))
		if ratio > best {
			bestN = n
			best = ratio
			fmt.Printf("%d/phi(%d)=%f\n", n, n, ratio)
		}
	}
	fmt.Println(bestN)
}
