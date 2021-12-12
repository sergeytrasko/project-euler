package main

import "fmt"

const max = 10_000_000

var primes = []int{}

func isPrime(p int) bool {
	for i := 2; i*i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func phi(n int) (int, []int) {
	primeDivisors := []int{}
	p := n
	for i := 0; i < len(primes) && primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			for p%primes[i] == 0 {
				p = p / primes[i]
			}
			primeDivisors = append(primeDivisors, primes[i])
		}
	}
	r := float64(n)
	for i := 0; i < len(primeDivisors); i++ {
		r *= float64(1.0 - 1.0/float64(primeDivisors[i]))
	}

	return int(r), primeDivisors
}

func orderDigits(n int) string {
	digits := [10]int{}
	s := ""
	for n > 0 {
		digits[n%10]++
		n = n / 10
	}
	for i := 0; i < len(digits); i++ {
		for j := 0; j < digits[i]; j++ {
			s = s + string(byte(i)+byte('0'))
		}
	}
	return s
}

func naive() {
	// fmt.Println(phi(8319823))
	// fmt.Printf("%s, %s, %t, %f\n", orderDigits(8319823), orderDigits(phi(8319823)), orderDigits(8319823) == orderDigits(phi(8319823)), float64(8319823)/float64(phi(8319823)))
	best := float64(10000000)
	bestN := 0
	for n := 2; n < max; n++ {
		if n%50_000 == 0 {
			fmt.Println(n)
		}
		p, divisors := phi(n)
		if orderDigits(n) == orderDigits(p) {
			ratio := float64(n) / float64(p)
			if ratio < best {
				bestN = n
				best = ratio
				fmt.Printf("%d/phi(%d)=%d/%d=%f, divisors=%d\n", n, n, n, p, ratio, divisors)
			}
		}
	}
	fmt.Println(bestN)
	// Naive approach shows that most probably N is a multiple of 2 prime numbers
}

func main() {
	for p := 2; p < max; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	fmt.Println(len(primes))
	best := float64(10000000)
	bestN := 0
	for p := 0; p < len(primes) && primes[p]*primes[p] < max; p++ {
		for q := p + 1; primes[q]*primes[p] < max; q++ {
			pp, qq := primes[p], primes[q]

			n := pp * qq
			// phi(p*q) = p*q*1(1-1/p)*(1-1/q) = (p-1)(q-1)
			phiN := (pp - 1) * (qq - 1)
			if orderDigits(n) == orderDigits(phiN) {
				ratio := float64(n) / float64(phiN)
				if ratio < best {
					best = ratio
					bestN = n
					fmt.Printf("%d/phi(%d)=%d/%d=%f\n", n, n, n, phiN, ratio)
				}
			}
		}
	}
	fmt.Println(bestN)
}
