package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const n = 15

func main() {
	m := [n][n]int{}
	file, _ := os.Open("in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		fmt.Println(tokens)
		for j := 0; j < i+1; j++ {
			num, _ := strconv.Atoi(tokens[j])
			m[i][j] = num
		}
		i = i + 1
	}
	file.Close()

	for r := 1; r < n; r++ {
		for j := 0; j <= r; j++ {
			mx := m[r-1][j]
			if j > 0 && (m[r-1][j-1] > m[r-1][j]) {
				mx = m[r-1][j-1]
			}
			m[r][j] = m[r][j] + mx
		}
	}
	for _, r := range m {
		fmt.Println(r)
	}
	max := 0
	for j := 0; j < n; j++ {
		if m[n-1][j] > max {
			max = m[n-1][j]
		}
	}
	fmt.Println(max)
}
