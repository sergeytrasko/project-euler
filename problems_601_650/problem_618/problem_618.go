package main

import "fmt"

func getLimit(n int) int64 {
	a, b := int64(1), int64(1)
	for k := 3; k <= n; k++ {
		a, b = b, a+b
	}
	return b
}

func sieve(max int64) []int64 {
	n := make([]bool, max)
	primes := []int64{}
	for i := int64(2); i < max; i++ {
		if !n[i] {
			primes = append(primes, i)
			for j := int64(1); ; j++ {
				if i*j >= max {
					break
				}
				n[i*j] = true
			}
		}
	}
	return primes
}

func main() {
	N := getLimit(24)
	m := int64(1e9)
	primes := sieve(N)

	b := make([][]int64, N+1)
	for n := 0; n < len(b); n++ {
		b[n] = make([]int64, len(primes))
	}
	for i := 0; i < len(primes); i++ {
		b[0][i] = 1
	}
	for n := int64(1); n <= N; n++ {
		for i := 0; i < len(primes); i++ {
			if i > 0 {
				b[n][i] += b[n][i-1] % m
			}
			if primes[i] <= n {
				b[n][i] += b[n-primes[i]][i] * primes[i] % m
			}
		}
	}
	sum := int64(b[1][len(primes)-1])
	x, y := int64(1), int64(1)
	for y != N {
		x, y = y, x+y
		sum += b[y][len(primes)-1]
	}
	fmt.Println(sum % m)
}
