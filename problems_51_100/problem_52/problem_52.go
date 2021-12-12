package main

import "fmt"

func normalizeDigits(n int) string {
	digits := [10]int{}
	for n > 0 {
		x := n % 10
		digits[x]++
		n = n / 10
	}
	s := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			s += fmt.Sprintf("%d", i)
		}
	}
	return s
}

func main() {
	n := 1
	for {
		d := normalizeDigits(n)
		if d == normalizeDigits(n*2) && d == normalizeDigits(n*3) && d == normalizeDigits(n*4) && d == normalizeDigits(n*5) && d == normalizeDigits(n*6) {
			break
		}
		n++
	}
	fmt.Println(n)
}
