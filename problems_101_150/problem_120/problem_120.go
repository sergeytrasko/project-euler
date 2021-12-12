package main

import (
	"fmt"
	"math/big"
)

func naiveMaxResidue(a int64) int64 {
	max := int64(0)
	for n := int64(1); n <= a*a; n++ {
		x := new(big.Int).Exp(big.NewInt(a+1), big.NewInt(n), nil)
		y := new(big.Int).Exp(big.NewInt(a-1), big.NewInt(n), nil)
		r := new(big.Int).Add(x, y)
		r = new(big.Int).Mod(r, big.NewInt(a*a))
		if r.Int64() > max {
			max = r.Int64()
		}
	}
	return max
}

func maxResidue(a int64) int64 {
	max := int64(0)
	for n := int64(1); n <= a*a; n++ {
		r := int64(0)
		if n%2 == 0 {
			r = 2
		} else {
			r = (2 * n * a) % (a * a)
		}
		if r > max {
			max = r
		}
	}
	return max
}

const maxA = 1000

func main() {
	// naiveSum := int64(0)
	sum := int64(0)
	for a := int64(3); a <= maxA; a++ {
		// r := naiveMaxResidue(a)
		r1 := maxResidue(a)
		// fmt.Printf("%d - %d - %d\n", a, r, r1)
		// fmt.Printf("%d - %d\n", a, r1)
		// naiveSum += r
		sum += r1
	}
	// fmt.Println(naiveSum)
	fmt.Println(sum)
}
