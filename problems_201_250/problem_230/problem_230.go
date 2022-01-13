package main

import "fmt"

func naive(a, b string, n int64) string {
	for int64(len(b)) < n {
		a, b = b, a+b
	}
	return b[n-1 : n]
}

func solve(a0, b0 string, n int64) string {
	target := a0 + b0
	a, b := a0, b0
	f := []int64{int64(len(a)), int64(len(b))}
	for i := 2; f[i-1] < n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}
	i := len(f) - 1
	for n > int64(len(target)) {
		left, _ := f[i-1], f[i-2]
		if n > left {
			n -= left
			i -= 2
		} else {
			i--
		}
	}
	if i%2 == 0 {
		return (a0 + b0)[n-1 : n]
	} else {
		return (b0 + a0)[n-1 : n]
	}
}

func main() {
	// fmt.Println(naive("1415926535", "8979323846", 35))
	// fmt.Println(solve("1415926535", "8979323846", 35))
	a := "1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	b := "8214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196"
	seven := int64(1)
	s := ""
	for n := int64(0); n <= 17; n++ {
		p := (127 + 19*n) * seven
		s = solve(a, b, p) + s
		fmt.Println(s)
		seven *= 7
	}
	fmt.Println(s)
}
