package main

import (
	"fmt"
	"math/big"
)

const n = 1000

func main() {
	mod := big.NewInt(int64(10000000000))
	sum := big.NewInt(0)
	for i := 1; i <= n; i++ {
		x := big.NewInt(int64(i))
		x = x.Exp(x, x, mod)
		sum = sum.Add(sum, x)
	}
	fmt.Println(sum.Mod(sum, mod))
}
