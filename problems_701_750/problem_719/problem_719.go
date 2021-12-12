package main

import (
	"fmt"
)

func canSplit(n int64, target int64, sum int64, groups int) bool {
	// fmt.Printf("n=%d, target=%d, sum=%d, groups=%d\n", n, target, sum, groups)
	if sum > target {
		return false
	}
	if sum == target && n == 0 && groups > 1 {
		return true
	}
	grp := int64(0)
	f := int64(1)
	for n > 0 {
		grp = grp + f*(n%10)
		f *= 10
		n /= 10
		if canSplit(n, target, sum+grp, groups+1) {
			return true
		}
	}
	return false
}

func isSNumber(n int64, nn int64) bool {
	return canSplit(n, nn, 0, 0)
}

func t(n int64) int64 {
	sum := int64(0)
	for i := int64(1); i*i <= n; i++ {
		if isSNumber(i*i, i) {
			// fmt.Printf("%d=sqrt(%d)\n", i*i, i)
			sum += i * i
		}
	}
	return sum
}

func main() {
	// fmt.Println(t(1e4)) //should be 41333
	fmt.Println(t(1e12))
}
