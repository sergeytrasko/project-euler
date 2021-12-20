package main

import "fmt"

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func countNaive(n int) int {
	cnt := 0
	for a := 1; ; a++ {
		r := float64(1)/float64(n) - float64(1)/float64(a)
		// fmt.Printf("%d, %f, %f\n", a, float64(1)/float64(a), r)
		if float64(1)/float64(a) < r {
			break
		}
		for b := 1; b <= a; b++ {
			d := gcd(a, b)
			if d == 1 {
				// x, y, z := a*(a+b), b*(a+b), a*b
				z := a * b
				if n%z == 0 {
					// k := n / z
					// fmt.Printf("1/%d+1/%d=1/%d\n", x*k, y*k, n)
					cnt++
				}
			}
		}
	}
	return cnt
}

var primes = []int{}

func isPrime(p int) bool {
	for i := 0; i < len(primes); i++ {
		pp := primes[i]
		if pp*pp > p {
			break
		}
		if p%pp == 0 {
			return false
		}
	}
	return true
}

func generatePrimes(max int) {
	primes = append(primes, 2)
	for p := 3; p < max; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	fmt.Printf("Found %d primes below %d\n", len(primes), max)
}

func squareDivisionsCount(n int) int {
	divisorsCount := 1
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		c := 0
		for n%p == 0 {
			n /= p
			c++
		}
		divisorsCount *= (1 + 2*c)
		if n == 1 {
			break
		}
	}
	return divisorsCount
}

func divisorsCount(n int) int {
	divisorsCount := 1
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		c := 0
		for n%p == 0 {
			n /= p
			c++
		}
		divisorsCount *= (1 + c)
	}
	return divisorsCount
}

func count(n int) int {
	// total number of solutions = number of divisors of n^2 - equivalent solutions
	cnt := squareDivisionsCount(n)
	// cnt := divisorsCount(n * n)
	return (cnt + 1) / 2
}

func main() {
	generatePrimes(100000)
	best := 0
	for n := 2; ; n++ {
		cnt2 := count(n)
		if cnt2 > best {
			best = cnt2
			fmt.Printf("%d - %d\n", n, cnt2)
			if best > 1000 {
				break
			}
		}
	}
}
