package main

import (
	"fmt"
)

func startsWith(n int64, s string) bool {
	for i := 0; i < len(s); i++ {
		if n%3 == 0 {
			if s[i:i+1] != "D" {
				return false
			}
			n = n / 3
		} else if n%3 == 1 {
			if s[i:i+1] != "U" {
				return false
			}
			n = (4*n + 2) / 3
		} else if n%3 == 2 {
			if s[i:i+1] != "d" {
				return false
			}
			n = (2*n - 1) / 3
		}
		if n == 1 {
			break
		}
	}
	return true
}

func main() {
	prefix, n := "UDDDUdddDDUDDddDdDddDDUDDdUUDd", int64(1e15)
	d := int64(1)
	for i := 0; i < len(prefix); i++ {
		s := prefix[:i+1]
		for !startsWith(n, s) {
			n += d
		}
		d *= 3
	}
	fmt.Println(n)
}
