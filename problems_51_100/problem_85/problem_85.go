package main

import "fmt"

func count(m int, n int, w int, h int) int {
	return (w - m + 1) * (h - n + 1)
}

func countRectangles(w int, h int) int {
	sum := 0
	for m := 1; m <= w; m++ {
		for n := 1; n <= h; n++ {
			sum += count(m, n, w, h)
		}
	}
	return sum
}

const max = 100
const target = 2_000_000

func main() {
	p := [max][max]int{}
	for w := 1; w < max; w++ {
		for h := 1; h < max; h++ {
			p[w][h] = countRectangles(w, h)
		}
	}
	closest := target
	area := 0
	for w := 1; w < max; w++ {
		for h := 1; h < max; h++ {
			d := target - p[w][h]
			if d < 0 {
				d = -d
			}
			if d < closest {
				// fmt.Printf("%d, %d\n", p[w][h], w*h)
				closest = d
				area = w * h
			}
		}
	}
	fmt.Println(area)
	// fmt.Println(closest)
}
