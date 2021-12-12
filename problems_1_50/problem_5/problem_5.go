package main

import "fmt"

func divisbleByBefore(n int, before int) bool {
	for i := 2; i <= before; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 1
	for {
		if divisbleByBefore(n, 20) {
			fmt.Println(n)
			break
		}
		n = n + 1
	}
}
