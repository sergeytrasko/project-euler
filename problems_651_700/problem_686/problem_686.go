package main

import (
	"fmt"
	"math"
)

const lg10_2 = math.Ln2 / math.Ln10

func p(l, n int64) int64 {
	x := l
	lMul := 1
	for x > 10 {
		lMul *= 10
		x /= 10
	}
	cnt := int64(0)
	for j := int64(1); ; j++ {
		numberOfDigits := float64(j) * lg10_2
		x := math.Mod(numberOfDigits, 1)
		e := int64(math.Trunc(math.Exp(math.Ln10*x) * float64(lMul)))
		if e == l {
			if cnt%25000 == 0 {
				fmt.Printf("%d, %d -> %d\n", j, cnt, e)
			}
			cnt++
			if cnt == n {
				return j
			}
		}
	}

}

func main() {
	fmt.Println(p(123, 678910))
}
