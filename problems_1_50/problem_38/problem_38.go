package main

import (
	"fmt"
	"strconv"
)

func is1_9Pandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	d := [10]int{}
	for i := 0; i < len(s); i++ {
		x := int(s[i] - byte('0'))
		d[x]++
	}
	if d[0] > 0 {
		return false
	}
	for i := 1; i < len(d); i++ {
		if d[i] != 1 {
			return false
		}
	}
	return true
}

func try(k int) (int, int) {
	s := ""
	for n := 1; n < 9; n++ {
		s = s + fmt.Sprintf("%d", k*n)
		if n > 1 && len(s) == 9 && is1_9Pandigital(s) {
			r, _ := strconv.Atoi(s)
			return r, n
		} else if len(s) > 9 {
			return -1, -1
		}
	}
	return -1, -1
}

func main() {
	k := 1
	max := 0
	bestK := 0
	for k < 10000 {
		r, n := try(k)
		if r > max {
			max = r
			bestK = k
			fmt.Printf("for k=%d and n=%d: 1 to 9 pandigital number found - %d\n", k, n, r)
		}
		k++
	}
	fmt.Println(bestK)
	fmt.Println(max)
}
