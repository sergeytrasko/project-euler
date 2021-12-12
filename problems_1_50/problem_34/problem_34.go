package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial(n int) *big.Int {
	return big.NewInt(1).MulRange(1, int64(n))
}

func digitFactorialSum(n *big.Int) *big.Int {
	s := n.String()
	sum := big.NewInt(0)
	for i := 0; i < len(s); i++ {
		x, _ := strconv.Atoi(string(s[i]))
		sum = sum.Add(sum, factorial(x))
	}
	return sum
}

func main() {
	n := 3
	sum := 0
	for n < 10000000 { // there should be some reasonable threshold where sum of factorials becomes bigger than the number
		f := big.NewInt(int64(n))
		digits := digitFactorialSum(f)
		if f.Cmp(digits) == 0 {
			fmt.Println(n)
			sum += n
		}
		n++
	}
	fmt.Printf("sum=%d\n", sum)
}
