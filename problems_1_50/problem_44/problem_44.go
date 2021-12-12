package main

import (
	"fmt"
	"math"
)

func pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func isPerfectSquare(x int) (int, bool) {
	s := int(math.Sqrt(float64(x)))
	return s, s*s == x
}

func isPentagonal(x int) (int, bool) {
	// x = Pn = n(3n - 1) / 2
	// 3n^2 - n - 2x = 0
	// n = (1+- sqrt(1 + 4 * 2x * 3)) / 6 = (1 + sqrt(1 + 24x)) / 6 - should be integer
	d := 1 + 24*x
	s, isSquare := isPerfectSquare(d)
	if !isSquare {
		return 0, false
	}
	return (1 + s) / 6, (1+s)%6 == 0
}

func main() {
	// fmt.Println(pentagonal(4))
	// fmt.Println(pentagonal(7))
	// fmt.Println(isPentagonal(pentagonal(4) + pentagonal(7)))
	k := 1
	maxJ := 10000
	bestDif := math.MaxInt
	for k < 10000 {
		pk := pentagonal(k)
		for j := 1; j < maxJ; j++ {
			pj := pentagonal(k + j)
			sum, sumOk := isPentagonal(pj + pk)
			dif, difOk := isPentagonal(pj - pk)
			if sumOk && difOk {
				fmt.Printf("k=%d, j=%d, Pk=%d, Pj=%d, P%d-P%d=P%d, P%d+P%d=P%d\n", k, k+j, pk, pj, k, k+j, dif, k, k+j, sum)
				if pj-pk < bestDif {
					bestDif = pj - pk
				}
			}
		}
		k++
	}
	fmt.Println(bestDif)
}
