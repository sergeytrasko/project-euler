package main

import (
	"fmt"
	"sort"
)

func sieve(max int) []int {
	n := make([]bool, max)
	primes := []int{}
	for i := 2; i < max; i++ {
		if !n[i] {
			primes = append(primes, i)
			for j := 1; ; j++ {
				if i*j >= max {
					break
				}
				n[i*j] = true
			}
		}
	}
	return primes
}

func isPrime(p int, primes []int) bool {
	i := sort.SearchInts(primes, p)
	return i < len(primes) && primes[i] == p
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

const max = 1e8

// const max = 100

func main() {
	primes := sieve(max)
	fmt.Printf("%d primes\n", len(primes))
	sum := int64(0)
	/*
		for i := 0; i < len(primes); i++ {
			a := primes[i]
			if i%10_000 == 0 {
				fmt.Println(i)
			}
			for j := i + 1; j < len(primes); j++ {
				b := primes[j]
				q := float64(b+1) / float64(a+1)
				cf := float64(b+1) * q
				if int(cf) > max {
					// fmt.Printf("break a=%d, b=%d, c=%d, %d\n", a, b, int(cf), j-i)
					break
				}
				if cf == float64(int(cf)) && isPrime(int(cf-1), primes) {
					// fmt.Printf("(%d, %d, %d)\n", a, b, int(cf-1))
					sum += (int64(a) + int64(b) + int64(cf-1))
				}
			}
		}
	*/
	// geometric sequence is : k*x*x, k*x*y, k*y*y, where gcd(x, y) = 1
	for x := 1; x*x < max; x++ {
		for k := 1; k*x*x < max; k++ {
			a := k*x*x - 1
			if !isPrime(a, primes) {
				continue
			}
			for y := 1; y < x; y++ {
				b := k*x*y - 1
				c := k*y*y - 1
				if !isPrime(b, primes) || !isPrime(c, primes) {
					continue
				}
				if gcd(x, y) != 1 {
					continue
				}
				sum += int64(a + b + c)
			}
		}
	}
	fmt.Println(sum)
}
