package main

import "fmt"

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

func numberOfSolutions(exponents []int) int64 {
	r := int64(1)
	for i := 0; i < len(exponents); i++ {
		r *= int64(1 + 2*exponents[i])
	}
	return (r + 1) / 2
}

func number(primes []int, exponents []int) int64 {
	r := int64(1)
	for i := 0; i < len(exponents); i++ {
		for j := 0; j < exponents[i]; j++ {
			r *= int64(primes[i])
		}
	}
	return r
}

func solve(primes []int, exponents []int, target int64, best *int64) int64 {
	nn := number(primes, exponents)
	if nn > *best {
		return *best
	}
	if n := numberOfSolutions(exponents); n > target {
		*best = nn
		// fmt.Printf("%d - %d - %d\n", nn, exponents, n)
		return nn
	}
	if n := solve(primes, append(exponents, 1), target, best); n < *best {
		*best = n
		return n
	}
	if len(exponents) == 1 || exponents[len(exponents)-2] > exponents[len(exponents)-1] {
		exponents[len(exponents)-1]++
		if n := solve(primes, exponents, target, best); n < *best {
			*best = n
			return n
		}
	}
	return *best
}

func main() {
	primes := sieve(100)[:15]
	best := int64(1e17)
	n := solve(primes, []int{0}, 4_000_000, &best)
	fmt.Println(n)
}
