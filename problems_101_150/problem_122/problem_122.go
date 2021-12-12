package main

import (
	"fmt"
	"math"
)

func findBest(n int, got int, operations int, possibilities []int, best *int) {
	if operations >= *best {
		return
	}
	if n == got {
		if operations < *best {
			// fmt.Println(possibilities)
			*best = operations
		}
		return
	}
	for i := 0; i < len(possibilities); i++ {
		p := possibilities[i]
		if got+p <= n {
			findBest(n, got+p, operations+1, append(possibilities, got+p), best)
		}
	}
}

func m(n int) int {
	//we know that binary exponention gives at most ones(n) + floor(log2(n)) - 1 operations
	//so better approach should give us less.
	//Where ones(n) = number of 1 in binary representation of n
	//As n <= 200 this means that binary representaion of n has less that 8 bits
	best := int(math.Log2(float64(n)) + 7)
	findBest(n, 1, 0, []int{1}, &best)
	return best
}

func main() {
	sum := 0
	for i := 1; i <= 200; i++ {
		sum += m(i)
	}
	fmt.Println(sum)
	// fmt.Println(m(15))
}
