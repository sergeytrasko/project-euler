package main

import (
	"fmt"
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

func tryDivisibility(p string) bool {
	factors := []int{1, 2, 3, 5, 7, 11, 13, 17}
	for i := 0; i < len(factors); i++ {
		s := p[i : i+3]
		// fmt.Printf("%s %d\n", s, factors[i])
		n, _ := strconv.Atoi(s)
		if n%factors[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	var available = []bool{}
	for i := 0; i < 10; i++ {
		available = append(available, true)
	}
	permute(&[]int{}, &available, 10)
	sum := int64(0)
	for _, p := range permutations {
		if tryDivisibility(p) {
			// fmt.Println(p)
			n, _ := strconv.Atoi(p)
			sum = sum + int64(n)
		}
	}
	fmt.Println(sum)

}
