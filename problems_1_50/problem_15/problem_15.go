package main

import "fmt"

const n = 20

func main() {
	m := [n + 1][n + 1]int{}
	m[0][0] = 1
	for y := 0; y < n+1; y = y + 1 {
		for x := 0; x < n+1; x = x + 1 {
			if x == 0 && y == 0 {
				continue
			}
			left, top := 0, 0
			if y > 0 {
				top = m[y-1][x]
			}
			if x > 0 {
				left = m[y][x-1]
			}
			m[y][x] = top + left
		}
	}
	// for _, r := range m {
	// 	fmt.Println(r)
	// }
	fmt.Println(m[n][n])
}
