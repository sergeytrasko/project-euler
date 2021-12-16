package main

import "fmt"

const max_t = 1_000_000

func main() {
	L := [max_t + 1]int{}        // L[n] number of t to form lamina with hole of any size
	for n := 1; n < max_t; n++ { //hole size
		t := 0
		k := n
		for {
			t += (k*4 + 4)
			if t > max_t {
				break
			}
			L[t]++
			k += 2
		}
	}
	cnt := 0
	for i := 1; i <= max_t; i++ {
		if L[i] >= 1 && L[i] <= 10 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
