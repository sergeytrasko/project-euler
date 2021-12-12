package main

import (
	"fmt"
	"strconv"
)

var primes = []int{}
var primesMap = map[int]bool{}

func isPrime(p int) bool {
	for i := 0; i < len(primes); i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	return true
	// for i := 2; i*i <= p; i++ {
	// 	if p%i == 0 {
	// 		return false
	// 	}
	// }
	// return true
}

func concat(a, b int) int {
	s := fmt.Sprintf("%d%d", a, b)
	c, _ := strconv.Atoi(s)
	return c
}

func isRemarkable(p []int) bool {
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			if i != j {
				n := concat(p[i], p[j])
				if !primesMap[n] {
					return false
				}
				// if !isPrime(concat(p[i], p[j])) {
				// 	return false
				// }
			}
		}
	}
	return true
}

func main() {
	for i := 2; i < 20_000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Printf("Generated %d primes\n", len(primes))
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			n := concat(primes[i], primes[j])
			if isPrime(n) {
				primesMap[n] = true
			}
			n = concat(primes[j], primes[i])
			if isPrime(n) {
				primesMap[n] = true
			}
		}
	}
	fmt.Println("Generated concatenate primes map")
	for a := 0; a < len(primes); a++ {
		pa := primes[a]
		for b := a + 1; b < len(primes); b++ {
			pb := primes[b]
			if !isRemarkable([]int{pa, pb}) {
				continue
			}
			for c := b + 1; c < len(primes); c++ {
				pc := primes[c]
				if !isRemarkable([]int{pa, pb, pc}) {
					continue
				}
				for d := c + 1; d < len(primes); d++ {
					pd := primes[d]
					if !isRemarkable([]int{pa, pb, pc, pd}) {
						continue
					}
					for e := d + 1; e < len(primes); e++ {
						pe := primes[e]
						if isRemarkable([]int{pa, pb, pc, pd, pe}) {
							fmt.Println([]int{pa, pb, pc, pd, pe})
							fmt.Println(pa + pb + pc + pd + pe)
						}
					}
				}
			}
		}
	}
	// fmt.Println(isRemarkable([]int{3, 7, 109, 673}))
}
