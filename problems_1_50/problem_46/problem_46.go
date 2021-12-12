package main

import (
	"fmt"
	"math"
)

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func isGoldbachNumber(n int) bool {
	for p := 3; p < n; p += 2 {
		if isPrime(p) {
			x := (n - p) / 2
			s := int(math.Sqrt(float64(x)))
			if s*s == x {
				fmt.Printf("%d = %d + 2*%d^2\n", n, p, s)
				return true
			}
		}
	}
	return false
}

func main() {
	p := 3
	for {
		if !isPrime(p) {
			fmt.Printf("Checking %d\n", p)
			if !isGoldbachNumber(p) {
				break
			}
		}
		p += 2
	}
	fmt.Println(p)
}
