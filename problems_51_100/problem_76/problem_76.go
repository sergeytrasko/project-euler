package main

import "fmt"

const max = 100

func main() {
	p := [max + 1][max + 1]int{}
	for i := 0; i < max; i++ {
		p[i][0] = 0
	}
	p[0][0] = 1
	for n := 0; n <= max; n++ {
		for k := 0; k <= max; k++ {
			if k > 0 && k <= n {
				p[n][k] = p[n][k-1] + p[n-k][k]
			} else if k > n {
				p[n][k] = p[n][n]
			}
		}
	}
	// for _, row := range p {
	// fmt.Println(row)
	// }
	fmt.Println(p[max][max] - 1) //exclude
}
