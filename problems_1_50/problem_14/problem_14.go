package main

import "fmt"

func distance(n int) int {
	cnt := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		cnt = cnt + 1
	}
	return cnt
}

func main() {
	max := 0
	n := 1
	best := 0
	for n < 1000000 {
		d := distance(n)
		if d > max {
			max, best = d, n
		}
		n = n + 1
	}
	fmt.Printf("%d gives %d\n", best, max)
}
