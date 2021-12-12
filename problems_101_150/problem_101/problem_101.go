package main

import "fmt"

type fn = func(int) int

func f(n int) int {
	// return n * n * n
	r := 1
	m := -n
	for i := 1; i <= 10; i++ {
		r += m
		m *= (-n)
	}
	return r
}

func lj(n int, k int, j int, f fn) float64 {
	r := 1.0
	for m := 1; m <= k; m++ {
		if m != j {
			r = r * float64(n-m) / float64(j-m)
		}
	}
	return r
}

func lagrange(n int, p int, f fn) int64 {
	l := float64(0.0)
	for j := 1; j <= n; j++ {
		l += float64(f(j)) * lj(p, n, j, f)
	}
	return int64(l)
}

//https://en.wikipedia.org/wiki/Lagrange_polynomial
func main() {
	maxN := 10
	sum := int64(0)
	for n := 1; n <= maxN; n++ {
		l := lagrange(n, n+1, f)
		fmt.Printf("n=%d, f(%d)=%d, OP(%d,n)=%d\n", n, n, f(n), n, l)
		sum += l
	}
	fmt.Println(sum)
}
