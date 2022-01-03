package main

import (
	"fmt"
	"math/big"
)

func powers(x *big.Int) []int {
	n := new(big.Int).Set(x)
	d := big.NewInt(1)
	p := 0
	for n.Cmp(d) >= 0 {
		d = new(big.Int).Mul(d, big.NewInt(2))
		p++
	}
	// fmt.Println(d)
	r := make([]int, p+1)
	for p >= 0 {
		// fmt.Printf("n=%d, d=%d, p=%d\n", n, d, p)
		if n.Cmp(d) >= 0 {
			n = new(big.Int).Sub(n, d)
			r[p] = 1
		}
		d = new(big.Int).Div(d, big.NewInt(2))
		p--
	}
	return r
}

func splitPowers(n *big.Int, powers *[]int, from int, cnt *int) {
	*cnt++
	for i := from; i > 0; i-- {
		if ((*powers)[i] == 1 || (*powers)[i] == 2) && (*powers)[i-1] == 0 {
			(*powers)[i]--
			(*powers)[i-1] = 2
			splitPowers(n, powers, i+1, cnt)
			(*powers)[i]++
			(*powers)[i-1] = 0
		}
	}
}

func bruteForce() {
	for x := int64(1); x <= 10; x++ {
		n := new(big.Int).Exp(big.NewInt(10), big.NewInt(x), nil)
		p := powers(n)
		cnt := 0
		splitPowers(n, &p, len(p)-1, &cnt)
		fmt.Printf("f(10^%d)=%d\n", x, cnt)
	}
}

func exlore() {
	for x := int64(1); x <= 100; x++ {
		n := big.NewInt(x)
		p := powers(n)
		cnt := 0
		splitPowers(n, &p, len(p)-1, &cnt)
		fmt.Printf("%d,", cnt)
	}
}

func solve(n *big.Int, cache map[string]int64) int64 {
	add := func(x *big.Int, a int) *big.Int {
		y := new(big.Int).Set(x)
		return y.Add(y, big.NewInt(int64(a)))
	}
	div2 := func(x *big.Int) *big.Int {
		y := new(big.Int).Set(x)
		return y.Div(y, big.NewInt(2))
	}
	even := func(x *big.Int) bool {
		y := new(big.Int).Set(x)
		return y.Mod(y, big.NewInt(2)).Cmp(big.NewInt(0)) == 0
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return 1
	}
	ns := n.String()
	if cache[ns] != 0 {
		return cache[ns]
	}
	var r int64
	if even(n) {
		r = solve(div2(n), cache)
	} else {
		r = solve(div2(add(n, 1)), cache) + solve(div2(add(n, -1)), cache)
	}
	cache[ns] = r
	return r
}

func main() {
	n := new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)
	cache := map[string]int64{}
	fmt.Println(solve(new(big.Int).Add(n, big.NewInt(1)), cache))
}
