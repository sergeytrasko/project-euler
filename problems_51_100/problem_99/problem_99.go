package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calc(base int, exp int) float64 {
	return float64(exp) * math.Log10(float64(base))
}

func main() {
	file, _ := os.Open("in")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lineNumber := 0
	best := 0.0
	bestLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		base, _ := strconv.Atoi(tokens[0])
		exp, _ := strconv.Atoi(tokens[1])
		lineNumber += 1
		c := calc(base, exp)
		if c > best {
			best = c
			bestLine = lineNumber
		}
	}
	fmt.Println(bestLine)

}
