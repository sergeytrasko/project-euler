package main

import (
	"fmt"
	"strconv"
)

func digitsPowers(n int) int {
	s := strconv.Itoa(n)
	sum := 0
	for i := 0; i < len(s); i++ {
		x, _ := strconv.Atoi(s[i : i+1])
		sum += x * x * x * x * x
	}
	return sum
}

func main() {
	max := 10000000 // there should be some reasonable limit
	sum := 0
	for n := 2; n < max; n++ {
		pow := digitsPowers(n)
		if n == pow {
			// fmt.Println(n)
			sum += n
		}
	}
	fmt.Println(sum)
}
