package main

import "fmt"

/*
x^2 - y^2 - z^2 = n
x = z+2a, y = z+a
...
(3a-z)(a+z) = n = pq
a = (p+q)/4
z = (3q-p)/4
*/

const max = 1_000_000

func check(p int, q int) bool {
	aIsNatural := (p+q)%4 == 0
	zIsNatural := 3*q-p > 0 && (3*q-p)%4 == 0
	// fmt.Printf("p=%d, q=%d, a=%f, z=%f\n", p, q, float64(p+q)/4.0, float64(3*q-p)/4.0)
	if aIsNatural && zIsNatural {
		// a := (p + q) / 4
		// z := (3*q - p) / 4
		// fmt.Printf("n=%d, x=%d, y=%d, z=%d\n", p*q, z+2*a, z+a, z)
		return true
	}
	return false
}

func main() {
	solutions := [max + 1]int{}
	for p := 1; p <= max; p++ {
		for q := 1; p*q <= max; q++ {
			n := p * q
			if check(p, q) {
				solutions[n]++
			}
		}
	}
	cnt := 0
	for i := 1; i < max; i++ {
		if solutions[i] == 10 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
