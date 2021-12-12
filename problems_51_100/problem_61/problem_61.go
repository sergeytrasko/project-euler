package main

import (
	"fmt"
	"math"
)

func triangular(n int) int {
	return n * (n + 1) / 2
}

func isTriangular(x int) bool {
	//x = Tn = n(n+1)/2 => n^2 +n - 2x = 0
	return quadratic(1, 1, -2*x)
}

func isSquare(x int) bool {
	s := int(math.Sqrt(float64(x)))
	return s*s == x
}

func quadratic(a int, b int, c int) bool {
	d := b*b - 4*a*c
	if !isSquare(d) {
		return false
	}
	return (-b+int(math.Sqrt(float64(d))))%(2*a) == 0
}

func isPentagonal(x int) bool {
	//x = Pn =n(3n-1)/2 => 3n^2 - n - 2x = 0
	return quadratic(3, -1, -2*x)
}

func isHexagonal(x int) bool {
	//x = Hn =n(2n-1) => 2n^2 - n - x = 0
	return quadratic(2, -1, -x)
}

func isHeptagonal(x int) bool {
	//x = Hn =n(5n-3)/2 => 5n^2 - 3n - 2x = 0
	return quadratic(5, -3, -2*x)
}

func isOctagonal(x int) bool {
	//x = On =n(3n-2) => 3n^2 - 2n - x = 0
	return quadratic(3, -2, -x)
}

func isFiguralNumber(x int) bool {
	return isSquare(x) || isPentagonal(x) || isHexagonal(x) || isHeptagonal(x) || isOctagonal(x)
}

type check func(x int) bool

func has(a []int, fn check) bool {
	cnt := 0
	for i := 0; i < len(a); i++ {
		if fn(a[i]) {
			cnt++
		}
	}
	/*
		for _, r := range a {
			if fn(r) {
				// return true
				cnt++
			}
		}
	*/
	return cnt == 1
	// return false
}

func contains(a []int, x int) bool {
	for _, r := range a {
		if r == x {
			return true
		}
	}
	return false
}

func find(start int, pos int, current []int) (bool, []int) {
	c := current[pos-1]
	next := c % 100
	if next < 10 {
		return false, []int{}
	}
	if pos == 6 {
		if next == start/100 {
			if /*has(current, isTriangular) &&*/ has(current, isSquare) && has(current, isPentagonal) && has(current, isHexagonal) && has(current, isHeptagonal) && has(current, isOctagonal) {
				return true, current
			}
		}
		return false, []int{}
	}
	for i := 10; i < 100; i++ {
		p := next*100 + i
		if isFiguralNumber(p) && !contains(current, p) {
			found, result := find(start, pos+1, append(current, p))
			if found {
				return true, result
			}
		}
	}
	return false, []int{}
}

func main() {
	n := 1
	for {
		t := triangular(n)
		if t > 10000 {
			break
		}
		if t >= 1000 {
			found, result := find(t, 1, []int{t})
			if found {
				fmt.Println(result)
				sum := 0
				for _, r := range result {
					fmt.Printf("%d - %t,%t,%t,%t,%t,%t\n", r, isTriangular(r), isSquare(r), isPentagonal(r), isHexagonal(r), isHeptagonal(r), isOctagonal(r))
					sum += r
				}
				fmt.Println(sum)
				// break
			}
		}
		n++
	}
}
