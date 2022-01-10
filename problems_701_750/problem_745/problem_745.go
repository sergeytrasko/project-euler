package main

import (
	"fmt"
	"math"
)

func explore() {
	s := func(n int) int {
		limit := int(math.Sqrt(float64(n))) + 1
		for i := limit; i >= 1; i-- {
			if n%(i*i) == 0 {
				return i * i
			}
		}
		return 1
	}

	sum := func(from, to int) int {
		ss := 0
		for i := from; i <= to; i++ {
			ss += s(i)
		}
		return ss
	}
	for i := 1; i < 10; i++ {
		fmt.Printf("%d,", s(i))
	}
	fmt.Println()
	for i := 1; i < 10; i++ {
		fmt.Printf("%d,", sum(1, i))
	}
	fmt.Println()
}

const N = 1e14
const mod = 1_000_000_007

func main() {
	// explore()
	limit := int64(math.Sqrt(N))
	c := make([]int64, limit+2)
	for i := 1; i < len(c); i++ {
		c[i] = 1
	}
	for i := limit; i >= 1; i-- {
		c[i] = N / (i * i)
		for j := int64(2); j <= limit/i; j++ {
			c[i] -= c[i*j]
		}
		// fmt.Printf("c(%d) = %d\n", i, c[i])
	}
	sum := int64(0)
	for i := int64(1); i <= limit; i++ {
		sum = (sum + i*i*c[i]) % mod
	}
	fmt.Println(sum)
}
