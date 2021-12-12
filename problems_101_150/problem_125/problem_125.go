package main

import (
	"fmt"
	"math"
	"strconv"
)

const max = 100_000_000

func isSumOfSquares(n int) bool {
	if n > max {
		return false
	}
	maxN := int(math.Sqrt(float64(n))) + 1
	sum, lower, upper := 1, 1, 1
	for sum != n {
		if sum < n {
			upper++
			sum += (upper * upper)
		} else if sum > n {
			sum -= (lower * lower)
			lower++
		}
		if upper > maxN {
			return false
		}
	}
	// fmt.Printf("%d, %d - %d\n", lower, upper, n)
	return upper > lower
}

func buildPalindrome(x, middle int) int {
	s := fmt.Sprintf("%d", x)
	s1 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s1 = s1 + s[i:i+1]
	}
	if middle >= 0 {
		s += fmt.Sprintf("%d", middle)
	}
	s += s1
	r, _ := strconv.Atoi(s)
	return r
}

func main() {
	sum := 0
	cnt := 0
	for n := 1; n <= 9; n++ {
		if isSumOfSquares(n) {
			sum += n
			cnt++
		}
	}
	for i := 1; i*i < max; i++ {
		n := buildPalindrome(i, -1)
		if isSumOfSquares(n) {
			// fmt.Printf("%d - %d\n", i, n)
			sum += n
			cnt++
		}
		for j := 0; j <= 9; j++ {
			n = buildPalindrome(i, j)
			if isSumOfSquares(n) {
				// fmt.Printf("%d, %d - %d\n", i, j, n)
				sum += n
				cnt++
			}
		}
	}
	fmt.Println(cnt)
	fmt.Println(sum)
}
