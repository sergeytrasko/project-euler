package main

import "fmt"

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	d := 2
	n := 1
	side := 3
	num, primes := 1, 0
	fmt.Println(n)
	for {
		for i := 0; i < 4; i++ {
			n += d
			// fmt.Println(n)
			if isPrime(n) {
				primes++
			}
			num++
		}
		ratio := float64(primes) / float64(num)
		fmt.Printf("side %d: %d/%d=%f\n", side, primes, num, float64(100.0*ratio))
		d += 2
		side += 2
		if ratio < 0.1 {
			break
		}
	}
}
