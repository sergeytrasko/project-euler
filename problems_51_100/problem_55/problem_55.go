package main

import (
	"fmt"
	"math/big"
)

const maxIterations = 50

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func reverse(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += s[i : i+1]
	}
	return r
}

func isLychrelNumber(n int) bool {
	s := big.NewInt(int64(n))
	for i := 1; i <= maxIterations; i++ {
		r, _ := big.NewInt(0).SetString(reverse(s.String()), 10)
		s = s.Add(s, r)
		if isPalindrome(s.String()) {
			// fmt.Printf("%d is not a Lychel number - found after %d iterations", n, i)
			return false
		}
	}
	return true
}

func main() {
	cnt := 0
	for i := 1; i < 10000; i++ {
		if isLychrelNumber(i) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
