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

func sieveRad(max int) []int64 {
	primes := sieve(max)

	rad := make([]int64, max+1)
	for i := 1; i < max; i++ {
		rad[i] = 1
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for j := 1; p*j < max; j++ {
			rad[p*j] *= int64(p)
		}
	}

	return rad
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

const maxC = 120_000

func main() {
	rad := sieveRad(maxC)
	cnt := 0
	sum := int64(0)
	for a := 1; a <= maxC/2; a++ {
		for b := a + 1; a+b < maxC; b++ {
			c := a + b
			ra, rb, rc := rad[a], rad[b], rad[c]
			if ra*rb*rc < int64(c) {
				// fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
				if gcd(a, b) != 1 {
					continue
				}
				cnt++
				sum += int64(c)
			}
		}
	}
	fmt.Println(cnt)
	fmt.Println(sum)
}
