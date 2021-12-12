package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l = l + 1
		r = r - 1
	}
	return true
}

func main() {
	max := 0
	for a := range [899]int{} {
		for b := range [899]int{} {
			x, y := a+100, b+100
			n := x * y
			if isPalindrome(n) && n > max {
				max = n
				fmt.Printf("%d = %d * %d\n", n, x, y)
			}
		}
	}
	fmt.Println(max)
}
