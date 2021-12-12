package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sign(x1 int, y1 int, x2 int, y2 int, px int, py int) int {
	return (x1-px)*(y2-y1) - (x2-x1)*(y1-py)
}

func containsOrigin(coords []int) bool {
	x1, y1, x2, y2, x3, y3 := coords[0], coords[1], coords[2], coords[3], coords[4], coords[5]
	a := sign(x1, y1, x2, y2, 0, 0)
	b := sign(x2, y2, x3, y3, 0, 0)
	c := sign(x3, y3, x1, y1, 0, 0)

	return (a > 0 && b > 0 && c > 0) || (a < 0 && b < 0 && c < 0)
}

func main() {
	cnt := 0
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		coords := []int{}
		for i := 0; i < 6; i++ {
			c, _ := strconv.Atoi(tokens[i])
			coords = append(coords, c)
		}
		if containsOrigin(coords) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
