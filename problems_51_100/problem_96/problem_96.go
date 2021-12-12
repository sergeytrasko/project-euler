package main

import (
	"bufio"
	"fmt"
	"os"
)

func canPut(field [][]int, x int, y int, value int) bool {
	return !alreadyInVertical(field, x, y, value) &&
		!alreadyInHorizontal(field, x, y, value) &&
		!alreadyInSquare(field, x, y, value)
}

func alreadyInVertical(field [][]int, x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(field [][]int, x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(field [][]int, x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if field[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int, field *[][]int) (bool, int) {
	if y == 9 {
		return true, (*field)[0][0]*100 + (*field)[0][1]*10 + (*field)[0][2]
	}
	if (*field)[y][x] != 0 {
		nx, ny := next(x, y)
		return solve(nx, ny, field)
	} else {
		for i := range [9]int{} {
			var v = i + 1
			if canPut(*field, x, y, v) {
				(*field)[y][x] = v
				nx, ny := next(x, y)
				if r, n := solve(nx, ny, field); r {
					return r, n
				}
				(*field)[y][x] = 0
			}
		}
		return false, 0
	}
}

func main() {
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		field := [][]int{}
		for i := 0; i < 9; i++ {
			scanner.Scan()
			line = scanner.Text()
			row := []int{}
			for j := 0; j < 9; j++ {
				row = append(row, int(line[j]-byte('0')))
			}
			field = append(field, row)
		}
		// fmt.Println(field)
		_, n := solve(0, 0, &field)
		sum += n
		// fmt.Println(n)
	}
	fmt.Println(sum)
}
