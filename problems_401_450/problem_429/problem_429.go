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

func countFactors(n int, p int) int {
	cnt := 0
	x := p
	for x <= n {
		cnt += n / x
		x *= p
	}
	return cnt
}

func powMod(a, b int, mod int64) int64 {
	return new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), big.NewInt(mod)).Int64()
}

const n = 100_000_000
const mod = 1_000_000_009

func main() {
	primes := sieve(n + 1)
	factors := make([]int, len(primes))
	for i := 0; i < len(primes); i++ {
		factors[i] = countFactors(n, primes[i])
	}
	sum := int64(1)
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		e := factors[i]
		//b(p^e, 2) = p^(2e) + 1
		sum = (sum * (powMod(p, 2*e, mod) + 1)) % mod
	}

	fmt.Println(sum)
}
