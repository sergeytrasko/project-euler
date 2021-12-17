package main

import (
	"fmt"
	"math/big"
)

func explore() {
	// primes := []int{2, 3, 5, 7, 11}
	// n := 32
	best := int64(2 * 3 * 5 * 7 * 11 * 13 * 17)
	for a, two := 0, int64(1); a < 30; a, two = a+1, two*2 {
		for b, three := 0, int64(1); b < 5; b, three = b+1, three+1 {
			for c, five := 0, int64(1); c < 5; c, five = c+1, five*5 {
				for d, seven := 0, int64(1); d < 5; d, seven = d+1, seven*7 {
					for e, eleven := 0, int64(1); e < 5; e, eleven = e+1, eleven*11 {
						for f, thriteen := 0, int64(1); f < 5; f, thriteen = f+1, thriteen*13 {
							for g, seventeen := 0, int64(1); g < 5; g, seventeen = g+1, seventeen*17 {
								if (1+a)*(1+b)*(1+c)*(1+d)*(1+e)*(1+f)*(1+g) != 128 {
									continue
								}
								if two*three*five*seven*eleven*thriteen*seventeen < best {
									best = two * three * five * seven * eleven * thriteen * seventeen
									fmt.Printf("%d - %d - %d - %d - %d - %d -%d\n", two, three, five, seven, eleven, thriteen, seventeen)
									fmt.Printf("%d - %d - %d - %d - %d - %d - %d\n", a, b, c, d, e, f, g)
									fmt.Printf("2^%d * 3^%d * 5^%d * 7^%d * 11^%d * 13^%d * 17^%d = %d\n", a, b, c, d, e, f, g, best)
								}
							}
						}
					}
				}
			}
		}
	}
}

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

type fact = struct {
	p     int64
	power int64
}

func main() {
	modulo := big.NewInt(500500507)
	n := 500_500
	primes := sieve(1e7)
	f := []fact{}
	i := 0
	for n > 0 {
		p := int64(primes[i])
		best := -1
		prevPower := int64(-1)
		for j := 0; j < len(f); j++ {
			q, power := f[j].p, f[j].power
			if power == prevPower && power == 1 {
				break
			}
			nextPower := power + 1
			d := new(big.Int).Exp(big.NewInt(q), big.NewInt(nextPower), nil).Int64()
			if d < p {
				best = j
				break
			}
			prevPower = power
		}
		if best == -1 {
			f = append(f, fact{p, 1})
			i++
		} else {
			f[best].power = f[best].power*2 + 1
		}
		n--
	}
	res := big.NewInt(1)
	for i := 0; i < len(f); i++ {
		s := new(big.Int).Exp(big.NewInt(f[i].p), big.NewInt(f[i].power), modulo)
		res = res.Mul(res, s)
		res = res.Mod(res, modulo)
	}
	fmt.Println(res)
}
