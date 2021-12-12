package main

import (
	"fmt"
	"math/big"
)

func digitSum(a int, b int) int {
	ba, bb := big.NewInt(int64(a)), big.NewInt(int64(b))
	digits := ba.Exp(ba, bb, nil).String()
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += int(digits[i] - byte('0'))
	}
	return sum
}

func main() {
	max := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			ds := digitSum(a, b)
			if ds > max {
				max = ds
			}
		}
	}
	fmt.Println(max)
}
