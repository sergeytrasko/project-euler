package main

import (
	"fmt"
	"math"
)

const m = 100

type point = struct {
	x int
	y int
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func countPointsOnLine(a, b point) int {
	x, y := b.x-a.x, b.y-a.y
	if x == 0 {
		return y - 1
	}
	if y == 0 {
		return x - 1
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return gcd(x, y) - 1
}

//https://en.wikipedia.org/wiki/Pick%27s_theorem
func countPointsInside(a, b, c, d int) int {
	// A = inner + boundry/2 - 1
	area := float64((a + c) * (b + d) / 2)
	boundry := float64(4 +
		countPointsOnLine(point{x: a, y: 0}, point{x: 0, y: b}) +
		countPointsOnLine(point{x: 0, y: b}, point{x: -c, y: 0}) +
		countPointsOnLine(point{x: -c, y: 0}, point{x: 0, y: -d}) +
		countPointsOnLine(point{x: 0, y: -d}, point{x: a, y: 0}))
	return int(math.Round(area - boundry/2 + 1))
}

func isSquare(n int) bool {
	x := int(math.Round(math.Floor(math.Sqrt(float64(n)))))
	return x*x == n
}

func main() {
	cnt := 0
	for a := 1; a <= m; a++ {
		for b := 1; b <= m; b++ {
			for c := 1; c <= m; c++ {
				for d := 1; d <= m; d++ {
					points := countPointsInside(a, b, c, d)
					if isSquare(points) {
						cnt++
					}
				}
			}
		}
	}
	fmt.Println(cnt)
}
