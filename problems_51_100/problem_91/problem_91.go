package main

import (
	"fmt"
	"math"
)

func distance(x1 int, y1 int, x2 int, y2 int) float64 {
	dx, dy := float64(x1-x2), float64(y1-y2)
	return math.Sqrt(dx*dx + dy*dy)
}

func pythagorean(a float64, b float64, c float64) bool {
	return a > 0 && b > 0 && c > 0 && math.Abs(a*a+b*b-c*c) < 0.00001
}

func isRightTriangele(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) bool {
	d1 := distance(x1, y1, x2, y2)
	d2 := distance(x2, y2, x3, y3)
	d3 := distance(x3, y3, x1, y1)

	return pythagorean(d1, d2, d3) || pythagorean(d2, d3, d1) || pythagorean(d1, d3, d2)
}

const max = 50

func main() {
	cnt := 0
	for x1 := 0; x1 <= max; x1++ {
		for y1 := 0; y1 <= max; y1++ {
			for x2 := 0; x2 <= max; x2++ {
				for y2 := 0; y2 <= max; y2++ {
					if isRightTriangele(x1, y1, x2, y2, 0, 0) {
						cnt++
					}
				}
			}
		}
	}
	fmt.Println(cnt / 2)
}
