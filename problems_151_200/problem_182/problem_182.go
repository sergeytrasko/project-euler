package main

import (
	"fmt"
)

const p = 1009
const q = 3643

// const p = 19
// const q = 37
const n = p * q
const phi = (p - 1) * (q - 1)

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func main() {
	bestE, best := []int{}, n
	for e := 1; e < phi; e++ {
		// fmt.Printf("e=%d, best so far=%d, bestE=%d\n", e, best, bestE)
		if gcd(e, phi) == 1 {
			fixedPoints := (1 + gcd(e-1, p-1)) * (1 + gcd(e-1, q-1))
			if fixedPoints < best {
				best = fixedPoints
				bestE = []int{e}
			} else if fixedPoints == best {
				bestE = append(bestE, e)
			}
		}
	}
	sum := 0
	for i := 0; i < len(bestE); i++ {
		sum += bestE[i]
	}
	fmt.Println(sum)
}
