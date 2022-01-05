package main

import "fmt"

func main() {
	loses := 0
	for n := int64(1); n <= 1<<30; n++ {
		res := n ^ (2 * n) ^ (3 * n)
		if res == 0 {
			loses++
		}
	}
	fmt.Println(loses)
}
