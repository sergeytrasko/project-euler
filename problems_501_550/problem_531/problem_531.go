package main

import "fmt"

func extendedGCD(a, b int64) (int64, int64) {
	x, y, u, v := int64(0), int64(1), int64(1), int64(0)
	for a != 0 {
		q, r := b/a, b%a
		m, n := x-u*q, y-v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}
	return x, y
}

func g(a, m, b, n int64) int64 {
	u, v := extendedGCD(m, n)
	g := u*m + v*n
	if a%g != b%g {
		return 0
	}
	mod := m * n / g
	r := a*v*(n/g) + b*u*(m/g)
	for r < 0 {
		r += mod * 1e6
	}
	return r % mod
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

func totient(n int64, primes []int64) int64 {
	nn := n
	factors := []int64{}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if n%p == 0 {
			factors = append(factors, p)
			for n%p == 0 {
				n /= p
			}
		}
		if n == 1 {
			break
		}
	}
	t := nn
	for i := 0; i < len(factors); i++ {
		p := factors[i]
		t = t * (p - 1) / p
	}
	return t
}

func main() {
	// fmt.Println(g(2, 4, 4, 6))
	// fmt.Println(g(3, 4, 4, 6))

	min := int64(1000000)
	max := int64(1005000)
	primes := sieve(max)
	p := map[int64]int64{}
	for i := min; i < max; i++ {
		p[i] = totient(i, primes)
	}
	sum := int64(0)
	for n := min; n < max; n++ {
		for m := n + 1; m < max; m++ {
			x := g(p[n], n, p[m], m)
			sum += x
		}
	}
	fmt.Println(sum)
}
