package main

import (
	"fmt"
	"math"
)

/*
N = r/k
f(N, k) = (N/k)^k, need to maximize f(N, k), where N is a parameter
basically f'(N, k) = 0

Using Wolfram Alpha
d
--((N/k)^k) = (N/k)^k*(ln(N/k) - 1) = 0
dk

(N/k)^k is always positive, therefore ln(N/k) - 1 = 0

N/k = e

k = round(N/e)
*/
func bestK(n int) int {
	return int(math.Round(float64(n) / math.E))
}

func d(n int) int {
	k := bestK(n)
	// fmt.Printf("For n=%d, k=%d\n", n, k)
	for k%2 == 0 {
		k /= 2
	}
	for k%5 == 0 {
		k /= 5
	}
	if n%k == 0 {
		// fmt.Println("Terminating")
		return -n
	} else {
		// fmt.Println("Nonterminating")
		return n
	}
}

func main() {
	sum := 0
	for n := 5; n <= 10000; n++ {
		sum += d(n)
	}
	fmt.Println(sum)
}
