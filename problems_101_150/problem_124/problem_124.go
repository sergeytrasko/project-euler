package main

import (
	"fmt"
	"sort"
)

type Rad = struct {
	n   int
	rad int
}

func rad(n int) int {
	m := 1
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			m *= i
			for n%i == 0 {
				n /= i
			}
		}
	}
	return m
}

func main() {
	m := []Rad{}
	for n := 1; n <= 100_000; n++ {
		m = append(m, Rad{n, rad(n)})
	}
	sort.Slice(m, func(i, j int) bool {
		ri, rj := m[i].rad, m[j].rad
		if ri == rj {
			return m[i].n < m[j].n
		}
		return ri < rj
	})
	fmt.Println(m[10_000-1].n)
}
