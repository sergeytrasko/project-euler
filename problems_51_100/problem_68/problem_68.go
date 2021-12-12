package main

import (
	"fmt"
	"sort"
)

var solutions = []string{}

func sum(p []int) int {
	sum := 0
	for _, r := range p {
		sum += r
	}
	return sum
}

func checkSolution(p []int) string {
	rows := [][]int{
		{p[0], p[1], p[2]},
		{p[3], p[2], p[4]},
		{p[5], p[4], p[6]},
		{p[7], p[6], p[8]},
		{p[9], p[8], p[1]},
	}
	s := sum(rows[0])
	for i := 1; i < len(rows); i++ {
		if s != sum(rows[i]) {
			return ""
		}
	}
	min := 11
	startPos := -1
	for i := 0; i < len(rows); i++ {
		if rows[i][0] < min {
			startPos = i
			min = rows[i][0]
		}
	}

	res := ""
	for i := 0; i < len(rows); i++ {
		r := rows[(startPos+i)%len(rows)]
		for j := 0; j < len(r); j++ {
			res += fmt.Sprintf("%d", r[j])
		}
	}

	return res
}

func permute(current *[]int, available *[]bool, size int) {
	l := len(*current)
	if l == size {
		p := []int{}
		for i := 0; i < len(*current); i++ {
			p = append(p, (*current)[i]+1)
		}
		solution := checkSolution(p)
		if len(solution) == 16 {
			solutions = append(solutions, solution)
		}
		return
	}
	for i := 0; i < len(*available); i++ {
		if (*available)[i] {
			*current = append((*current), i)
			(*available)[i] = false
			permute(current, available, size)
			(*available)[i] = true
			*current = (*current)[0:l]
		}
	}
}

func main() {
	var available = []bool{}
	for i := 0; i < 10; i++ {
		available = append(available, true)
	}
	permute(&[]int{}, &available, 10)
	sort.Strings(solutions)
	fmt.Println(solutions[len(solutions)-1])
}
