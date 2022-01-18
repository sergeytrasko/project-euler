package main

import (
	"fmt"
	"math/big"
)

const mod = 1234567891011

var memo = map[int64]int64{}

func f(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	if memo[n] > 0 {
		return memo[n]
	}
	res := int64(0)
	if n%2 == 1 {
		fn := big.NewInt(f(n / 2))
		fn1 := big.NewInt(2 * f(n/2+1))
		fn1 = fn1.Sub(fn1, fn)
		fn1 = fn1.Mul(fn1, fn)
		fn1 = fn1.Mod(fn1, big.NewInt(mod))
		res = fn1.Int64()
	} else {
		fn := big.NewInt(f((n + 1) / 2))
		fn1 := big.NewInt(f((n - 1) / 2))
		fn = fn.Mul(fn, fn)
		fn1 = fn1.Mul(fn1, fn1)
		fn = fn.Add(fn, fn1)
		fn = fn.Mod(fn, big.NewInt(mod))
		res = fn.Int64()
	}
	memo[n] = res
	return res
}

func main() {
	N := int64(1e14)
	x := N
	cnt := 0
	sum := int64(0)
	for cnt < 100_000 {
		if big.NewInt(x).ProbablyPrime(1) {
			sum = (sum + f(x-1)) % mod
			cnt++
		}
		x++
	}
	fmt.Println(sum)
}
