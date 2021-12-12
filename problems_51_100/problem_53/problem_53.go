package main

import (
	"fmt"
	"math/big"
)

const maxN = 101

var million = big.NewInt(1_000_000)

func main() {
	p := [maxN][maxN]*big.Int{}
	for i := 0; i < maxN; i++ {
		p[i][0] = big.NewInt(1)
		p[i][i] = big.NewInt(1)
	}
	cnt := 0
	for n := 1; n < maxN; n++ {
		for k := 1; k < n; k++ {
			p[n][k] = big.NewInt(0).Set(p[n-1][k])
			p[n][k] = p[n][k].Add(p[n][k], p[n-1][k-1])
			if p[n][k].Cmp(million) > 0 {
				cnt++
				// fmt.Println(p[n][k])
			}
		}
	}
	fmt.Println(cnt)
}
