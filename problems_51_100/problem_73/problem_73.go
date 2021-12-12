package main

import (
	"fmt"
	"math"
)

const n = 12_000

func less(a int64, b int64, c int64, d int64) bool {
	// a/b < c/d = > a*d < c*b
	return a*d < c*b
}

//https://en.wikipedia.org/wiki/Farey_sequence#Next_term
func main() {
	a, b := int64(0), int64(1)
	c, d := int64(1), int64(n)
	cnt := int64(0)
	for {
		// fmt.Printf("%d/%d\n", c, d)
		p := int64(math.Floor(float64(n+b)/float64(d)))*c - a
		q := int64(math.Floor(float64(n+b)/float64(d)))*d - b
		a, b, c, d = c, d, p, q
		if less(1, 3, a, b) {
			if !less(a, b, 1, 2) {
				break
			}
			// fmt.Printf("%d/%d\n", a, b)
			cnt++
		}
	}
	fmt.Println(cnt)
}
