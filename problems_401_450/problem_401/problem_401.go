package main

import (
	"fmt"
	"math"
	"math/big"
)

func SIGMA2(n int64, mod int64) int64 {
	f := func(x int64) int64 {
		r := new(big.Int).Div(
			new(big.Int).Mul(
				big.NewInt(x),
				new(big.Int).Mul(
					big.NewInt(x+1),
					big.NewInt(2*x+1),
				),
			),
			big.NewInt(6),
		)
		return r.Mod(r, big.NewInt(mod)).Int64()
		// return (x * (x + 1) * (2*x + 1) / 6)
	}
	s := int64(math.Sqrt(float64(n)))
	sum := int64(0)
	for k := int64(1); k <= s; k++ {
		x := n / k
		sum = (sum + f(x) + (k*x%mod)*(k%mod)) % mod
	}
	r := sum - s*f(s)
	for r < 0 {
		r += mod
	}
	return r % mod
}

const mod = 1e9

func main() {
	// for n := int64(1); n <= 6; n++ {
	// 	fmt.Printf("SIGMA2(%d)=%d\n", n, SIGMA2(n, mod))
	// }
	fmt.Println(SIGMA2(1e15, mod))
}
