package main

import "fmt"

func generateRepunit(n int, b int64) int64 {
	x := int64(0)
	for i := 0; i < n; i++ {
		x = x*int64(b) + 1
	}
	return x
}

func sign(x int64) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

func isIntegerBase(x int64, n int) bool {
	l, r := int64(1), x
	for l < r {
		m := (l + r) / 2
		rl := x - generateRepunit(n, l)
		rm := x - generateRepunit(n, m)
		if rm == 0 {
			return true
		}
		if sign(rl)*sign(rm) > 0 {
			l = m
		} else {
			r = m
		}
	}
	return false
}

const limit = 1e12

func main() {
	sum := int64(1) //1 is also a strong rep-unit
	reps := make(map[int64]bool)
	for b := int64(2); ; b++ {
		found := false
		for n := 3; ; n++ {
			r := generateRepunit(n, b)
			if r > limit {
				break
			}
			for m := 2; m < n; m++ {
				if isIntegerBase(r, m) {
					if !reps[r] {
						sum += r
						reps[r] = true
					}
					break
				}
			}
			found = true
		}
		if !found {
			break
		}
	}
	fmt.Println(sum)
}
