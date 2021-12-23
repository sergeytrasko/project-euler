package main

import "fmt"

const maxLen = 18
const fiveToTen = 9765625

func toDigits(index int) [10]int {
	r := [10]int{}
	for d := 0; d <= 9; d++ {
		r[d] = index % 5
		index /= 5
	}
	return r
}

func toIndex(d [10]int) int {
	index := 0
	for i := 9; i >= 0; i-- {
		index = index*5 + d[i]
	}
	return index
}

func copy(digits [10]int) [10]int {
	x := [10]int{}
	for i := 0; i <= 9; i++ {
		x[i] = digits[i]
	}
	return x
}

func main() {
	dp := [maxLen + 1][fiveToTen]int64{}
	for d := 1; d <= 9; d++ {
		digits := [10]int{}
		digits[d] = 1
		dp[1][toIndex(digits)] = 1
	}
	for l := 2; l <= maxLen; l++ {
		fmt.Printf("L = %d\n", l)
		for index := 0; index < fiveToTen; index++ {
			digits := toDigits(index)
			for d := 0; d <= 9; d++ {
				x := copy(digits)
				if x[d] <= 3 {
					x[d]++
				}
				dp[l][toIndex(x)] += dp[l-1][index]
			}
		}
	}
	count := int64(0)
	for index := 0; index < fiveToTen; index++ {
		digits := toDigits(index)
		lessThan3 := true
		for d := 0; d <= 9; d++ {
			lessThan3 = lessThan3 && (digits[d] <= 3)
		}
		if lessThan3 {
			count += dp[maxLen][index]
		}
	}

	fmt.Println(count)
}
