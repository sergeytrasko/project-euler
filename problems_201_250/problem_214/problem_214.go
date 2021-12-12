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

func totient(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	t := n
	factors := []int{}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if p*p > n {
			break
		}
		if n%p == 0 {
			factors = append(factors, p)
			for n%p == 0 {
				n /= p
			}
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	for _, r := range factors {
		t = t * (r - 1) / r
	}
	return t
}

func chain(n int, primes []int, chainLengthCache map[int]int) int {
	if n == 1 {
		return 1
	}
	if chainLengthCache[n] != 0 {
		return chainLengthCache[n]
	}
	t := totient(n, primes)
	l := 1 + chain(t, primes, chainLengthCache)
	chainLengthCache[n] = l
	return l
}

func main() {
	primes := sieve(40_000_000)
	// fmt.Printf("Found %d primes\n", len(primes))
	cache := make(map[int]int)
	sum := 0
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		l := 1 + chain(p-1, primes, cache)
		cache[p] = l
		if l == 25 {
			// fmt.Printf("chain(%d)=%d\n", p, l)
			sum += p
		}
	}
	fmt.Println(sum)
}
