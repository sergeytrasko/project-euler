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

func factorial(p int) *big.Int {
	return big.NewInt(1).MulRange(1, int64(p))
}

func main() {
	n := factorial(100)
	fmt.Println(sumDigits(n))
}
