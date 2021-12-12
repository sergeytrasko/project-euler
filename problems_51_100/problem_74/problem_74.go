package main

import "fmt"

func fact(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

func digitsFactorial(n int) int {
	sum := 0
	for n > 0 {
		sum += fact(n % 10)
		n = n / 10
	}
	return sum
}

func chainLength(n int) int {
	visited := map[int]int{}
	visited[n] = 1
	len := 1
	for {
		n = digitsFactorial(n)
		if visited[n] > 0 {
			break
		}
		len++
		visited[n] = len
	}
	return len
}

func main() {
	cnt := 0
	for i := 1; i < 1_000_000; i++ {
		l := chainLength(i)
		if l == 60 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
