package main

import "fmt"

func digits(set []int) []int {
	res := []int{}
	for _, r := range set {
		if r == 6 || r == 9 {
			res = append(res, []int{6, 9}...)
		} else {
			res = append(res, r)
		}
	}
	return res
}

func checkSets(a []int, b []int) bool {
	r := [100]bool{}
	for _, aa := range a {
		for _, bb := range b {
			r[aa*10+bb] = true
			r[bb*10+aa] = true
		}
	}
	return r[1] && r[4] && r[9] && r[16] && r[25] && r[36] && r[49] && r[64] && r[81]
}

func choose(cur []int, n int, p int, target int, res *[][]int) {
	if p == n {
		if len(cur) == target {
			*res = append(*res, cur)
		}
		return
	}
	choose(cur, n, p+1, target, res)
	next := append([]int{}, cur...)
	choose(append(next, p), n, p+1, target, res)
}

func main() {
	combinations := [][]int{}
	choose([]int{}, 10, 0, 6, &combinations)
	cnt := 0
	for i := 0; i < len(combinations); i++ {
		for j := i + 1; j < len(combinations); j++ {
			if checkSets(digits(combinations[i]), digits(combinations[j])) {
				cnt++
				// fmt.Printf("%d - %d\n", combinations[i], combinations[j])
			}
		}
	}
	fmt.Println(cnt)
}
