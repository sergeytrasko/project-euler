package main

import (
	"fmt"
	"math"
	"sort"
)

func explore() {
	cnt := 0
	for k := int64(1); cnt <= 15; k++ {
		a := 5*k*k + 2*k + 1
		s := int64(math.Round(math.Sqrt(float64(a))))
		if s*s == a {
			cnt++
			fmt.Printf("%d - %d - %d\n", cnt, k, a)
		}
	}
}

func solveMagicFibonacci() {
	a, b := 1, 1
	k := 15
	for n := 3; n <= 2*k+1; n++ {
		c := a + b
		a, b = b, c
	}
	fmt.Println(a * b)
}

func solveDiophantine() {
	type next = func(x, y int64) (int64, int64)

	solve := func(limit int, x, y int64, nextFunction next) []int64 {
		res := []int64{}
		for n := 1; n <= limit; n++ {
			if x > 0 {
				res = append(res, x)
			}
			x, y = nextFunction(x, y)
		}
		return res
	}
	pairs := [][]int64{
		{0, 1},
		{2, 5},
		{0, -1},
		{-1, 2},
		{-1, -2},
	}

	res := []int64{}
	for i := 0; i < len(pairs); i++ {
		x, y := pairs[i][0], pairs[i][1]
		res = append(res, solve(20, x, y, func(x, y int64) (int64, int64) {
			return -9*x - 4*y - 2, -20*x - 9*y - 4
		})...)
		res = append(res, solve(10, x, y, func(x, y int64) (int64, int64) {
			return -9*x + 4*y - 2, 20*x - 9*y + 4
		})...)
	}
	m := map[int64]bool{}
	for _, r := range res {
		m[r] = true
	}
	res = []int64{}
	for k, _ := range m {
		res = append(res, k)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	fmt.Println(res[14])
}

func main() {
	solveMagicFibonacci()
	solveDiophantine()
}
