package main

import (
	"fmt"
	"math/big"
)

func generateRightTruncatable(n int64, digitSum int64, depth int, result *[]int64) {
	if n > 10 && big.NewInt(n).ProbablyPrime(10) {
		nn := n / 10
		d := n % 10
		if big.NewInt(nn / (digitSum - d)).ProbablyPrime(10) {
			*result = append(*result, n)
		}
	}
	if n%int64(digitSum) != 0 {
		return
	}
	if depth == 1 {
		return
	}
	for i := int64(0); i <= 9; i++ {
		generateRightTruncatable(n*10+i, digitSum+i, depth-1, result)
	}
}

func main() {
	res := []int64{}
	for i := int64(1); i <= 9; i++ {
		generateRightTruncatable(int64(i), i, 14, &res)
	}
	sum := int64(0)
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	fmt.Println(sum)
}
