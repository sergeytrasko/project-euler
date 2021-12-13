package main

import (
	"fmt"
	"sort"
)

func isPandigital(n []int64) bool {
	digits := usedDigits(n)
	for i := 0; i < len(digits); i++ {
		if digits[i] != 1 {
			return false
		}
	}
	return true
}

func usedDigits(n []int64) []int {
	digits := make([]int, 10)
	for _, r := range n {
		for r > 0 {
			digits[r%10]++
			r /= 10
		}
	}
	return digits
}

func allDigitsDistinct(n []int64) bool {
	digits := usedDigits(n)
	for i := 0; i < len(digits); i++ {
		if digits[i] > 1 {
			return false
		}
	}
	return true
}

func choose(size int, current []int, available []bool, pos int, result *[][]int) {
	if pos >= len(available) {
		return
	}
	if len(current) == size {
		permutations := [][]int{}
		permute([]int{}, append([]int{}, current...), &permutations)
		*result = append(*result, permutations...)
		return
	}
	for pos < len(available) && !available[pos] {
		pos++
	}
	if pos < len(available) {
		choose(size, current, available, pos+1, result)
		available[pos] = false
		tmp := append([]int{}, current...)
		choose(size, append(tmp, pos), available, pos, result)
		available[pos] = true
	}
}

func permute(cur []int, available []int, result *[][]int) {
	if len(available) == 0 {
		(*result) = append((*result), cur)
	}
	for i := 0; i < len(available); i++ {
		tmp := append([]int{}, available...)
		permute(append(cur, available[i]), append(tmp[:i], tmp[i+1:]...), result)
	}
}

func availableDigits(usedDigits []int) []bool {
	available := make([]bool, 10)
	for i := 0; i < len(available); i++ {
		available[i] = true
	}
	for i, r := range usedDigits {
		if r > 0 {
			available[i] = false
		}
	}
	return available
}

func generatePandigitals(digits int, usedDigits []int) []int64 {
	available := availableDigits(usedDigits)
	result := []int64{}
	c := [][]int{}
	choose(digits, []int{}, available, 0, &c)
	// not clear why, but "choose" generates duplicates
	for _, r := range c {
		if r[0] == 0 {
			continue
		}
		n := int64(0)
		for i := 0; i < len(r); i++ {
			n = n*10 + int64(r[i])
		}
		result = append(result, n)
	}
	return result
}

func firstDigit(n int64) int64 {
	for {
		if n < 10 {
			return n
		}
		n /= 10
	}
}

func bestConcat(n []int64) string {
	d := [10]int64{}
	for i := 0; i < len(n); i++ {
		d[firstDigit(n[i])] = n[i]
	}
	s := ""
	for i := 9; i > 0; i-- {
		if d[i] > 0 {
			s = s + fmt.Sprintf("%d", d[i])
		}
	}
	return s
}

func solve(n []int64, solutions *[]string) {
	if len(n) >= 3 && isPandigital(n) {
		m := []int64{}
		for i := 1; i < len(n); i++ {
			m = append(m, n[0]*n[i])
		}
		if isPandigital(m) {
			r := bestConcat(m)
			fmt.Printf("%d - %s\n", n, r)
			*solutions = append(*solutions, r)
		}
		return
	}
	used := usedDigits(n)
	for l := 1; l < 10; l++ {
		nn := generatePandigitals(l, used)
		if len(nn) == 0 {
			break
		}
		for _, r := range nn {
			dd := usedDigits([]int64{n[0] * r})
			for i := 0; i < 10; i++ {
				if used[i] > 0 && dd[i] > 0 {
					continue
				}
			}
			if len(n) > 1 && r > n[len(n)-1] {
				// let n[1], n[2]... be decreasing
				continue
			}
			next := append([]int64{}, n...)
			next = append(next, r)
			solve(next, solutions)
		}
	}
}

func main() {
	solutions := []string{}
	for i := 1; i <= 5; i++ {
		fmt.Printf("Processing %d-digit n\n", i)
		nn := generatePandigitals(i, []int{})
		for _, r := range nn {
			if r != 1 {
				solve([]int64{r}, &solutions)
			}
		}
	}
	fmt.Printf("Found %d solutions\n", len(solutions))
	sort.Strings(solutions)
	fmt.Println(solutions[len(solutions)-1])
}
