package main

import (
	"fmt"
	"math"
)

func p(m int) int {
	/*
		Investigating in Wolfram Alpha:
		https://www.wolframalpha.com/input/?i=maximize%28a*b%5E2*c%5E3*d%5E4*e%5E5%29%2Ca%2Bb%2Bc%2Bd%2Be%3D5%2Ca%3E0%2Cb%3E0%2Cc%3E0%2Cd%3E0%2Ce%3E0

		it turns out that for:
		m = 2: x1 = 2/3, x2 = 4/3
		m = 3: x1 = 2/4, x2 = 4/4, x3 = 6/4
		m = 4: x1 = 2/5, x2 = 4/5, x3 = 6/5, x4 = 8/5
		m = 5: x1 = 2/6, x2 = 4/6, x3 = 6/6, x4 = 8/6, x5 = 10/6

		pattern: x(i, m) = 2*i/(m+1)
	*/
	res := 1.0
	for i := 1; i <= m; i++ {
		xi := float64(2*i) / float64(m+1)
		res = res * math.Pow(xi, float64(i))
	}
	return int(math.Round(math.Floor(res)))
}

func main() {
	sum := 0
	for m := 2; m <= 15; m++ {
		sum += p(m)
	}
	fmt.Println(sum)
}
