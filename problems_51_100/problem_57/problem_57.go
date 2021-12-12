package main

import (
	"fmt"
	"math/big"
)

func expansion(depth int) (*big.Int, *big.Int) {
	// 1/(2 + p/q) = 1/((2q + p)/q) = q / (2q + p)
	p, q := big.NewInt(1), big.NewInt(2)
	for n := 0; n < depth; n++ {
		two := big.NewInt(2)
		p, q = q, p.Add(p, two.Mul(two, q))
	}
	// 1 + p/q = (p + q) / q
	p, q = p.Add(p, q), q
	return p, q
}

func main() {
	cnt := 0
	for i := 0; i < 1000; i++ {
		p, q := expansion(i)
		// fmt.Printf("%d/%d\n", p, q)
		pDigits, qDigits := len(p.String()), len(q.String())
		if pDigits > qDigits {
			cnt++
			// fmt.Printf("%d - %d/%d\n", i, p, q)
		}
	}
	fmt.Println(cnt)
}
