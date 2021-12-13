package main

import "fmt"

var primes = []int{}

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for p := 2; p < 100; p++ {
		if isPrime(p) {
			primes = append(primes, p)
		}
	}
	pp := 1
	d := 1
	ratio := float64(15499) / float64(94744)
	for i := 0; ; i++ {
		pp *= (primes[i] - 1)
		d *= primes[i]
		fmt.Printf("prime %d - %f - %f\n", primes[i], float64(pp)/float64(d-1), ratio)
		if (float64(pp) / float64(d-1)) < ratio {
			d /= primes[i]
			pp /= (primes[i] - 1)
			for j := 1; ; j++ {
				fmt.Printf("%d/%d=%.10f, ratio=%.10f\n", j*pp, j*d-1, float64(j*pp)/float64(j*d-1), ratio)
				if (float64(j*pp) / float64(j*d-1)) < ratio {
					fmt.Println(j * d)
					break
				}
			}
			break
		}
	}
}
