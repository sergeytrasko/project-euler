package main

import (
	"fmt"
	"math/big"
)

func isPrime(p int64) bool {
	pp := big.NewInt(p)
	return pp.ProbablyPrime(1)
}

func naive() {
	max := int64(150_000_000)
	sum := int64(0)
	for n := int64(10); n < max; n+=10 {
		if n%2 == 1 || n%3 == 0 || n%7 == 0 || n%13 == 0 || n%10 != 0 {
			continue
		}
		if isPrime(n*n+1) &&
			isPrime(n*n+3) &&
			!isPrime(n*n+5) &&
			isPrime(n*n+7) &&
			isPrime(n*n+9) &&
			!isPrime(n*n+11) &&
			isPrime(n*n+13) &&
			!isPrime(n*n+15) &&
			!isPrime(n*n+17) &&
			!isPrime(n*n+19) &&
			!isPrime(n*n+21) &&
			!isPrime(n*n+23) &&
			!isPrime(n*n+25) &&
			isPrime(n*n+27) {
			fmt.Println(n)
			sum += n
		}
	}
	fmt.Println(sum)
	fmt.Println("-----")
}

func sieve() {
	const max = 150_000_000
	const dLen = 6
	d := []int64{1, 3, 7, 9, 13, 27}
	t := [max+1][dLen]int64{}
	for n := int64(0); n <= max; n++ {
		for i := 0; i < len(d); i++ {
			t[n][i] = n*n+d[i]
		}
	}
	primeSum := int64(0)
	for n := int64(1); n <= max; n++ {
		prime := true
		for j := 0; j < len(d); j++ {
			p := t[n][j]
			fmt.Printf("n=%d, j=%d, p=%d\n", n, d[j], p)
			if p > 1 {
				if p != n*n+d[j] {
					prime = false
				}
				for i := j; i < len(d); i++ {
					for k := int64(1); n+k*p <= max; k++ {
						for t[n+k*p][i]%p == 0 {
							t[n+k*p][i] /= p
						}
					}
					for k := int64(1); -n+k*p <= max; k++ {
						if -n+k*p >= 0 {
							for t[-n+k*p][i]%p == 0 {
								t[-n+k*p][i] /= p
							}
						}
					}
				}
			} else {
				prime = false
			}
		}
		if prime && n >= 10 {
			fmt.Println(n)
			primeSum += n
		}
	}
	fmt.Println(primeSum)
}

func main() {
	naive()
	// sieve()
}
