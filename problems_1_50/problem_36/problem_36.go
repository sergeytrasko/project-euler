package main

import (
	"fmt"
	"strconv"
)

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

func main() {
	sum := 0
	for n := 1; n < 1000000; n++ {
		if isPalindrome(strconv.FormatInt(int64(n), 10)) && isPalindrome(strconv.FormatInt(int64(n), 2)) {
			fmt.Printf("%s, %s \n", strconv.FormatInt(int64(n), 10), strconv.FormatInt(int64(n), 2))
			sum += n
		}
	}
	fmt.Println(sum)
}
