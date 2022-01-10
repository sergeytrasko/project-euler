package main

import (
	"fmt"
	"math/big"
)

const mod = 1e16
const div = 250
const setSize = 250250

func main() {
	s := [setSize + 1]int64{}
	for i := int64(1); i <= setSize; i++ {
		s[i] = new(big.Int).Exp(big.NewInt(i), big.NewInt(i), big.NewInt(div)).Int64()
	}
	ways := [div]int64{}
	ways[0] = 1
	for index := 1; index <= setSize; index++ {
		n := s[index]
		prev := [div]int64{}
		for i := int64(0); i < div; i++ {
			prev[i] = ways[i]
		}
		for i := int64(0); i < div; i++ {
			ways[(i+n)%div] = (ways[(i+n)%div] + prev[i]) % mod
		}
	}
	fmt.Println(ways[0] - 1)
}
