package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func sumDigits(n *big.Int) int {
	s := n.String()
	sum := 0
	for i := 0; i < len(s); i++ {
		d, _ := strconv.Atoi(string(s[i]))
		sum = sum + d
	}
	return sum
}

func twoPower(p int) *big.Int {
	two := big.NewInt(2)
	n := big.NewInt(1)
	for i := 0; i < p; i++ {
		n = n.Mul(n, two)
	}
	return n
}

func main() {
	n := twoPower(1000)
	fmt.Println(sumDigits(n))
}
