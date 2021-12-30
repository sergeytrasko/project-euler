package main

import (
	"fmt"
	"math/big"
)

func main() {
	limit := int64(5e15)
	fmt.Println(limit)
	cnt := 0
	for n := int64(1); ; n++ {
		p := n*n + (n+1)*(n+1)
		if n%100_000 == 0 {
			fmt.Printf("n=%d, p=%d, found primes=%d\n", n, p, cnt)
		}
		if p > limit {
			break
		}
		if big.NewInt(p).ProbablyPrime(0) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
