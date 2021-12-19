package main

import (
	"fmt"
	"math/big"
)

func a(n int) int {
	k := 1
	for x := 1; x != 0; x, k = (x*10+1)%n, k+1 {
	}
	return k
}

func main() {
	n := 2
	cnt := 0
	sum := 0
	for cnt < 25 {
		if n%5 != 0 && n%2 != 0 {
			k := a(n)
			if (n-1)%k == 0 && !big.NewInt(int64(n)).ProbablyPrime(10) {
				sum += n
				cnt++
			}
		}
		n++
	}
	fmt.Println(sum)
}
