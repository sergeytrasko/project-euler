package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const size = 40

//https://en.wikipedia.org/wiki/Kruskal%27s_algorithm
func main() {
	m := [size][size]int{}
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	row := 0
	sum := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		for i := 0; i < size; i++ {
			if tokens[i] == "-" {
				m[row][i] = -1
			} else {
				v, _ := strconv.Atoi(tokens[i])
				m[row][i] = v
				sum += v
			}
		}
		row++
	}
	sum = sum / 2

	type edge = struct {
		u      int
		v      int
		weight int
	}

	edges := []edge{}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if m[i][j] > 0 {
				edges = append(edges, edge{u: i, v: j, weight: m[i][j]})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})
	// fmt.Println(edges)
	v := [size]int{}
	for i := 0; i < size; i++ {
		v[i] = i
	}
	treeSize := 0
	cnt := 0
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		if v[e.u] != v[e.v] {
			// fmt.Printf("Connecting %d (%d) and %d (%d)\n", e.u, v[e.u], e.v, v[e.v])
			treeSize += e.weight
			cnt++
			sU, sV := v[e.u], v[e.v]
			for j := 0; j < size; j++ {
				if v[j] == sU {
					v[j] = sV
				}
			}
			// fmt.Printf("%d - %d\n", e.weight, v)
		}
	}
	// fmt.Printf("%d iterations\n", cnt)
	// fmt.Println(sum)
	// fmt.Println(treeSize)
	fmt.Println(sum - treeSize)
}
