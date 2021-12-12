package main

import "fmt"

func countDivisors(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt = cnt + 1
		}
	}
	return cnt
}

func triangNumberDivisors(p int) int {
	if p%2 == 0 {
		return countDivisors(p/2) * countDivisors(p+1)
	} else {
		return countDivisors(p) * countDivisors((p+1)/2)
	}
}

func main() {
	best := 0
	n := 1
	for {
		cnt := triangNumberDivisors(n)
		if cnt > best {
			fmt.Printf("T%d has %d divisors\n", n, cnt)
			best = cnt
		}
		if cnt > 500 {
			break
		}
		n++
	}
	fmt.Println(n * (n + 1) / 2)
}
