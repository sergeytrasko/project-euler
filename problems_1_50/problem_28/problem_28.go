package main

import "fmt"

const n = 1001

var m = [n][n]int{}

func fill() {
	m[n/2][n/2] = 1
	// print()
	direction := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	x, y, p, delta, dir := n/2+1, n/2-1, 2, 2, 0
	for p < n*n-1 {
		for i := 0; i < delta; i++ {
			x += direction[dir][0]
			y += direction[dir][1]
			// fmt.Printf("x=%d,y=%d,dir=%d,p=%d\n", x, y, dir, p)
			m[y][x] = p
			// print()
			p++
		}
		dir++
		if dir == 4 {
			x += 1
			y -= 1
			dir = 0
			delta += 2
			// print()
		}
	}
}

func print() {
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println()
}

func countDiagonals() int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += m[i][i] + m[n-i-1][i]
	}
	sum -= m[n/2][n/2]
	return sum
}

func main() {
	fill()
	fmt.Println(countDiagonals())
}
