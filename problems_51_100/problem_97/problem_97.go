package main

import (
	"fmt"
	"math/big"
)

func naive() string {
	p := new(big.Int).Add(
		new(big.Int).Mul(
			big.NewInt(28433),
			new(big.Int).Exp(big.NewInt(2), big.NewInt(7830457), nil),
		),
		big.NewInt(1),
	)
	s := p.String()
	return s[len(s)-10:]
}

func modular() int64 {
	mul := int64(28433)
	exp := int64(7830457)
	modulo := int64(10_000_000_000)
	// p = (mul * 2^exp + 1) mod 10_000_000_000
	x := int64(1)
	for i := int64(0); i < exp; i++ {
		x = (x * 2) % int64(modulo)
	}
	x = (mul*x + 1) % int64(modulo)
	return x
}

func main() {
	fmt.Println(naive())
	fmt.Println(modular())
}
