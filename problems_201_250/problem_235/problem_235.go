package main

import (
	"fmt"
	"math"
)

const target = -600_000_000_000.0

func u(k int, r float64) float64 {
	return float64(900-3*k) * math.Pow(r, float64(k-1))
}

func s(n int, r float64) float64 {
	sum := float64(0)
	for i := 1; i <= n; i++ {
		sum += u(i, r)
	}
	return sum
}

func f(r float64) float64 {
	return s(5000, r) - target
}

func sign(r float64) int {
	if r < 0 {
		return -1
	} else {
		return 1
	}
}

func main() {
	// r is between 1 and 1.0005 accoring to WolframAlpha
	// fmt.Printf("%.10f\n", f(0.99))
	fmt.Printf("%.10f\n", f(1.002))
	fmt.Printf("%.10f\n", f(1.02))
	l, r := 1.002, 1.02
	for r-l > 1e-13 {
		m := (l + r) / 2
		fl, fr, fm := f(l), f(r), f(m)
		fmt.Printf("[%.12f, %.12f] -> [%.15f, %.15f]\n", l, r, fl, fr)
		if sign(fl)*sign(fm) > 0 {
			l = m
		} else {
			r = m
		}
	}
	m := (l + r) / 2
	fmt.Printf("%.12f\n", f(m))
}
