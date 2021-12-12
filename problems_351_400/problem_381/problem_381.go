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

func s(p int64) int64 {
	if p < 5 {
		return 0
	}
	sum := int64(0)
	for k := int64(1); k <= 5; k++ {
		sum += factMod(p-k, p)
	}
	return sum % p
}

func powMod(x, y, p int64) int64 {
	res := int64(1)
	for y > 0 {
		if y%2 == 1 {
			res = (res * x) % p
		}
		y /= 2
		x = (x * x) % p
	}
	return res
}

func inverse(a, p int64) int64 {
	/*
		Fermat Little Theorem:
		gcd(a, p) = 1 => a^(p-1)=1(mod p)
		If p is prime, then
		a^(p-2) = a^-1 (mod p)
	*/
	// we know that p >= 5, therefore can choose a = 2
	return powMod(a, p-2, p)
}

func factMod(n int64, p int64) int64 {
	res := int64(p - 1)
	for i := n + 1; i < p; i++ {
		res = (res * inverse(i, p)) % p
	}
	return res
}

func main() {
	primes := sieve(1e8)
	fmt.Println(len(primes))
	sum := int64(0)
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if p >= 5 && p < 1e8 {
			sum += s(int64(p))
		}
	}
	fmt.Println(sum)
}
