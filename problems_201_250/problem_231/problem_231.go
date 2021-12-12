package main

import (
	"fmt"
)

const max = 20_000_000

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

func primeDivisors(n int, primes []int) map[int]int {
	d := make(map[int]int)
	for i := 0; primes[i]*primes[i] <= n; i++ {
		p := primes[i]
		for n%p == 0 {
			n /= p
			d[p]++
		}
	}
	d[n]++
	return d
}

//https://en.wikipedia.org/wiki/Binomial_coefficient#Multiplicative_formula
func cPrimeFactorsSum(n, k int, primes []int) int {
	if k > n-k {
		k = n - k
	}
	x, y := make(map[int]int), make(map[int]int)
	for i := 1; i <= k; i++ {
		for p, v := range primeDivisors(n+1-i, primes) {
			x[p] += v
		}
		for p, v := range primeDivisors(i, primes) {
			y[p] += v
		}
	}
	// fmt.Println(x)
	// fmt.Println(y)
	sum := 0
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		sum += p * (x[p] - y[p])
	}
	return sum
}

func main() {
	primes := sieve(max)
	fmt.Println(len(primes))
	// fmt.Println(cPrimeFactorsSum(10, 3, primes))
	fmt.Println(cPrimeFactorsSum(20_000_000, 15_000_000, primes))
}
