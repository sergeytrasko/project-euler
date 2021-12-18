package main

import "fmt"

func main() {
	a, b := int64(1), int64(1)
	for i := int64(41); i <= 50; i++ {
		a *= i
	}
	for i := int64(61); i <= 70; i++ {
		b *= i
	}
	fmt.Printf("%.9f\n", 7*(1-float64(a)/float64(b)))
}
