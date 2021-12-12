package main

import (
	"fmt"
	"math"
)

//https://en.wikipedia.org/wiki/Periodic_continued_fraction#Canonical_form_and_repetend
func chain(n int) ([]int, int) {
	type factorization = struct {
		m int
		d int
		a int
	}
	facts := []factorization{}
	r := []int{}
	a0 := int(math.Round(math.Floor(math.Sqrt(float64(n)))))
	m := 0
	d := 1
	a := a0
	r = append(r, a0)
	facts = append(facts, factorization{m: m, d: d, a: a})
	i := 0
	for {
		m = d*a - m
		d = (n - m*m) / d
		a = (a0 + m) / d
		for j := 0; j < len(facts); j++ {
			f := facts[j]
			if f.m == m && f.d == d && f.a == a {
				return r, i - j + 1
			}
		}
		r = append(r, a)
		facts = append(facts, factorization{m: m, d: d, a: a})
		i++
	}
}

func isSquare(n int) bool {
	s := int(math.Round(math.Floor(math.Sqrt(float64(n)))))
	return s*s == n
}

func main() {
	cnt := 0
	for i := 1; i <= 10000; i++ {
		if !isSquare(i) {
			frac, period := chain(i)
			fmt.Printf("sqrt(%d)=%d - %d\n", i, frac, period)
			if period%2 == 1 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
