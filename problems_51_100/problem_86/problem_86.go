package main

import (
	"fmt"
	"math"
)

func isSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

const target = 1_000_000

func main() {
	cnt := 0
	for a := 1; ; a++ {
		for b := 1; b <= a; b++ {
			for c := 1; c <= b; c++ {
				if isSquare((b+c)*(b+c) + a*a) {
					// fmt.Printf("(%d, %d, %d)\n", a, b, c)
					cnt++
					if cnt > target {
						fmt.Println(a)
						return
					}
				}
			}
		}
		// fmt.Printf("Max side=%d, solutions=%d\n", a, cnt)
	}
}
