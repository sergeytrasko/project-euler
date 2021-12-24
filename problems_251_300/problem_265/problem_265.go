package main

import (
	"fmt"
)

func count(current int, value int, l int, target int, bits int, visited *[]bool, sum *int) {
	if l == target {
		if current == 0 {
			*sum = *sum + (value >> bits)
		}
		return
	}
	x := ((current << 1) & ((1 << bits) - 1))
	if !(*visited)[x] {
		(*visited)[x] = true
		count(x, value*2, l+1, target, bits, visited, sum)
		(*visited)[x] = false
	}
	if !(*visited)[x+1] {
		(*visited)[x+1] = true
		count(x+1, value*2+1, l+1, target, bits, visited, sum)
		(*visited)[x+1] = false
	}
	if x == 0 && l == target-1 {
		count(x, value*2, l+1, target, bits, visited, sum)
	}
}

func s(n int) int {
	sum := 0
	target := 1
	for i := 1; i <= n; i++ {
		target *= 2
	}
	visited := make([]bool, target)
	visited[0] = true
	count(0, 0, 0, target, n, &visited, &sum)
	return sum
}

func main() {
	// fmt.Println(s(3))
	fmt.Println(s(5))
}
