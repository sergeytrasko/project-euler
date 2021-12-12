package main

import (
	"fmt"
	"math"
)

const max = 50_000_000

var primes = []int{2}
var found = [max]bool{}

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func pow(index int, p int) int64 {
	m := int64(1)
	for i := 0; i < p; i++ {
		m *= int64(primes[index])
	}
	return m
}

func main() {
	p := 3
	for p <= int(math.Sqrt(max)) {
		if isPrime(p) {
			primes = append(primes, p)
		}
		p += 2
	}
	l := len(primes)
	// fmt.Println(primes)
	for a := 0; a < l; a++ {
		for b := 0; b < l; b++ {
			s := pow(a, 2) + pow(b, 3)
			if s < max {
				for c := 0; c < l && (s+pow(c, 4)) < max; c++ {
					ss := s + pow(c, 4)
					if ss < max {
						// fmt.Println(ss)
						found[ss] = true
					}
				}
			}
		}
	}
	cnt := 0
	for i := 0; i < max; i++ {
		if found[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
