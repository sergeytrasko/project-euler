package main

import "fmt"

const N = 50

func calc(remaining int, blockLen int, cache *[]int) int {
	cnt := 0
	if remaining < blockLen {
		return cnt
	}
	if (*cache)[remaining] > 0 {
		return (*cache)[remaining]
	}
	for p := 0; p <= remaining-blockLen; p++ {
		cnt++
		cnt += calc(remaining-p-blockLen, blockLen, cache)
	}
	(*cache)[remaining] = cnt
	return cnt
}

func count(blockLen int) int {
	cache := make([]int, N+1)
	return calc(N, blockLen, &cache)
}

func main() {
	fmt.Println(count(2) + count(3) + count(4))
}
