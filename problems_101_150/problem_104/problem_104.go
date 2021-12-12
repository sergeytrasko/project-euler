package main

import (
	"fmt"
	"math/big"
)

func isPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	digits := [10]int{}
	for i := 0; i < len(s); i++ {
		digits[int(s[i]-byte('0'))]++
	}
	for i := 1; i <= 9; i++ {
		if digits[i] != 1 {
			return false
		}
	}
	return true
}

func main() {
	cutoff := big.NewInt(1_000_000_000)
	a, b := big.NewInt(1), big.NewInt(1)
	k := 3
	for {
		c := new(big.Int).Add(a, b)
		a, b = b, c
		if k%10000 == 0 {
			fmt.Printf("%d\n", k)
		}
		if c.Cmp(cutoff) > 0 {
			last := new(big.Int).Set(c)
			lastDigits := last.Mod(last, cutoff).String()
			if isPandigital(lastDigits) {
				first := c.String()[0:9]
				if isPandigital(first) {
					fmt.Printf("%d - %d\n", k, c)
					break
				}
			}
		}
		k++
	}
	fmt.Println(k)
}
