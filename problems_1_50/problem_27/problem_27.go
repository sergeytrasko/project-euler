package main

import "fmt"

func isPrime(n int64) bool {
	if n < 0 {
		return false
	}
	for p := int64(2); p*p <= n; p++ {
		if n%p == 0 {
			return false
		}
	}
	return true
}

func count(a int64, b int64) int64 {
	n := int64(0)
	for {
		p := n*n + a*n + b
		if !isPrime(p) {
			break
		}
		n++
	}
	return n - 1
}

func main() {
	longest := int64(0)
	bestMult := 0
	limit := 1000
	for a := -limit; a <= limit; a++ {
		for b := -limit; b <= limit; b++ {
			n := count(int64(a), int64(b))
			if n > longest {
				bestMult = a * b
				longest = n
			}
		}
	}
	fmt.Println(bestMult)
}
