package main

import (
	"fmt"
	"math/big"
)

// b^e mod m
func modExp(b, e, m int64) int64 {
	r := int64(1)
	for e > 0 {
		if e%2 == 1 {
			r = (r * b) % m
		}
		b = (b * b) % m
		e /= 2
	}
	return r
}

// degree of p in n!
func factExp(n, p int64) int64 {
	e, u, t := int64(0), p, n
	for u <= t {
		e += t / u
		u *= p
	}
	return e
}

func toBaseB(n, b int64) []int64 {
	d := []int64{}
	for n > 0 {
		d = append(d, n%b)
		n /= b
	}
	return d
}

func fermatBinomial(n, k, p int64) int64 {
	if k > n {
		return 0
	}
	numDegree, denDegree := factExp(n, p)-factExp(n-k, p), factExp(k, p)
	if numDegree > denDegree {
		return 0
	}
	num := int64(1)
	for i := n; i > n-k; i-- {
		cur := int64(i)
		for cur%p == 0 {
			cur /= p
		}
		num = (num * cur) % p
	}
	denom := int64(1)
	for i := int64(1); i < k+1; i++ {
		cur := int64(i)
		for cur%p == 0 {
			cur /= p
		}
		denom = (denom * cur) % p
	}
	return (num * modExp(denom, p-2, p)) % p
}

func lucasBinomial(n, k, p int64) int64 {
	np, kp := toBaseB(n, p), toBaseB(k, p)
	res := int64(1)
	for i := len(np) - 1; i >= 0; i-- {
		ni, ki := np[i], int64(0)
		if i < len(kp) {
			ki = kp[i]
		}
		res = (res * fermatBinomial(ni, ki, p)) % p
	}
	return res
}

func extendedGCD(a, b int64) (int64, int64) {
	x, y, u, v := int64(0), int64(1), int64(1), int64(0)
	for a != 0 {
		q, r := b/a, b%a
		m, n := x-u*q, y-v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}
	return x, y
}

func chineeseReminderTheorem(congurences []struct{ a, m int64 }) int64 {
	m := int64(1)
	for i := 0; i < len(congurences); i++ {
		m *= congurences[i].m
	}
	res := int64(0)
	for i := 0; i < len(congurences); i++ {
		s, _ := extendedGCD(m/congurences[i].m, congurences[i].m)
		res += (congurences[i].a * s * m) / congurences[i].m
	}
	for res < 0 {
		res += m
	}
	return res % m
}

func binomial_naive(n, k, m int64) int64 {
	nn := new(big.Int).MulRange(1, n)
	kk1 := new(big.Int).MulRange(1, k)
	kk2 := new(big.Int).MulRange(1, n-k)

	kk1 = kk1.Mul(kk1, kk2)
	nn = nn.Div(nn, kk1)
	nn = nn.Mod(nn, big.NewInt(m))
	return nn.Int64()
}

func sieve(min, max int) []int {
	n := make([]bool, max)
	primes := []int{}
	for i := 2; i < max; i++ {
		if !n[i] {
			if i >= min {
				primes = append(primes, i)
			}
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

func naive(n, k int64, minP, maxP int) {
	primes := sieve(minP, maxP)
	res := int64(0)
	for pi := 0; pi < len(primes); pi++ {
		p := int64(primes[pi])
		for qi := pi + 1; qi < len(primes); qi++ {
			q := int64(primes[qi])
			for ri := qi + 1; ri < len(primes); ri++ {
				r := int64(primes[ri])
				res += binomial_naive(n, k, p*q*r)
			}
		}
	}
	fmt.Println(res)
}

func solve(n, k int64, minP, maxP int) {
	primes := sieve(minP, maxP)
	binomials := map[int64]int64{}
	for i := 0; i < len(primes); i++ {
		p := int64(primes[i])
		binomials[p] = lucasBinomial(n, k, p)
	}
	res := int64(0)
	for pi := 0; pi < len(primes); pi++ {
		p := int64(primes[pi])
		for qi := pi + 1; qi < len(primes); qi++ {
			q := int64(primes[qi])
			for ri := qi + 1; ri < len(primes); ri++ {
				r := int64(primes[ri])
				res += chineeseReminderTheorem([]struct {
					a int64
					m int64
				}{
					{binomials[p], p},
					{binomials[q], q},
					{binomials[r], r},
				})
			}
		}
	}
	fmt.Println(res)
}

func main() {
	// naive(1e4, 1e3, 1, 100)
	// solve(1e4, 1e3, 1, 100)
	solve(1e18, 1e9, 1000, 5000)
}
