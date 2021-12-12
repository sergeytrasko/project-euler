package main

import (
	"fmt"
	"math/big"
)

func mul2(n *big.Int) *big.Int {
	return new(big.Int).Mul(big.NewInt(2), n)
}

func square(n *big.Int) *big.Int {
	x := new(big.Int)
	x.Set(n)
	return x.Mul(x, x)
}

func isSquare(n *big.Int) bool {
	x := new(big.Int).Set(n)
	s := x.Sqrt(x)
	return square(s).Cmp(n) == 0
}

/*
y - number of blues
x - total number
y/x * (y-1)/(x-1) = 1/2

x^2-2y^2-x+2y=0
https://www.alpertron.com.ar/QUAD.HTM

x(n+1) = 3x(n)+4y(n)-3
y(n+1) = 2n(x)+3y(n)-2

We know that x = 21 and y = 15 is a solution
*/
func main() {
	min := int64(1e12)
	x, y := int64(21), int64(15)
	for y < min {
		fmt.Printf("blue=%d, total=%d\n", y, x)
		x, y = 3*x+4*y-3, 2*x+3*y-2
	}
}
