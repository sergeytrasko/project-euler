package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/oddg/hungarian-algorithm"
)

const n = 15

func read() [][]int {
	m := make([][]int, n)
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for strings.Contains(line, "  ") {
			line = strings.ReplaceAll(strings.TrimSpace(line), "  ", " ")
		}
		tokens := strings.Split(line, " ")
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			num, _ := strconv.Atoi(tokens[j])
			m[i][j] = num
		}
		i = i + 1
	}
	return m
}

func convertToMax(m *[][]int) {
	max := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if (*m)[r][c] > max {
				max = (*m)[r][c]
			}
		}
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			(*m)[r][c] = max - (*m)[r][c]
		}
	}
}

func main() {
	m := read()
	convertToMax(&m)
	solution, _ := hungarianAlgorithm.Solve(m)
	fmt.Println(solution)
	m = read()
	sum := 0
	for r := 0; r < n; r++ {
		sum += m[r][solution[r]]
	}
	fmt.Println(sum)
}
