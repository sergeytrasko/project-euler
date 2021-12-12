package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Floor(math.Pow(float64(2), float64(30.403243784-x*x))) * 1e-9
}

func main() {
	u, up := float64(-1), float64(0)
	for n := 1; n < 100000; n++ {
		next := f(u)
		// fmt.Printf("f(%d) = %f\n", n, next)
		u, up = next, u
	}
	fmt.Println(u + up)
}
