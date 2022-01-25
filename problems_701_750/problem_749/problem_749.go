package main

import (
	"fmt"
)

func power(a int, b int) int64 {
	r := int64(1)
	for b > 0 {
		r *= int64(a)
		b--
	}
	return r
}

func digitSumPower(digits []int, k int) int64 {
	s := int64(0)
	for i := 0; i < len(digits); i++ {
		s += power(digits[i], k)
	}
	return s
}

func toDigits(n int64) []int {
	d := []int{}
	for n > 0 {
		d = append(d, int(n%10))
		n /= 10
	}
	return d
}

func explore() {
	sum := int64(0)
	for n := int64(1); n < 1_000_000; n++ {
		digits := toDigits(n)
		for k := 2; k < 10; k++ {
			s := digitSumPower(digits, k)
			if s == n-1 || s == n+1 {
				fmt.Printf("%d - %d\n", n, k)
				sum += n
				break
			}
		}
	}
	fmt.Println(sum)
}

func isNearPowerSum(d digits, n int64) bool {
	for n > 0 {
		x := n % 10
		if d[x] > 0 {
			d[x]--
		} else {
			return false
		}
		n /= 10
	}
	for i := 1; i <= 9; i++ {
		if d[i] > 0 {
			return false
		}
	}
	return true
}

type digits = [10]int

func digitSumPower2(d digits, k int) int64 {
	sum := int64(0)
	for i := 0; i < len(d); i++ {
		sum += int64(d[i]) * power(i, k)
	}
	return sum
}

func split(current, target, usedDigits, index int, d digits, n int) int64 {
	if current > target {
		return 0
	}
	if current == target {
		s := int64(0)
		d0 := d[0]
		d[0] = n - usedDigits // fill as many 0 as possible
		for k := 2; k <= n; k++ {
			dps := digitSumPower2(d, k)
			if isNearPowerSum(d, dps-1) {
				fmt.Printf("%d - %d, k=%d\n", d, dps-1, k)
				s += dps - 1
			}
			if isNearPowerSum(d, dps+1) {
				fmt.Printf("%d - %d, k=%d\n", d, dps+1, k)
				s += dps + 1
			}
		}
		d[0] = d0
		return s
	}
	if index >= len(d) {
		return 0
	}
	s := int64(0)
	for i := 0; i <= n-usedDigits; i++ {
		d[index] = i
		s += split(current+index*i, target, usedDigits+i, index+1, d, n)
		d[index] = 0
	}
	return s
}

func main() {
	// explore()
	const n = 16
	sum := int64(0)
	for i := 1; i <= n*9; i++ {
		sum += split(0, i, 0, 1, digits{}, n)
	}
	fmt.Println(sum)
}
