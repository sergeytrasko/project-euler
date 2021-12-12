package main

import (
	"fmt"
	"math"
)

func triangular(n int) int {
	return n * (n + 1) / 2
}

func isPerfectSquare(x int) (int, bool) {
	s := int(math.Sqrt(float64(x)))
	return s, s*s == x
}

func isPentagonal(x int) (int, bool) {
	// x = Pn = n(3n - 1) / 2
	// 3n^2 - n - 2x = 0
	// n = (1+- sqrt(1 + 4 * 2x * 3)) / 6 = (1 + sqrt(1 + 24x)) / 6 - should be integer
	d := 1 + 24*x
	s, isSquare := isPerfectSquare(d)
	if !isSquare {
		return 0, false
	}
	return (1 + s) / 6, (1+s)%6 == 0
}

func isHexagonal(x int) (int, bool) {
	// x = Hn = n(2n - 1)
	// 2n^2 - n - x = 0
	// n = (1 +- sqrt(1 + 4 * x * 2)) / 4 = (1 + sqrt(1 + 8x)) / 4
	d := 1 + 8*x
	s, isSquare := isPerfectSquare(d)
	if !isSquare {
		return 0, false
	}
	return (1 + s) / 4, (1+s)%4 == 0
}

func main() {
	n := 1
	for {
		t := triangular(n)
		p, isPent := isPentagonal(t)
		h, isHex := isHexagonal(t)
		if isPent && isHex {
			fmt.Printf("T%d=P%d=H%d=%d\n", n, p, h, t)
			// break
		}
		n++
	}
}
