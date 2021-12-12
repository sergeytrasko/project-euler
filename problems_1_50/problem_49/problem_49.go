package main

import "fmt"

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func sortDigits(n int) string {
	digits := [10]int{}
	for n > 0 {
		digits[n%10]++
		n = n / 10
	}
	s := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			s = s + fmt.Sprintf("%d", i)
		}
	}
	return s
}

func sameDigits(a int, b int, c int) bool {
	d := sortDigits(a)
	return d == sortDigits(b) && d == sortDigits(c)
}

func main() {
	for n := 1001; n < 9999; n += 2 {
		if isPrime(n) && n != 1487 {
			for d := 2; d < (9999-n)/2; d += 2 {
				if isPrime(n+d) && isPrime(n+2*d) && sameDigits(n, n+d, n+2*d) {
					fmt.Printf("%d%d%d\n", n, n+d, n+2*d)
				}
			}
		}
	}
}
