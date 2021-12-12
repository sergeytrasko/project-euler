package main

import "fmt"

var values = []int{0, 1, 2, 5, 10, 20, 50, 100, 200}

const max = 200
const vl = 9

func main() {
	p := [max + 1][vl + 1]int{}
	for n := 1; n <= max; n++ {
		for k := 1; k < vl; k++ {
			if n == 1 || k == 1 {
				p[n][k] = 1
			}
			if n > values[k] {
				p[n][k] = p[n][k-1] + p[n-values[k]][k]
			} else if n < values[k] {
				p[n][k] = p[n][k-1]
			} else if n == values[k] {
				p[n][k] = p[n][k-1] + 1
			}
		}
	}
	// for _, row := range p {
	// 	fmt.Println(row)
	// }
	fmt.Println(p[max][vl-1])
}
