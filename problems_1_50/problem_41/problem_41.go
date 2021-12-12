package main

import (
	"fmt"
	"sort"
)

var pandigital = []int{}

func isPrime(p int) bool {
	if p == 1 {
		return false
	}
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func toNumber(a []int) int {
	n := 0
	for i := 0; i < len(a); i++ {
		n = n*10 + a[i]
	}
	return n
}

func permute(current *[]int, available *[]bool, size int) {
	l := len(*current)
	if l == size {
		// fmt.Println(*current)
		n := toNumber(*current)
		if isPrime(n) {
			pandigital = append(pandigital, n)
		}
		return
	}
	for i := 1; i < len(*available); i++ {
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
	for n := 1; n <= 9; n++ {
		available := []bool{false}
		for i := 0; i < n; i++ {
			available = append(available, true)
		}
		permute(&[]int{}, &available, n)
	}
	sort.Ints(pandigital)
	fmt.Println(pandigital[len(pandigital)-1])
}
