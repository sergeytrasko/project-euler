package main

import (
	"fmt"
	"math"
	"strconv"
)

func isSquare(n int64) bool {
	nn := int64(math.Round(math.Floor(math.Sqrt(float64(n)))))
	return nn*nn == n
}

func ok(p int64) bool {
	ways := 0
	for m := int64(1); m*m*m < p; m++ {
		n2 := p - m*m*m
		if isSquare(n2) {
			ways++
			if ways > 4 {
				return false
			}
		}
	}
	return ways == 4
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
	// fmt.Println(ok(5229225))
	r := []int{}
	n := 1
	for len(r) < 5 {
		for i := -1; i <= 9; i++ {
			p := buildPalindrome(n, i)
			if ok(int64(p)) {
				// fmt.Println(p)
				r = append(r, p)
			}
		}
		n++
	}
	sum := 0
	for i := 0; i < 5; i++ {
		sum += r[i]
	}
	fmt.Println(sum)
}
