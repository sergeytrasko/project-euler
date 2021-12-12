package main

import "fmt"

func main() {
	a, b := 0, 1
	var sum = 0
	for a < 4000000 {
		c := a + b
		a, b = b, c
		if a%2 == 0 {
			sum = sum + a
		}
	}
	fmt.Println(sum)
}
