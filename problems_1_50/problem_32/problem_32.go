package main

import (
	"fmt"
	"sort"
	"strconv"
)

var permutations = []string{}

func permute(current *[]int, available *[]bool, size int) {
	l := len(*current)
	if l == size {
		// fmt.Println(*current)
		s := ""
		for i := 0; i < len(*current); i++ {
			s += fmt.Sprint((*current)[i])
		}
		permutations = append(permutations, s)
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

func tryMultiplications(p string) []int {
	res := []int{}
	for i := 1; i < len(p)-2; i++ {
		for j := i + 1; j < len(p)-1; j++ {
			a, _ := strconv.Atoi(p[0:i])
			b, _ := strconv.Atoi(p[i:j])
			c, _ := strconv.Atoi(p[j:])
			if int64(a)*int64(b) == int64(c) {
				fmt.Printf("%s: %d * %d = %d\n", p, a, b, c)
				res = append(res, c)
			}
		}
	}
	return res
}

func main() {
	// tryMultiplications("391867254")
	var available = []bool{false}
	for i := 1; i < 10; i++ {
		available = append(available, true)
	}
	permute(&[]int{}, &available, 9)

	multiplications := []int{}
	for _, p := range permutations {
		multiplications = append(multiplications, tryMultiplications(p)...)
	}
	sort.Ints(multiplications)
	res := 0
	prev := 0
	for _, m := range multiplications {
		if m != prev {
			res += m
			prev = m
		}
	}

	fmt.Println(res)
}
