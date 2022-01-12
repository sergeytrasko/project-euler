package main

import (
	"fmt"
)

func isPowerOf2(x int) bool {
	n := x
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}

func main() {
	p, q := 1, 12345
	total, perfect := 0, 0
	for x := 2; ; x++ {
		if isPowerOf2(x) {
			perfect++
		}
		total++
		k := x*x - x
		if perfect*q < total*p {
			fmt.Printf("P(%d)=%d/%d = %0.10f (need %0.10f)\n", k, perfect, total, float64(perfect)/float64(total), float64(1)/float64(12345))
			fmt.Println(k)
			break
		}
	}
}
