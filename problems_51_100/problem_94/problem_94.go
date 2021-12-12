package main

import (
	"fmt"
	"math"
)

func isSquare(x int64) bool {
	s := int64(math.Sqrt(float64(x)))
	return s*s == x
}

const max = 1_000_000_000

func main() {
	// 2 * a + a + 1 <= max
	// 3a + 1 <= max
	// a is odd, as (a+-1)/2 should be an integer to have triangle's height to be an integer as well
	// h = sqrt(a^2 - ((a+-1)/2)^)
	// h should be an integer <=> a^2 - ((a+-1)/2)^ should be a perfect square
	sum := int64(0)
	maxA := int64((max - 1) / 3)
	for a := int64(5); a < maxA; a += 2 {
		if isSquare(a*a - (a+1)*(a+1)/4) {
			fmt.Printf("(%d, %d, %d) - %d\n", a, a, a+1, 3*a+1)
			sum += (3*a + 1)
		}
		if isSquare(a*a - (a-1)*(a-1)/4) {
			fmt.Printf("(%d, %d, %d) - %d\n", a, a, a-1, 3*a-1)
			sum += (3*a - 1)
		}
	}
	fmt.Println(sum)
}
