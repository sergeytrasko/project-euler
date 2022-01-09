package main

import (
	"fmt"
	"math"
	"math/big"
)

const n = 20
const mod = 1_000_000_000

type digits = [10]int

func split(current, target, usedDigits, index int, d digits) int64 {
	if current > target {
		return 0
	}
	if current == target {
		s := int64(0)
		d0 := d[0]
		d[0] = n - usedDigits // fill as many 0 as possible
		s += sum(d)
		d[0] = d0
		return s
	}
	if index >= len(d) {
		return 0
	}
	s := int64(0)
	for i := 0; i <= n-usedDigits; i++ {
		d[index] = i
		s += split(current+index*index*i, target, usedDigits+i, index+1, d)
		d[index] = 0
	}
	return s
}

func sum(d digits) int64 {
	s := int64(0)
	s1 := int64(0)
	for i := 0; i < len(d); i++ {
		s += int64(d[i])
		s1 += int64(i * d[i])
	}
	r := new(big.Int).MulRange(1, s)
	for i := 0; i < len(d); i++ {
		r = r.Div(r, new(big.Int).MulRange(1, int64(d[i])))
	}
	r = r.Mul(r, big.NewInt(s1))
	r = r.Div(r, big.NewInt(s))
	r = r.Mul(r,
		new(big.Int).Div(
			new(big.Int).Sub(
				new(big.Int).Exp(big.NewInt(10), big.NewInt(s), nil),
				big.NewInt(1),
			),
			big.NewInt(9),
		),
	)
	r = r.Mod(r, big.NewInt(mod))
	return r.Int64()
}

func bruteForce(n int) int64 {
	limit := int64(1)
	for i := 1; i <= n; i++ {
		limit *= 10
	}
	s := int64(0)
	for i := int64(1); i < limit; i++ {
		ds := int64(0)
		x := i
		for x > 0 {
			d := x % 10
			ds += d * d
			x /= 10
		}
		sqrt := int64(math.Sqrt(float64(ds)))
		if sqrt*sqrt == ds {
			s += i
		}
	}
	return s
}

func main() {
	// fmt.Println(bruteForce(n))
	s := int64(0)
	for i := 1; i*i <= n*81; i++ {
		s += split(0, i*i, 0, 1, digits{})
	}
	fmt.Println(s % mod)
}
