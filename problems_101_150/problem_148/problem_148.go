package main

import (
	"fmt"
)

func explore() int {
	const max = 7*7*2 + 1
	c := [max + 1][max + 1]int64{}
	c[0][0] = 1
	cnt := 0
	for n := 1; n < max; n++ {
		if (n-1)%7 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d\t", n-1)
		r := 0
		for k := 1; k <= n; k++ {
			c[n][k] = (c[n-1][k] + c[n-1][k-1]) % 7
			v := "."
			if c[n][k]%7 != 0 {
				v = "*"
				cnt++
				r++
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
	return cnt
}

func countRecursive(row int) int64 {
	d, r := 1, row
	for r >= 7 {
		r /= 7
		d *= 7
	}

	if row < 7 {
		return int64(row + 1)
	} else {
		mul := int64((row / d) + 1)
		res := mul * countRecursive(row%d)
		return res
	}
}

func count(row int) int64 {
	d := int64(1)
	for row > 0 {
		d *= int64(row%7 + 1)
		row /= 7
	}
	return d
}

func main() {
	// explore()
	cnt := int64(0)
	for i := 0; i < 1e9; i++ {
		cnt += count(i)
	}
	fmt.Println(cnt)
}
