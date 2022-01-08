package main

import (
	"fmt"
	"sort"
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

func drs(n int) int {
	return (n-1)%9 + 1
}

func mdrs(n int, spf []int, cache map[int]int) int {
	factors := getFactor(factorize(n, spf))
	best := 0
	for _, f := range factors {
		var r = 0
		if f == n {
			r = drs(n)
		} else {
			r = drs(f) + cache[n/f]
		}
		if r > best {
			best = r
		}
	}
	cache[n] = best
	return best
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	sort.Ints(list)
	return list
}

func getFactor(f factorization) []int {
	res := []int{1}
	for divisor, times := range f {
		for i := 1; i <= times; i++ {
			newDivisors := []int{}
			for j := 0; j < len(res); j++ {
				newDivisors = append(newDivisors, res[j]*divisor)
			}
			res = unique(append(res, newDivisors...))
		}
	}
	return res[1:] //exclude 1 from divisors
}

const max = 1e6

func main() {
	spf := sieveSmallestPrimeFactors(max)
	cache := map[int]int{}
	for n := 1; n < max; n++ {
		r := mdrs(n, spf, cache)
		cache[n] = r
	}
	sum := 0
	for n := 1; n < max; n++ {
		sum += cache[n]
	}
	fmt.Println(sum)
}
