package main

import (
	"fmt"
	"math"
)

var k = int64(1)
var s = []int64{}

func nextS() int64 {
	var v = int64(0)
	if k <= 55 {
		v = (100003-200003*k+300007*k*k*k)%1000000 - 500000
		s = append(s, v)
	} else {
		v = (s[55-24]+s[0]+1000000)%1000000 - 500000
		s = s[1:]
		s = append(s, v)
	}
	k++
	return v
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubSum(x []int64) int64 {
	best := int64(math.MinInt)
	current := int64(0)
	for i := 0; i < len(x); i++ {
		current = max(x[i], current+x[i])
		best = max(best, current)
	}
	return best
}

const N = 2000

func main() {
	m := [N][N]int64{}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			m[y][x] = nextS()
		}
	}
	best := int64(math.MinInt)
	// horizontal lines
	for i := 0; i < N; i++ {
		best = max(best, maxSubSum(m[i][:]))
	}
	// vertical lines
	for i := 0; i < N; i++ {
		x := []int64{}
		for j := 0; j < N; j++ {
			x = append(x, m[j][i])
		}
		best = max(best, maxSubSum(x))
	}
	// diagonal lines
	for i := 0; i < N; i++ {
		x := []int64{}
		// above main diagonal
		for j := i; j < N; j++ {
			x = append(x, m[i][j])
		}
		best = max(best, maxSubSum(x))
		// below main diagonal
		x = []int64{}
		for j := i; j < N; j++ {
			x = append(x, m[j][i])
		}
		best = max(best, maxSubSum(x))
	}
	// anti diagonal
	for i := 0; i < N; i++ {
		x := []int64{}
		for j := N - i - 1; j >= 0; j-- {
			x = append(x, m[i][j])
		}
		best = max(best, maxSubSum(x))
		x = []int64{}
		for j := N - i - 1; j >= 0; j-- {
			x = append(x, m[j][i])
		}
		best = max(best, maxSubSum(x))
	}

	fmt.Println(best)
}
