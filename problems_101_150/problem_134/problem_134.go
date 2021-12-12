package main

import (
	"fmt"
	"math/big"
)

func isPrime(p int64) bool {
	for i := int64(2); i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func digitsMultiplier(p int64) int64 {
	a := int64(1)
	for p > 0 {
		a *= 10
		p /= 10
	}
	return a
}

func eucledian(a, m int64) (int64, int64, int64) {
	x := *new(big.Int)
	y := *new(big.Int)
	z := new(big.Int).GCD(&x, &y, big.NewInt(int64(a)), big.NewInt(int64(m)))
	return x.Int64(), y.Int64(), z.Int64()
}

func solve(p1, p2 int64) int64 {
	/*
		n ends with p1 and p2 divides n
		a = 10^(number of digits of p1)
		n = x*a + p1 = 0 (mod p2)
		n = x*a = -p1 (mod p2)
	*/
	a := digitsMultiplier(p1) % p2
	b := p2 - p1
	m := p2
	// now solving ax=b(mod m) for x
	// fmt.Printf("%dx=%d(mod %d)\n", a, b, m)
	p, q, _ := eucledian(a, m)
	// fmt.Printf("p=%d, q=%d\n", p, q)
	x0 := (b * p / (a*p + m*q)) % m
	for x0 < 0 {
		x0 += m
	}
	// fmt.Printf("p=%d, q=%d, x0=%d, gcd=%d\n", p, q, x0, gcd)

	// fmt.Printf("x0=%d\n", x0)
	return x0*digitsMultiplier(p1) + p1
}

var primes = []int64{}

const max = 1_000_000

func main() {
	// fmt.Println(solve(5, 7))
	for p := int64(5); p < max; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	for p := int64(max); ; p++ {
		if isPrime(p) {
			primes = append(primes, p)
			break
		}
	}
	sum := int64(0)
	for i := 1; i < len(primes); i++ {
		p1, p2 := primes[i-1], primes[i]
		s := solve(p1, p2)
		// fmt.Printf("p1=%d, p2=%d, s=%d\n", p1, p2, s)
		if s < 0 || s%p2 != 0 {
			fmt.Printf("p1=%d, p2=%d, s=%d\n", p1, p2, s)
			fmt.Println("error!!!")
			return
		}
		sum += s
	}
	fmt.Println(sum)
}
