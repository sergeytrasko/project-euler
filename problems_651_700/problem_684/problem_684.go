package main

import (
	"fmt"
)

func modExp(base, exponent, modulus int64) int64 {
	if modulus == 1 {
		return 0
	}
	if exponent == 0 {
		return 1
	}

	result := modExp(base, exponent/2, modulus)
	result = (result * result) % modulus
	if exponent&1 != 0 {
		return ((base % modulus) * result) % modulus
	}
	return result % modulus
}

func s(n int64, mod int64) int64 {
	if n <= 9 {
		return n
	}
	return (modExp(10, n/9, mod)*(n%9+1) - 1) % mod
}

func ss(n int64, mod int64) int64 {
	if n < 10 {
		sum := int64(0)
		for i := int64(1); i <= n; i++ {
			sum = (sum + s(i, mod)) % mod
		}
		return sum % mod

	}
	chunks := n / 9
	sum := (54 + 540*(modExp(10, chunks-1, mod)-1)/9 - chunks*9) % mod
	for sum < 0 {
		sum += mod
	}
	start := chunks*9 + 1
	for i := start; i <= n; i++ {
		sum = (sum + s(i, mod)) % mod
	}
	return sum % mod
}

func explore() {
	const chunk = 9
	for i := int64(0); i < 5; i++ {
		sum := int64(0)
		for j := int64(1); j <= chunk; j++ {
			sum += s(i*chunk+j, 1_000_000_007)
		}
		fmt.Printf("s(%d...%d)=%d", i*chunk, i*chunk+chunk, sum)
	}
}

func f(m int) int64 {
	a, b := int64(0), int64(1)
	for n := 1; n <= m; n++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	mod := int64(1_000_000_007)
	sum := int64(0)
	for n := 2; n <= 90; n++ {
		sum = (sum + ss(f(n), mod)) % mod
	}
	fmt.Println(sum)
}
