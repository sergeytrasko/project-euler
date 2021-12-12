package main

import "fmt"

const minLen = 3
const N = 50

var cache = [N + 1]int{}

func calc(remaining int) int {
	cnt := 1
	if remaining < minLen {
		return cnt
	}
	if cache[remaining] > 0 {
		return cache[remaining]
	}
	for p := 0; p <= remaining-minLen; p++ {
		for block := minLen; block <= remaining-p; block++ {
			cnt += calc(remaining - p - block - 1)
		}
	}
	cache[remaining] = cnt
	return cnt
}

func main() {
	fmt.Println(calc(50))
}
