package main

import (
	"fmt"
	"sort"
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

func main() {
	var available = []bool{}
	for i := 0; i < 10; i++ {
		available = append(available, true)
	}
	permute(&[]int{}, &available, 10)
	sort.Strings(permutations)
	// fmt.Println(len(permutations))
	// fmt.Println(permutations[0])
	// fmt.Println(permutations[len(permutations)-1])
	fmt.Println(permutations[999999])
}
