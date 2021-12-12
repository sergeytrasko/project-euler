package main

import (
	"fmt"
	"math"
)

const max = 100_000_000

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(b-a, a)
	}
}

func main() {
	cnt := 0
	for m := 2; m < int(math.Sqrt(float64(max/2))); m++ {
		for n := 1; n < m; n++ {
			if (m+n)%2 == 1 && gcd(m, n) == 1 {
				a, b, c := m*m-n*n, 2*m*n, m*m+n*n
				p := a + b + c
				if p > max {
					break
				}
				k := max / p
				if c%(a-b) == 0 {
					// fmt.Printf("(%d, %d, %d), p=%d\n", a, b, c, p)
					cnt += k
				}
			}
		}
	}
	fmt.Println(cnt)
}
