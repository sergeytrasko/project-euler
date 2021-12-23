package main

import (
	"fmt"
)

func divides(x int) bool {
	a, b, c := 1, 1, 1
	for n := 4; ; n++ {
		d := (a + b + c)%x
		if d == 0 {
			return true
		}
		a, b, c = b, c, d
		if a == 1 && b == 1 && c == 1 {
			return false
		}
	}
	return false
}

func main() {
	cnt:= 0
	for n := 3; cnt < 124; n+=2{
		if !divides(n) {
			cnt++
			fmt.Printf("%d - %d\n", cnt, n)
		}
	}
}
