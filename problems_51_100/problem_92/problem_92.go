package main

import "fmt"

// const maxDigitSum = 9 * 9 * 9 //9 digits max, each is 9 max - next number will be less than 9 * 9^2
const max = 10_000_000

var cache = [max]int{}

func digitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		x := n % 10
		sum += x * x
		n = n / 10
	}
	return sum
}

func chain(start int) int {
	n := start
	// fmt.Print(n)
	for n != 1 && n != 89 {
		// fmt.Print("->")
		n = digitSquareSum(n)
		// fmt.Print(n)
		if cache[n] != 0 {
			// fmt.Printf("(known to coverge to %d)", cache[n])
			n = cache[n]
		}
	}
	// fmt.Println()
	cache[start] = n
	return n
}

func main() {
	cnt := 0
	for n := 1; n < max; n++ {
		res := chain(n)
		if res == 89 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
