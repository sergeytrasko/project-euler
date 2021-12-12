package main

import (
	"fmt"
	"math"
	"sort"
)

func permute(current *[]int, available *[]int, size int, outcomes *[]float64) {
	l := len(*current)
	if l == size {
		results := calc(*current, 0, size-1)
		*outcomes = distinct(append((*outcomes), results...))
		// fmt.Println(results)
		return
	}
	for i := 0; i < len(*available); i++ {
		d := (*available)[i]
		if d >= 0 {
			*current = append((*current), d)
			(*available)[i] = -1
			permute(current, available, size, outcomes)
			(*available)[i] = d
			*current = (*current)[0:l]
		}
	}
}

func distinct(x []float64) []float64 {
	sort.Float64s(x)
	r := []float64{x[0]}
	prev := x[0]
	for i := 1; i < len(x); i++ {
		if prev != x[i] {
			prev = x[i]
			r = append(r, prev)
		}
	}

	return r
}

func calc(digits []int, left int, right int) []float64 {
	// fmt.Printf("calc [%d;%d]\n", left, right)
	if left == right {
		return []float64{float64(digits[left])}
	}
	res := []float64{}
	for i := left; i < right; i++ {
		l := calc(digits, left, i)
		r := calc(digits, i+1, right)
		for n := 0; n < len(l); n++ {
			for m := 0; m < len(r); m++ {
				res = append(res, l[n]+r[m])
				res = append(res, l[n]-r[m])
				res = append(res, l[n]*r[m])
				if r[m] != 0 {
					res = append(res, l[n]/r[m])
				}
			}
		}
	}
	return distinct(res)
}

func results(digits []int) int {
	outcomes := []float64{}
	permute(&[]int{}, &digits, 4, &outcomes)
	sort.Float64s(outcomes)
	n := 0
	for i := 0; i < len(outcomes); i++ {
		if outcomes[i] != math.Floor(outcomes[i]) {
			continue
		}
		if outcomes[i] > 0 {
			n++
			if int(outcomes[i]) != n {
				break
			}
		}
	}
	return n
}

func main() {
	// outcomes := []int{}
	// permute(&[]int{}, &[]int{1, 2, 3, 4}, 4, &outcomes)
	// sort.Ints(outcomes)
	// fmt.Println(outcomes)
	// fmt.Println(results([]int{1, 2, 5, 8}))
	// expressions([]int{1, 2, 3}, 0, "")
	// fmt.Println(calc([]int{1, 2}, 0, 1))
	// fmt.Println(calc([]int{1, 2, 3}, 0, 2))

	best := 0
	bestSet := []int{}
	for a := 0; a <= 9; a++ {
		for b := a + 1; b <= 9; b++ {
			for c := b + 1; c <= 9; c++ {
				for d := c + 1; d <= 9; d++ {
					r := results([]int{a, b, c, d})
					if r > best {
						best = r
						bestSet = []int{a, b, c, d}
						// fmt.Println(best)
						// fmt.Println(bestSet)
					}
				}
			}
		}
	}
	// fmt.Println(best)
	// fmt.Println(bestSet)
	fmt.Printf("%d%d%d%d\n", bestSet[0], bestSet[1], bestSet[2], bestSet[3])
}
