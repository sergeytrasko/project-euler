package main

import (
	"fmt"
	"math"
	"math/big"
)

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

func generateAdmissible(current int64, usedFactors int, index int, primes []int, N int64, a *[]int64) {
	if index == len(primes) {
		return
	}
	if current > N {
		return
	}
	if usedFactors > 1 {
		*a = append(*a, current)
	}
	p := int64(primes[index])
	c := current * p
	for c < N {
		generateAdmissible(c, usedFactors+1, index+1, primes, N, a)
		c *= p
	}
}

func admissible(N int64) []int64 {
	a := []int64{}
	for p := int64(2); p < N; p *= 2 {
		a = append(a, p)
	}
	primes := sieve(int(math.Sqrt(float64(N) + 1)))
	generateAdmissible(1, 0, 0, primes, N, &a)
	return a
}

func main() {
	a := admissible(1e9)
	res := map[int64]bool{}
	for i := 0; i < len(a); i++ {
		n := a[i]
		for m := int64(2); ; m++ {
			if big.NewInt(n + m).ProbablyPrime(1) {
				res[m] = true
				break
			}
		}
	}
	sum := int64(0)
	for m := range res {
		sum += m
	}
	fmt.Println(sum)
}
