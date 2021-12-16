package main

import "fmt"

func g(k int) int {
	return k * (3*k - 1) / 2
}

func main() {
	const max = 100000
	const mod = 1_000_000
	p := [max]int{}
	p[0] = 1
	n := 1
	for {
		sum := 0
		k := 1
		for {
			term := n - g(k)
			if term < 0 {
				break
			}
			sign := 1
			if k%2 == 0 {
				sign = -1
			}
			sum = (sum + sign*p[term] + mod) % mod
			k = -k
			if k > 0 {
				k++
			}
		}
		p[n] = sum
		// fmt.Printf("P(%d) mod %d = %d\n", n, mod, sum)
		if sum == 0 {
			break
		}
		n++
	}
	fmt.Println(n)
}
