package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	powers := []string{}
	maxA, maxB := 100, 100
	for a := 2; a <= maxA; a++ {
		for b := 2; b <= maxB; b++ {
			x := big.NewInt(int64(a))
			power := x.Exp(x, big.NewInt(int64(b)), nil)
			powers = append(powers, power.String())
		}
	}
	sort.Strings(powers)
	// fmt.Println(powers)
	cntDistinct := 0
	prev := ""
	for i := 0; i < len(powers); i++ {
		if powers[i] != prev {
			cntDistinct++
			prev = powers[i]
		}
	}
	fmt.Println(cntDistinct)
}
