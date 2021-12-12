package main

import "fmt"

const max = 20_000
const mod = 1_000_000_007

type factorization = map[int]int

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

func factor(n int, primes []int) factorization {
	res := factorization{}
	for i := 0; i < len(primes) && primes[i] <= n; i++ {
		p := primes[i]
		for n%p == 0 {
			res[p]++
			n /= p
		}
		if n == 1 {
			break
		}
	}
	return res
}

func add(f1 *factorization, f2 factorization, sign int) {
	for f := range f2 {
		(*f1)[f] += (sign * f2[f])
		if (*f1)[f] == 0 {
			delete(*f1, f)
		}
	}
}

//https://en.wikipedia.org/wiki/Binomial_coefficient#Products_of_binomial_coefficients
func b(n int, facts [max + 1]factorization) factorization {
	f := factorization{}
	for k := 1; k <= n; k++ {
		fk := facts[k]
		add(&f, fk, 2*k-n-1)
	}
	return f
}

//https://stackoverflow.com/questions/1522825/calculating-sum-of-geometric-series-mod-m
func geomSequenceSum(n, b, mod int64) int64 {
	t := int64(1)
	e := b % mod
	total := int64(0)
	for n > 0 {
		if n%2 == 1 {
			total = (e*total + t) % mod
		}
		t = ((e + 1) * t) % mod
		e = (e * e) % mod
		n = n / 2
	}
	return total
}

//https://www2.math.upenn.edu/~deturck/m170/wk3/lecture/sumdiv.html
func d(n int, facts [max + 1]factorization, mod int64) int64 {
	f := b(n, facts)
	fmt.Printf("%d - %d\n", n, len(f))
	res := int64(1)
	for p := range f {
		m := int64(1)
		// sum of geometric sequence
		// a1 = 1, r = p
		// m = sum = (1 - r^n)/(1-r)
		m = geomSequenceSum(int64(f[p]+1), int64(p), mod)
		// naive implementation
		// x := int64(1)
		// for i := 1; i <= f[p]; i++ {
		// 	x = (x * int64(p)) % mod
		// 	m = (m + x) % mod
		// }
		res = (res * m) % mod
	}

	return res
}

func s(n int, facts [max + 1]factorization, mod int64) int64 {
	res := int64(0)
	for i := 1; i <= n; i++ {
		res = (res + d(i, facts, mod)) % mod
	}
	return res
}

func main() {
	primes := sieve(max)
	facts := [max + 1]factorization{}
	for i := 2; i <= max; i++ {
		facts[i] = factor(i, primes)
	}
	fmt.Println(s(20000, facts, mod))
	// fmt.Println(s(5, facts, mod)) //5736
	// fmt.Println(s(100, facts, mod)) //332792866

}
