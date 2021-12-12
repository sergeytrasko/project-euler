package main

import (
	"fmt"
	"sort"
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

const N = 1_000_000_000

func main() {
	primes := sieve(100)
	candidates := []int{1}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		m := p
		x := []int{}
		for m <= N {
			x = append(x, m)
			m *= p
		}
		y := append([]int{}, candidates...)
		for j := 0; j < len(x); j++ {
			for k := 0; k < len(y); k++ {
				z := x[j] * y[k]
				if z > N {
					break
				}
				candidates = append(candidates, z)
			}
		}
		sort.Ints(candidates)
	}
	fmt.Println(len(candidates))
}
