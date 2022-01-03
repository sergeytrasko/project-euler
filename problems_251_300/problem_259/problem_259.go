package main

import (
	"fmt"
	"sort"
)

type decimal = struct {
	x int64
	y int64
}

func gcd(x, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func simplify(dec decimal) decimal {
	if dec.x == 0 {
		return decimal{x: 0, y: 1}
	}
	d := gcd(dec.x, dec.y)
	return decimal{x: dec.x / d, y: dec.y / d}
}

func add(a, b decimal) decimal {
	return simplify(decimal{
		x: a.x*b.y + b.x*a.y,
		y: a.y * b.y,
	})
}

func sub(a, b decimal) decimal {
	return simplify(decimal{
		x: a.x*b.y - b.x*a.y,
		y: a.y * b.y,
	})
}

func mul(a, b decimal) decimal {
	return simplify(decimal{
		x: a.x * b.x,
		y: a.y * b.y,
	})
}

func div(a, b decimal) decimal {
	return simplify(decimal{
		x: a.x * b.y,
		y: a.y * b.x,
	})
}

func distinct(x []decimal) []decimal {
	sort.Slice(x, func(i, j int) bool {
		return x[i].x*x[j].y-x[j].x*x[i].y < 0
	})
	r := []decimal{x[0]}
	prev := x[0]
	for i := 1; i < len(x); i++ {
		if prev.x != x[i].x || prev.y != x[i].y {
			prev = x[i]
			r = append(r, prev)
		}
	}

	return r
}

func calc(digits []int64, left int, right int) []decimal {
	if left == right {
		return []decimal{{x: digits[left], y: 1}}
	}
	res := []decimal{}
	concat := int64(0)
	for i := left; i <= right; i++ {
		concat = concat*10 + digits[i]
	}
	res = append(res, decimal{x: concat, y: 1})
	for i := left; i < right; i++ {
		l := calc(digits, left, i)
		r := calc(digits, i+1, right)
		for n := 0; n < len(l); n++ {
			for m := 0; m < len(r); m++ {
				res = append(res, add(l[n], r[m]))
				res = append(res, sub(l[n], r[m]))
				res = append(res, mul(l[n], r[m]))
				if r[m].y != 0 {
					res = append(res, div(l[n], r[m]))
				}
			}
		}
	}
	return distinct(res)
}

func main() {
	digits := []int64{}
	for i := 1; i <= 9; i++ {
		digits = append(digits, int64(i))
	}
	res := calc(digits, 0, len(digits)-1)
	sum := int64(0)
	for _, r := range res {
		if r.y == 1 && r.x*r.y > 0 {
			sum += int64(r.x)
		}
	}
	fmt.Println(sum)
}
