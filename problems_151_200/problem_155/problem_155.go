package main

import (
	"fmt"
	"sort"
)

type frac = struct {
	x, y int64
}

func reduce(f frac) frac {
	d := gcd(f.x, f.y)
	return frac{x: f.x / d, y: f.y / d}
}

func gcd(x, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func sequential(f []frac) frac {
	inverse := []frac{}
	for i := 0; i < len(f); i++ {
		inverse = append(inverse, frac{x: f[i].y, y: f[i].x})
	}
	r := parallel(inverse)
	return frac{x: r.y, y: r.x}
}

func parallel(f []frac) frac {
	y := int64(1)
	for i := 0; i < len(f); i++ {
		y *= f[i].y
	}
	x := int64(0)
	for i := 0; i < len(f); i++ {
		x += f[i].x * y / f[i].y
	}
	return reduce(frac{x: x, y: y})
}

func distinct(f []frac) []frac {
	sort.Slice(f, func(i int, j int) bool {
		a, b := f[i], f[j]
		return a.x*b.y < a.y*b.x
	})
	res := []frac{}
	prev := frac{0, 0}
	for i := 0; i < len(f); i++ {
		if prev.x != f[i].x || prev.y != f[i].y {
			res = append(res, f[i])
			prev = f[i]
		}
	}
	return res
}

var c = frac{x: 1, y: 1}

var cache = [19][]frac{}

func solve(remaining int) []frac {
	if len(cache[remaining]) > 0 {
		return cache[remaining]
	}
	if remaining == 1 {
		return []frac{c}
	}
	res := []frac{}
	for i := 1; i < remaining; i++ {
		a := solve(i)
		for j := i; i+j <= remaining; j++ {
			b := solve(j)
			for _, aa := range a {
				for _, bb := range b {
					res = append(res, sequential([]frac{aa, bb}))
					res = append(res, parallel([]frac{aa, bb}))
				}
			}
		}
	}
	res = distinct(res)
	cache[remaining] = res
	return res
}

func main() {
	res := []frac{}
	for i := 1; i <= 18; i++ {
		sols := solve(i)
		res = distinct(append(res, sols...))
		fmt.Printf("D(%d)=%d\n", i, len(res))
	}
	fmt.Println(len(res))
}
