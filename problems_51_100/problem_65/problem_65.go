package main

import (
	"fmt"
	"math/big"
)

const max_depth = 1000 + 1

var z = []*big.Int{}

func fillZ() {
	k := 1
	for len(z) < max_depth {
		z = append(z, big.NewInt(1))
		z = append(z, big.NewInt(int64(2*k)))
		z = append(z, big.NewInt(1))
		k++
	}
}

func expansion(n int, depth int) (*big.Int, *big.Int) {
	if n == depth {
		return big.NewInt(1), new(big.Int).Set(z[n])
	}
	ep, eq := expansion(n+1, depth)

	p := new(big.Int).Set(eq)
	q := ep.Add(ep, eq.Mul(eq, z[n]))
	// fmt.Printf("level %d - %d/%d, %d\n", n, p, q, z[n])
	return p, q
}

func calc(depth int) (*big.Int, *big.Int) {
	p, q := expansion(0, depth)

	two := big.NewInt(2)
	pp := two.Mul(two, q)
	pp = pp.Add(pp, p)
	qq := q
	return pp, qq
}

const n = 100

func main() {
	fillZ()
	p, q := calc(n - 2)
	fmt.Printf("%d/%d\n", p, q)
	s := p.String()
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - byte('0'))
	}
	fmt.Println(sum)
}
