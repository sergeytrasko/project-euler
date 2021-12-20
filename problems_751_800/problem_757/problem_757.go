package main

import (
	"fmt"
	"math"
)

const max = 1e14

func main() {
	solutions := make(map[int64]struct{}, int(math.Sqrt(max)))
	limit := int64(math.Sqrt(math.Sqrt(max)))
	found := struct{}{}
	for x := int64(1); x < limit; x++ {
		for y := x; ; y++ {
			n := x * y * (x + 1) * (y + 1)
			if n > max {
				break
			}
			solutions[n] = found
		}
	}
	fmt.Println(len(solutions))
}
