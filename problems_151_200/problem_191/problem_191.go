package main

import (
	"fmt"
	"strings"
)

func countNaive(s string, left int, count *int) {
	if strings.Contains(s, "AAA") || strings.Count(s, "L") > 1 {
		return
	}
	if left == 0 {
		*count = *count + 1
		return
	}
	countNaive(s+"A", left-1, count)
	countNaive(s+"O", left-1, count)
	countNaive(s+"L", left-1, count)
}

func countNaive2(left int, absentInARow int, late int) int {
	if absentInARow == 3 || late > 1 {
		return 0
	}
	if left == 0 {
		return 1
	}
	return countNaive2(left-1, 0, late) + //on-time
		countNaive2(left-1, absentInARow+1, late) + //absent
		countNaive2(left-1, 0, late+1) //late
}

func main() {
	// cnt := 0
	// countNaive("", 10, &cnt)
	// fmt.Println(cnt)
	fmt.Println(countNaive2(30, 0, 0))
}
