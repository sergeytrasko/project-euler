package main

import (
	"fmt"
	"sort"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

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

func mebius(n int) []int {
	primes := sieve(n + 1)
	m := make([]int, n+1)
	for i := 0; i < len(m); i++ {
		m[i] = 1
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for pos := p; pos < len(m); pos += p {
			m[pos] = -m[pos]
		}
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for pos := p * p; pos < len(m); pos += p * p {
			m[pos] = 0
		}
	}
	return m
}

func isPrime(p int, primes []int) bool {
	i := sort.SearchInts(primes, p)
	return i < len(primes) && primes[i] == p
}

func permutations(min int, max int, result *[][]int) {
	var p func(cur []int, available []int, result *[][]int)
	p = func(cur []int, available []int, result *[][]int) {
		if len(available) == 0 {
			(*result) = append((*result), cur)
		}
		for i := 0; i < len(available); i++ {
			tmp := append([]int{}, available...)
			p(append(cur, available[i]), append(tmp[:i], tmp[i+1:]...), result)
		}
	}
	avail := []int{}
	for i := min; i <= max; i++ {
		avail = append(avail, i)
	}
	p([]int{}, avail, result)
}

func main() {
	primes := sieve(100_000_000)
	// primes := sieve(100)
	fmt.Println(len(primes))
	fmt.Println(isPrime(2, primes))
	fmt.Println(isPrime(4, primes))
	fmt.Println(isPrime(31, primes))
	fmt.Println(isPrime(21, primes))
	fmt.Println(isPrime(100, primes))
	fmt.Println(isPrime(97, primes))
}
