package main

import (
	"fmt"
	"sort"
)

func isPandigital(x int64, base int) bool {
	d := map[int]int{}
	for x > 0 {
		d[int(x%int64(base))]++
		x /= int64(base)
	}
	return len(d) == base
}

func toNumber(d []int) int64 {
	n := len(d)
	x := int64(0)
	for i := 0; i < n; i++ {
		x = x*int64(n) + int64(d[i])
	}
	return x
}

func isSuperPandigital(d []int) bool {
	n := len(d)
	x := toNumber(d)
	for i := n - 1; i >= 2; i-- {
		if !isPandigital(x, i) {
			return false
		}
	}
	return true
}

func permutations(d []int, limit int) {
	findCeil := func(x []int, first int, l, h int) int {
		idx := l
		for i := l + 1; i <= h; i++ {
			if x[i] > first && x[i] < x[idx] {
				idx = i
			}
		}
		return idx
	}
	sum := int64(0)
	sort.Ints(d)
	finished := false
	for !finished {
		// fmt.Println(d)
		if d[0] != 0 {
			if isSuperPandigital(d) {
				fmt.Println(d)
				sum += toNumber(d)
				limit--
				if limit == 0 {
					fmt.Println(sum)
					break
				}
			}
		}
		i := 0
		for i = len(d) - 2; i >= 0; i-- {
			if d[i] < d[i+1] {
				break
			}
		}
		if i == -1 {
			finished = true
		} else {
			ceil := findCeil(d, d[i], i+1, len(d)-1)
			d[i], d[ceil] = d[ceil], d[i]
			tail := d[i+1:]
			sort.Ints(tail)
			d = append(append([]int{}, d[0:i+1]...), tail...)
		}
	}
}

const N = 12

func main() {
	d := []int{}
	for i := 0; i < N; i++ {
		d = append(d, i)
	}
	d[0] = 1
	d[1] = 0

	permutations(d, 10)
	// fmt.Println(isSuperPandigital([]int{1, 2, 4, 0, 3}))
	// fmt.Println(isSuperPandigital([]int{1, 0, 9, 3, 2, 6, 5, 7, 8, 4}))
}
