package main

import "fmt"

func main() {
	for a := 1; a < 1000; a = a + 1 {
		for b := a + 1; b < 1000; b = b + 1 {
			c := 1000 - a - b
			if c <= 0 {
				continue
			}
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
			}
		}
	}
}
