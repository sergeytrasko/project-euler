package main

import (
	"fmt"
	"math"
	"math/big"
)

const max = 64_000_000

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

func gcd(x, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func main() {
	primes := sieve(max)
	fmt.Printf("Found %d primes\n", len(primes))
	sigma2 := [max + 1]int64{}
	// sigma2[1] = 1
	for i := 1; i < max; i++ {
		sigma2[i] = 1
	}
	// https://en.wikipedia.org/wiki/Divisor_function
	// sigma2(n) = mul ( (p^(a+1)*2 - 1)/(p^2-1) )
	// if gcd(a, b) = 1 then sigma2(a, b) = sigma2(a) * sigma2(b)
	for i := 0; i < len(primes); i++ {
		p := int64(primes[i])
		n := p
		j := 1
		for n < max {
			// fmt.Println(n)
			for k := int64(1); n*k < max; k++ {
				if gcd(p, k) == 1 {
					s := new(big.Int).Exp(big.NewInt(p), big.NewInt(int64((j+1)*2)), nil)
					s = s.Sub(s, big.NewInt(1))
					s = new(big.Int).Div(s, big.NewInt(p*p-1))
					sigma2[n*k] *= s.Int64()
				}
			}
			n *= p
			j++
		}
	}
	sum := int64(0)
	for i := int64(1); i < max; i++ {
		s := int64(math.Sqrt(float64(sigma2[i])))
		if s*s == sigma2[i] {
			sum += i
		}
	}
	fmt.Println(sum)
}
