package main

import (
	"fmt"
	"math/big"
)

func check(n int64) bool {
	n2 := big.NewInt(n)
	n2 = n2.Mul(n2, n2)
	s := n2.String()
	if len(s) == 19 {
		ok := true
		for i := 0; i < 9; i++ {
			if int(s[2*i]-byte('0')) != i+1 {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func main() {
	//n^2 = 1_2_3_4_5_6_7_8_9_0
	//(n*10)^ = 1_2_3_4_5_6_7_8_900 => n ends with ****30 or ****70
	//n^2 - 17 digis + 900 => n is 8 digits
	lastTwoDigits := []int{30, 70}
	for n := int64(10_000_000); n < int64(100_000_000-1); n++ {
		for i := 0; i < len(lastTwoDigits); i++ {
			x := int64(n)*100 + int64(lastTwoDigits[i])
			if check(x) {
				fmt.Println(x)
				return
			}
		}
	}
}
