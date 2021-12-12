package main

import (
	"fmt"
)

const maxLen = 16
const maxDigits = 65536 // 2^16

func main() {
	m := [maxLen + 1][maxDigits]int64{}
	for i := 1; i < 16; i++ {
		m[1][1<<i] = 1
	}
	for i := 2; i <= maxLen; i++ {
		for d := 0; d < maxDigits; d++ {
			for j := 0; j < 16; j++ {
				m[i][d|(1<<j)] += m[i-1][d]
			}
		}
	}
	mask := (1 << 0) + (1 << 1) + (1 << 0xA)
	sum := int64(0)
	for l := 1; l <= maxLen; l++ {
		for d := 0; d < maxDigits; d++ {
			if d&mask == mask {
				sum += m[l][d]
			}
		}
	}
	fmt.Printf("%X\n", sum)
}
