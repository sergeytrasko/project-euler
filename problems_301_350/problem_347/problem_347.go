package main

import (
	"fmt"
	"math"
)

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

func mm(p, q, n int) int {
	max := 0
	for i := 1; ; i++ {
		pp := int(math.Pow(float64(p), float64(i)))
		if pp > n {
			break
		}
		for j := 1; ; j++ {
			qq := int(math.Pow(float64(q), float64(j)))
			if pp*qq > n {
				break
			}
			if pp*qq > max {
				max = pp * qq
			}
		}
	}

	return max
}

func findM(n int, primes []int) int {
	best := 0
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			m := mm(primes[i], primes[j], n)
			if m > best {
				best = m
			}
		}
	}
	return best
}

// const max = 10_000_000

// const max = 100 + 1
const N = 10_000_000

func main() {
	primes := sieve(N + 1)
	fmt.Println(len(primes))
	m := [N + 1]bool{}
	for pp := 0; pp < len(primes); pp++ {
		p := primes[pp]
		for qq := pp + 1; qq < len(primes); qq++ {
			best := -1
			q := primes[qq]
			// pq := p * q
			pMul := 1
			for k := 1; ; k++ {
				if pMul*p > N {
					break
				}
				pMul = pMul * p
				qMul := pMul
				for l := 1; ; l++ {
					if qMul*q > N {
						break
					}
					qMul = qMul * q
					if qMul > best {
						best = qMul
					}
				}
			}
			if p*q > N {
				break
			}
			m[best] = true
		}
	}
	// fmt.Println(m)
	sum := 0
	for i := 0; i <= N; i++ {
		if m[i] {
			sum += i
		}
	}
	fmt.Println(sum)

	// x := [max]int{}
	// for i := 0; i < max; i++ {
	// 	x[i] = findM(i, primes)
	// 	if m[i] != 0 && m[i] != x[i] {
	// 		fmt.Printf("!!!%d - %d - %d\n", i, m[i], x[i])
	// 	}
	// }
	// fmt.Println(x)
	// sum := 0
	// for i := 0; i < N; i++ {
	// 	sum += m[i]
	// }
	// fmt.Println(sum)

	/*
		x := [N + 1]bool{}
		for pp := 0; pp < len(primes); pp++ {
			for qq := pp + 1; qq < len(primes); qq++ {
				bestM := mm(primes[pp], primes[qq], N)
				x[bestM] = true
			}
		}
		sum2 := 0
		for i := 0; i <= N; i++ {
			if x[i] {
				sum2 += i
			}
		}
		fmt.Println(sum2)
	*/
}
