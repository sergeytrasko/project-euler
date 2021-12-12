package main

import "fmt"

func isBouncy(n int) bool {
	if n < 10 {
		return false
	}
	// x := n
	prevDigit := n % 10
	n = n / 10
	increasing, decreasing := false, false
	for n > 0 {
		thisDigit := n % 10
		if prevDigit > thisDigit {
			increasing = true
		} else if prevDigit < thisDigit {
			decreasing = true
		}
		prevDigit = thisDigit
		n = n / 10
	}
	// fmt.Printf("%d - increasing %t, decreasing %t\n", x, increasing, decreasing)
	return increasing && decreasing
}

func main() {
	bouncyCnt := 0
	n := 1
	for {
		if isBouncy(n) {
			bouncyCnt++
		}
		if n*99 == bouncyCnt*100 {
			break
		}
		n++
	}
	fmt.Println(n)
}
