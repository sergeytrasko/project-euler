package main

import "fmt"

func sum(list []int) int {
	s := 0
	for i := 0; i < len(list); i++ {
		s += list[i]
	}
	return s
}

func process(batch int, sheets []int) float64 {
	prob := 0.0
	ss := sum(sheets)
	if ss == 1 {
		if sheets[len(sheets)-1] == 1 {
			return 0 // last batch
		} else {
			prob = 1.0 // only one sheet
		}
	}

	for i := 0; i < len(sheets); i++ {
		if sheets[i] > 0 {
			// cut the sheet in halves
			s := append([]int{}, sheets...)
			s[i]--
			for j := i + 1; j < len(s); j++ {
				s[j]++
			}
			prob += (float64(sheets[i]) / float64(ss)) * process(batch+1, s)
		}
	}
	return prob
}

func main() {
	fmt.Printf("%.6f\n", process(0, []int{0 /*A1*/, 1 /*A2*/, 1 /*A3*/, 1 /*A4*/, 1 /*A5*/}))
}
