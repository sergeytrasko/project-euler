package main

import "fmt"

func permutations(min int, max int, result *[][]int) {
	var p func(cur []int, available []int, result *[][]int)
	p = func(cur []int, available []int, result *[][]int) {
		if len(available) == 0 {
			(*result) = append((*result), cur)
		}
		for i := 0; i < len(available); i++ {
			tmp := append([]int{}, available...)
			p(append(cur, available[i]), append(tmp[:i], tmp[i+1:]...), result)
		}
	}
	avail := []int{}
	for i := min; i <= max; i++ {
		avail = append(avail, i)
	}
	p([]int{}, avail, result)
}

func naivePandigital() int {
	res := [][]int{}
	permutations(0, 9, &res)
	cnt := 0
	for i := 0; i < len(res); i++ {
		if res[i][0] != 0 {
			x := 0
			for j := len(res[i]) - 1; j >= 0; j-- {
				x = x*10 + res[i][j]
			}
			if x%11 == 0 {
				cnt++
			}
		}
	}
	return cnt
}

const dig = 2
const dd = dig + 1
const max = dd * dd * dd * dd * dd * dd * dd * dd * dd * dd // dig^10

func indexToNumbers(x int) [10]int {
	res := [10]int{}
	for i := 9; i >= 0; i-- {
		res[i] = x % dd
		x /= dd
	}
	return res
}

func numbersToIndex(n [10]int) int {
	res := 0
	for i := 0; i <= 9; i++ {
		res = res*dd + n[i]
	}
	return res
}

func main() {
	mods := [1 + 10*dig][max][11]int{}
	mods[0][0][0] = 1
	div := -1
	for l := 1; l <= 10*dig; l++ {
		div = -div
		for d := 0; d < max; d++ {
			numbers := indexToNumbers(d)
			for n := 0; n < 10; n++ {
				if l == 1 && n == 0 {
					continue // cannot start with 0
				}
				if numbers[n] < dig {
					// digit n is not yet used in full
					numbers[n]++
					for j := 0; j < 11; j++ {
						mods[l][numbersToIndex(numbers)][j] += mods[l-1][d][(j+11+div*n)%11]
					}
					numbers[n]--
				}
			}
		}
	}
	fmt.Println(mods[10*dig][max-1][0])
	// fmt.Println(naivePandigital())

}
