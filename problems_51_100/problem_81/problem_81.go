package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const max = 80

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	m := [max][max]int{}
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		for j := 0; j < max; j++ {
			num, _ := strconv.Atoi(tokens[j])
			m[i][j] = num
		}
		i += 1
	}

	for i := 1; i < max; i++ {
		m[i][0] = m[i][0] + m[i-1][0]
		m[0][i] = m[0][i] + m[0][i-1]
	}
	for i := 1; i < max; i++ {
		for j := 1; j < max; j++ {
			m[i][j] = m[i][j] + min(m[i][j-1], m[i-1][j])
		}
	}
	fmt.Println(m[max-1][max-1])

}
