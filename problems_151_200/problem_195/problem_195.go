package main

import (
	"fmt"
	"math"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func area(a, b, c int) float64 {
	s := float64(a+b+c) / 2
	return math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
}

func count(maxRadius int) int {
	cnt := 0
	for m := 2; m < 4*maxRadius; m++ {
		for n := 1; 2*n <= m; n++ {
			if gcd(n, m) == 1 {
				a, b, c := m*m-m*n+n*n, 2*m*n-n*n, m*m-n*n
				if a == b && b == c {
					continue
				}
				d := 1
				if a%3 == 0 && b%3 == 0 && c%3 == 0 {
					d = 3
				}
				s := area(a, b, c)
				r := s / (float64(a+b+c) / 2)
				if r > float64(maxRadius)*3 {
					break
				}
				r = r / float64(d)
				if r < float64(maxRadius) {
					cnt += int(float64(maxRadius) / r)
				}
			}
		}
	}
	return cnt
}

func main() {
	// fmt.Printf("T(%d) = %d, expected %d\n", 100, count(100), 1234)
	// fmt.Printf("T(%d) = %d, expected %d\n", 1000, count(1000), 22767)
	// fmt.Printf("T(%d) = %d, expected %d\n", 10000, count(10000), 359912)
	fmt.Println(count(1053779))
}
