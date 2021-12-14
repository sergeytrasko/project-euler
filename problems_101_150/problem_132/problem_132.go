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

const limit = 1e9

func main() {
	primes := sieve(1_000_000)
	cnt := 0
	sum := 0
	for i := 0; i < len(primes) && cnt < 40; i++ {
		p := primes[i]
		x := new(big.Int).Exp(big.NewInt(10), big.NewInt(limit), big.NewInt(int64(9*p)))
		if x.Cmp(big.NewInt(1)) == 0 {
			cnt++
			sum += p
		}
	}
	fmt.Println(sum)
}
