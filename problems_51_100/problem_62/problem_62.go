package main

import (
	"fmt"
	"math/big"
)

func cube(n int) string {
	nn := big.NewInt(int64(n))
	return nn.Exp(nn, big.NewInt(int64(3)), nil).String()
}

func cubeWithSortedDigits(n int) string {
	s := cube(n)
	digits := [10]int{}
	for i := 0; i < len(s); i++ {
		digits[int(s[i]-byte('0'))]++
	}
	ss := ""
	for i := 0; i < len(digits); i++ {
		for j := 0; j < digits[i]; j++ {
			ss += fmt.Sprintf("%d", i)
		}
	}
	return ss
}

func main() {
	m := map[string][]int{}
	n := 1
	for {
		s := cubeWithSortedDigits(n)
		if m[s] == nil {
			m[s] = []int{n}
		} else {
			l := m[s]
			l = append(l, n)
			m[s] = l
			if len(l) == 5 {
				// fmt.Println(l)
				fmt.Println(cube(l[0]))
				break
			}
		}
		n++
	}
}
