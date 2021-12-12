package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	maxDivisor := math.Sqrt(float64(n)) + 1
	d := 3
	for d < int(maxDivisor) {
		if n%d == 0 {
			return false
		}
		d = d + 2
	}
	return true
}

func main() {
	n := 3
	sum := 2
	for n < 2000000 {
		if isPrime(n) {
			sum = sum + n
		}
		n = n + 2
	}
	fmt.Println(sum)
}
