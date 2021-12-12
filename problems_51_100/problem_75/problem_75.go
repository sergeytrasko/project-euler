package main

import "fmt"

func gcd(a, b int64) int64 {
	if a == b {
		return a
	} else if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(b-a, a)
	}
}

const limit = 1_500_000

func main() {
	solutions := [limit + 1]int64{}
	for m := int64(1); m*m < limit; m++ {
		for n := int64(1); n < m; n++ {
			if (m+n)%2 == 1 && gcd(m, n) == 1 {
				a, b, c := m*m-n*n, 2*m*n, m*m+n*n
				l := a + b + c
				for k := int64(1); k*l <= limit; k++ {
					solutions[k*l]++
				}
			}
		}
	}
	cnt := 0
	for i := 0; i < len(solutions); i++ {
		if solutions[i] == 1 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
