package main

import (
	"fmt"
	"strconv"
)

func isPrime(p int) bool {
	if p == 1 {
		return false
	}
	for n := 2; n*n <= p; n++ {
		if p%n == 0 {
			return false
		}
	}
	return true
}

func rightTruncatable(p int) bool {
	for {
		p = p / 10
		if !isPrime(p) {
			return false
		}
		if p == 0 {
			return true
		}
	}
}

func leftTruncatable(p int) bool {
	s := strconv.Itoa(p)
	for i := 0; i < len(s); i++ {
		ss := s[i:]
		n, _ := strconv.Atoi(ss)
		if !isPrime(n) {
			return false
		}
	}
	return true
}

func main() {
	n := 11
	cntRemaining := 11
	sum := 0
	for cntRemaining > 0 {
		if isPrime(n) && leftTruncatable(n) && rightTruncatable(n) {
			// fmt.Println(n)
			sum += n
			cntRemaining--
		}
		n += 2
	}
	fmt.Println(sum)
}
