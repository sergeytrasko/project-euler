package main

import (
	"fmt"
	"math"
	"strconv"
)

func numberOfNumbers(len int) int {
	if len == 1 {
		return 9
	}
	return int(math.Pow10(len)-math.Pow10(len-1)) - 1
}

func d(n int) int {
	p := 0
	l := 1
	for {
		nr := numberOfNumbers(l)
		if p+nr*l > n {
			// fmt.Printf("%d should be between %d and %d\n", n, p, p+nr*l)
			break
		}
		p += nr * l
		l++
	}
	first := int(math.Pow10(l - 1))
	// fmt.Printf("n=%d, p=%d, l=%d, first=%d\n", n, p, l, first)
	number := first + (n-p)/l
	digitPos := (n - p) % l
	// fmt.Printf("number=%d, digitPos=%d\n", number, digitPos)
	digit, _ := strconv.Atoi(strconv.Itoa(number)[digitPos : digitPos+1])
	return digit
}

func dd(s string, p int) int {
	d := s[p-1 : p]
	digit, _ := strconv.Atoi(d)
	return digit
}

func main() {
	s := ""
	for i := 1; len(s) < 1000000; i++ {
		s = s + strconv.Itoa(i)
	}
	fmt.Println(dd(s, 1) * dd(s, 10) * dd(s, 100) * dd(s, 1000) * dd(s, 10000) * dd(s, 100000) * dd(s, 1000000))

	// for i := 1; i < 200; i++ {
	// fmt.Printf("%d -> %d, %s\n", i, d(i-1), s[i-1:i])
	// }
	// fmt.Println(numberOfNumbers(3))
	// fmt.Println(d(12))
	// fmt.Println(d(1) * d(10) * d(100) * d(1000) * d(10000) * d(100000) * d(1000000))
}
