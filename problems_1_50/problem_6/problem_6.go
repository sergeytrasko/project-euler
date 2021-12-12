package main

import "fmt"

func squareSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i*i
	}
	return sum
}

func sumSquare(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum * sum
}

func main() {
	fmt.Println(sumSquare(100) - squareSum(100))
}
