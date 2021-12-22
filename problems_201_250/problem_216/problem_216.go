package main

import (
	"fmt"
)

func main() {
	const max = 50_000_000
	t := [max + 1]int64{}
	for n := int64(1); n <= max; n++ {
		t[n] = 2*n*n - 1
	}
	primes := 0
	for n := int64(2); n <= max; n++ {
		p := t[n]
		if p > 1 {
			if p == 2*n*n-1 {
				primes++
			}
			for k := int64(1); n+k*p <= max; k++ {
				for t[n+k*p]%p == 0 {
					t[n+k*p] /= p
				}
			}
			for k := int64(1); -n+k*p <= max; k++ {
				for t[-n+k*p]%p == 0 {
					t[-n+k*p] /= p
				}
			}
		}
	}
	fmt.Println(primes)
}
