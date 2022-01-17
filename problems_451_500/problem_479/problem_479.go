package main

import (
	"fmt"
	"math"
	"math/big"
)

func explore(n int) int {
	sum := 0
	for p := 1; p <= n; p++ {
		s := 0
		for k := 1; k <= n; k++ {
			s += int(math.Pow(float64(1-k*k), float64(p)))
		}
		sum += s
	}
	return sum
}

func solve(n int64) int64 {
	m := big.NewInt(1_000_000_007)
	sum := big.NewInt(0)
	for k := int64(1); k <= n; k++ {
		pow := new(big.Int).Exp(big.NewInt(1-k*k), big.NewInt(n+1), m)
		pow = new(big.Int).Sub(big.NewInt(1), pow)
		pow = pow.Mul(pow, new(big.Int).Exp(big.NewInt(k*k), big.NewInt(-1), m))
		pow = pow.Mod(pow, m)
		sum = sum.Add(sum, pow)
		sum = sum.Mod(sum, m)
	}
	sum = sum.Sub(sum, big.NewInt(n))
	return sum.Mod(sum, m).Int64()
}

func main() {
	// fmt.Println(explore(4))
	// fmt.Println(solve(4))
	fmt.Println(solve(1e6))
}
