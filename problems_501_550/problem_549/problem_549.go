package main

import (
	"fmt"
)

type factorization = map[int]int

func sieveSmallestPrimeFactors(max int) []int {
	spf := make([]int, max+1)
	for i := 1; i <= max; i++ {
		spf[i] = i
	}
	for i := 4; i <= max; i += 2 {
		spf[i] = 2
	}
	for i := 3; i*i <= max; i++ {
		if spf[i] == i {
			for j := i * i; j <= max; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}
	return spf
}

func factorize(n int, spf []int) factorization {
	f := factorization{}
	for n > 1 {
		factor := spf[n]
		f[factor]++
		n = n / factor
	}
	return f
}

func minFactorial(n int, f factorization) int {
	res := 1
	for k, v := range f {
		// find min y such that for x = y! we have x a multiple of k^v
		y := 0
		for v > 0 {
			next := y + k
			for next%k == 0 {
				next /= k
				v--
			}
			y += k
		}
		if y > res {
			res = y
		}
	}
	return res
}

func solve() {
	const N = 100_000_000
	spf := sieveSmallestPrimeFactors(N)
	sum := int64(0)
	for i := 2; i <= N; i++ {
		f := factorize(i, spf)
		sum += int64(minFactorial(i, f))
	}
	fmt.Println(sum)
}

func main() {
	solve()
}
