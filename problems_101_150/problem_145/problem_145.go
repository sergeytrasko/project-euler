package main

import "fmt"

func reverse(n int64) int64 {
	x := int64(0)
	for n > 0 {
		x = x*10 + n%10
		n /= 10
	}
	return x
}

func onlyOddDigits(n int64) bool {
	for n > 0 {
		if n%2 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func main() {
	// n := 409
	// fmt.Println(reverse(n))
	// fmt.Println(onlyOddDigits(n + reverse(n)))
	cnt := 0
	for n := int64(1); n < 1_000_000_000; n++ {
		if n%10 == 0 {
			continue
		}
		if onlyOddDigits(n + reverse(n)) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
