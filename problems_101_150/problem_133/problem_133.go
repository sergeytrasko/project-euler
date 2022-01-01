package main

import (
	"fmt"
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

func isFactor(p int) bool {
	m := int64(9 * p)
	phi := big.NewInt(int64(6 * (p - 1)))
	reminders := map[int64]bool{}
	for n := int64(0); n <= m; n++ {
		x := new(big.Int).Exp(big.NewInt(10), big.NewInt(n), phi)
		x = x.Exp(big.NewInt(10), x, big.NewInt(m))
		r := x.Int64()
		if r == 1 {
			return true
		}
		if reminders[r] {
			return false
		}
		reminders[r] = true
	}
	return false
}

func main() {
	primes := sieve(100_000)
	sum := 0
	for _, p := range primes {
		if !isFactor(p) {
			sum += p
		}
	}
	fmt.Println(sum)
}
