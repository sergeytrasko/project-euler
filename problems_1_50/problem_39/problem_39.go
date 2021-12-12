package main

import "fmt"

func countSolutions(p int) int {
	cnt := 0
	for a := 1; a < p; a++ {
		for b := a; b < p; b++ {
			c := p - a - b
			if c < 0 {
				continue
			}
			if a*a+b*b == c*c {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	max, bestP := 0, 0
	for p := 1; p < 1000; p++ {
		cnt := countSolutions(p)
		if cnt > max {
			max, bestP = cnt, p
		}
	}
	fmt.Println(bestP)
}
