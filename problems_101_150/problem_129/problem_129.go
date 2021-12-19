package main

import "fmt"

func a(n int) int {
	k := 1
	for x := 1; x != 0; x, k = (x*10+1)%n, k+1 {
	}
	return k
}

func main() {
	n := 1_000_000
	for {
		if n%5 != 0 && n%2 != 0 {
			if a(n) > 1_000_000 {
				break
			}
		}
		n++
	}
	fmt.Println(n)
}
