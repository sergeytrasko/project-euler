package main

import "fmt"

func choose(s []int, taken []int, p int) bool {
	if p == len(s) {
		a, b := []int{}, []int{}
		for i := 0; i < len(s); i++ {
			if taken[i] == 1 {
				a = append(a, s[i])
			} else if taken[i] == 2 {
				b = append(b, s[i])
			}
		}
		if len(a) == 0 || len(b) == 0 {
			return true
		}
		sa, sb := sum(a), sum(b)
		// fmt.Printf("A:%d (sum = %d), B:%d (sum=%d)\n", a, sa, b, sb)
		if sa == sb {
			return false
		}
		if len(a) > len(b) {
			return sa > sb
		}
		if len(b) > len(a) {
			return sb > sa
		}
		return true
	}
	if !choose(s, taken, p+1) {
		return false
	}
	taken[p] = 1
	if !choose(s, taken, p+1) {
		return false
	}
	taken[p] = 2
	if !choose(s, taken, p+1) {
		return false
	}
	taken[p] = 0
	return true
}

func isSpecialSumSet(s []int) bool {
	taken := make([]int, len(s))
	return choose(s, taken, 0)
}

func sum(s []int) int {
	ss := 0
	for _, r := range s {
		ss += r
	}
	return ss
}

const delta = 3

func tryCandidates(s []int, p int, best *[]int, bestSum *int) {
	if p == len(s) {
		if isSpecialSumSet(s) && sum(s) < *bestSum {
			*best = s
			*bestSum = sum(s)
			fmt.Printf("%d - %d\n", *best, *bestSum)
		}
		return
	}
	original := s[p]
	for i := -delta; i <= delta; i++ {
		s[p] = original + i
		tryCandidates(s, p+1, best, bestSum)
		s[p] = original
	}
}

func main() {
	n6 := []int{11, 18, 19, 20, 22, 25}
	b := n6[len(n6)/2]
	best := []int{b}
	for i := 0; i < len(n6); i++ {
		best = append(best, n6[i]+b)
	}
	bestSum := sum(best)
	candidate := []int{}
	for i := 0; i < len(best); i++ {
		candidate = append(candidate, best[i])
	}
	fmt.Printf("%d - %d\n", best, bestSum)
	tryCandidates(candidate, 0, &best, &bestSum)
	for i := 0; i < len(best); i++ {
		fmt.Print(best[i])
	}
	fmt.Println()
}
