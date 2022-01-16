package main

import (
	"fmt"
	"math/big"
)

const n = 2020
const mod = 1e16

type digits = [10]int

func split(current, target, usedDigits, index int, d digits) int64 {
	if current > target {
		return 0
	}
	if current == target {
		s := int64(0)
		d0, c0 := d[0], d[current]
		d[current]++
		d[0] = n - 1 - usedDigits // fill as many 0 as possible
		s += sum(d)
		d[0] = d0
		d[current] = c0
		return s
	}
	if index >= len(d) {
		return 0
	}
	s := int64(0)
	for i := 0; i < n-usedDigits; i++ {
		d[index] = i
		s += split(current+index*i, target, usedDigits+i, index+1, d)
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

func main() {
	// s := int64(0)
	// for i := 1; i*i <= n*81; i++ {
	// 	s += split(0, i*i, 0, 1, digits{})
	// }
	// fmt.Println(s % mod)
	s := int64(0)
	for i := 1; i <= 9; i++ {
		s += split(0, i, 0, 1, digits{})
	}
	fmt.Println(s % mod)
}
