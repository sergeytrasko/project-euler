package main

import "fmt"

const max = 50

func main() {
	m := [max + 1]int{}
	m[0] = 1
	for i := 1; i <= max; i++ {
		for l := 1; l <= 4; l++ {
			if i >= l {
				m[i] += m[i-l]
			}
		}
		/*
			m[i] += m[i-1] //skipping the cell
			if i >= 2 {
				m[i] += m[i-2] //red block of 2
			}
			if i >= 3 {
				m[i] += m[i-3] //green block of 3
			}
			if i >= 4 {
				m[i] += m[i-4] //blue block of 4
			}
		*/
	}
	fmt.Println(m[max])
}
