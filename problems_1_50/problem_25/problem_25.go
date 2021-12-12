package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(1), big.NewInt(1)
	n := 1
	for len(a.String()) < 1000 {
		c := a.Add(a, b)
		a, b = b, c
		n++
	}
	fmt.Println(n)
}
