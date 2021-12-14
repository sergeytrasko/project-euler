package main

import (
	"fmt"
	"math/big"
)

func phi(n int64, primes []int64) int64 {
	ph := n
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if n%p == 0 {
			ph = ph * (p - 1) / p
			for n%p == 0 {
				n /= p
			}
		}
	}
	return ph
}

func sieve(max int64) []int64 {
	n := make([]bool, max)
	primes := []int64{}
	for i := int64(2); i < max; i++ {
		if !n[i] {
			primes = append(primes, i)
			for j := int64(1); ; j++ {
				if i*j >= max {
					break
				}
				n[i*j] = true
			}
		}
	}
	return primes
}

func tetration(a, b, m int64, primes []int64) int64 {
	fmt.Printf("calc t(%d,%d,%d)\n", a, b, m)
	if m == 1 {
		//a^^b mod 1 = 0
		fmt.Printf("t(%d,%d,%d)=%d\n", a, b, m, 0)
		return 0
	}
	if b == 1 {
		//a^^1 = a
		fmt.Printf("t(%d,%d,%d)=%d\n", a, b, m, a)
		return a
	}
	p := phi(m, primes)
	modulo := p

	t1 := tetration(a, b-1, modulo, primes) % modulo
	fmt.Printf("%d^(%d mod %d) mod %d\n", a, t1, modulo, m)
	r := new(big.Int).Exp(big.NewInt(a), big.NewInt(t1), big.NewInt(m)).Int64()
	fmt.Printf("result t(%d,%d,%d)=%d\n", a, b, m, r)
	return r
}

func main() {
	// const a, b, m = 3, 3, 1000 // should be 987
	// const a, b, m = 9, 2, 1000 // should be 489
	const a, b, m = 1777, 1855, 1e8
	primes := sieve(m)
	fmt.Println(tetration(a, b, m, primes))
}
