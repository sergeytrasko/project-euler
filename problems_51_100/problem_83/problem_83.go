package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const max = 80
const START = max * max
const FINISH = max*max + 1

func readMatrix() [max][max]int {
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
	return m
}

func vertexIndex(x, y int) int {
	return y*max + x
}

func coords(index int) (int, int) {
	return index % max, index / max
}

func outgoingEdges(from int) []int {
	e := []int{}
	if from == FINISH {
		return []int{}
	} else if from == START {
		e = append(e, vertexIndex(0, 0))
	} else {
		x, y := coords(from)
		if y > 0 {
			e = append(e, vertexIndex(x, y-1))
		}
		if y < max-1 {
			e = append(e, vertexIndex(x, y+1))
		}
		if x > 0 {
			e = append(e, vertexIndex(x-1, y))
		}
		if x < max-1 {
			e = append(e, vertexIndex(x+1, y))
		}
		if x == max-1 && y == max-1 {
			e = append(e, FINISH)
		}
	}
	return e
}

func edge(from, to int, m [max][max]int) int {
	if to == FINISH {
		return 0
	} else {
		x, y := coords(to)
		return m[y][x]
	}
}

// Dijkstra algorithm
func main() {
	maxD := 10000000
	m := readMatrix()
	visited := [max*max + 2]bool{}
	d := [max*max + 2]int{}
	for i := 0; i < len(d); i++ {
		if i != START {
			d[i] = maxD
		}
	}
	for {
		v := -1
		minD := maxD + 1
		for i := 0; i < len(d); i++ {
			if i != FINISH && minD > d[i] && !visited[i] {
				minD = d[i]
				v = i
			}
		}
		if v == -1 {
			break
		}
		visited[v] = true
		// fmt.Printf("v=%d, d=%d\n", v, d)
		e := outgoingEdges(v)
		// fmt.Printf("Edges=%d\n", e)
		for i := 0; i < len(e); i++ {
			u := e[i]
			if !visited[u] {
				// fmt.Printf("u=%d, d[u]=%d, d[v]=%d, edge(v, u)=%d\n", u, d[u], d[v], edge(v, u, m))
				if d[u] > d[v]+edge(v, u, m) {
					d[u] = d[v] + edge(v, u, m)
				}
			}
		}
	}
	// fmt.Println(d)
	fmt.Println(d[FINISH])

}
