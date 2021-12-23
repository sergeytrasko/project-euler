package main

import (
	"fmt"
	"math"
	"sort"
)

func explore() {
	cnt := 0
	for k := int64(1); cnt <= 30; k++ {
		a := 5*k*k + 14*k + 1
		s := int64(math.Round(math.Sqrt(float64(a))))
		if s*s == a {
			cnt++
			fmt.Printf("%d - %d - %d\n", cnt, k, a)
		}
	}
}

type next = func(x, y int64) (int64, int64)

func solve(limit int, x, y int64, nextFunction next) []int64 {
	res := []int64{}
	for n := 1; n <= limit; n++ {
		if x > 0 {
			res = append(res, x)
		}
		x, y = nextFunction(x, y)
	}
	return res
}

func main() {
	pairs := [][]int64{
		{2, -7},
		{0, -1},
		{0, 1},
		{-4, 5},
		{-3, 2},
		{-3, -2},
	}

	res := []int64{}
	for i := 0; i < len(pairs); i++ {
		x, y := pairs[i][0], pairs[i][1]
		res = append(res, solve(20, x, y, func(x, y int64) (int64, int64) {
			return -9*x - 4*y - 14, -20*x - 9*y - 28
		})...)
		res = append(res, solve(10, x, y, func(x, y int64) (int64, int64) {
			return -9*x + 4*y - 14, 20*x - 9*y + 28
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
	sum := int64(0)
	for i := 0; i < 30; i++ {
		sum += res[i]
	}
	fmt.Println(sum)
}
