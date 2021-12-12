package main

import "fmt"

func divisorSum(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	sum, n := 0, 1
	for n < 10000 {
		m := divisorSum(n)
		if n > m && divisorSum(m) == n {
			sum += (m + n)
			fmt.Printf("%d, %d\n", n, m)
		}
		n++
	}
	fmt.Println(sum)
}
