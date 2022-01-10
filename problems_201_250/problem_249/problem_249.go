package main

import "fmt"

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

const mod = 1e16

func main() {
	s := sieve(5000)
	max := 0
	for _, r := range s {
		max += r
	}
	fmt.Println(max)
	ways := make([]int64, max+1)
	ways[0] = 1
	for _, p := range s {
		for i := max; i >= p; i-- {
			ways[i] = (ways[i] + ways[i-p]) % mod
		}
	}
	s = sieve(max + 1)
	sum := int64(0)
	for _, p := range s {
		sum = (sum + ways[p]) % mod
	}
	fmt.Println(sum)
}
