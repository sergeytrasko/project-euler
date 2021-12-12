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
	n := 600851475143
	maxDivisor := math.Sqrt(float64(n)) + 1
	d := 3
	best := 0
	for d < int(maxDivisor) {
		if n%d == 0 && isPrime(d) {
			best = d
		}
		d = d + 2
	}
	fmt.Println(best)
}
