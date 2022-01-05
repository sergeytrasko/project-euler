package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("1504170715041707", 10)
	m, _ := new(big.Int).SetString("4503599627370517", 10)
	eulercoins := []*big.Int{a}

	for n := int64(1); ; n++ {
		nn := big.NewInt(int64(n))
		nn = nn.Mul(nn, a)
		e := nn.Mod(nn, m)
		last := eulercoins[len(eulercoins)-1]
		if e.Cmp(last) < 0 {
			fmt.Printf("%d - %d, dif=%d\n", n, e, new(big.Int).Sub(new(big.Int).Set(last), e))
			eulercoins = append(eulercoins, e)
			if e.Cmp(big.NewInt(n)) < 0 {
				break
			}
		}
	}
	limit := eulercoins[len(eulercoins)-1].Int64()
	aInv := new(big.Int).Exp(new(big.Int).Set(a), big.NewInt(-1), m)
	max := m
	for n := int64(1); n < limit; n++ {
		nn := big.NewInt(int64(n))
		nn = nn.Mul(nn, aInv)
		e := nn.Mod(nn, m)
		if e.Cmp(max) < 0 {
			eulercoins = append(eulercoins, big.NewInt(n))
			fmt.Printf("%d - %d\n", e, n)
			max = e
		}
	}
	sum := big.NewInt(0)
	for _, e := range eulercoins {
		sum = sum.Add(sum, e)
	}
	fmt.Println(sum)
}
