package main

import "fmt"

var peter = []int{}
var colin = []int{}

func roll(remaining, sum, max int, results *[]int) {
	if remaining == 0 {
		(*results)[sum]++
	}
	if remaining > 0 {
		for i := 1; i <= max; i++ {
			roll(remaining-1, sum+i, max, results)
		}
	}
}

func countOutcomesMoreThan(outcomes []int, n int) int {
	sum := 0
	for i := n + 1; i < len(outcomes); i++ {
		sum += outcomes[i]
	}
	return sum
}

func main() {
	peter = make([]int, 9*4+1)
	colin = make([]int, 6*6+1)
	roll(9, 0, 4, &peter)
	roll(6, 0, 6, &colin)
	// fmt.Println(peter)
	// fmt.Println(colin)
	ps := 0
	cs := 0
	for i := 1; i < 37; i++ {
		ps += peter[i]
		cs += colin[i]
	}
	peterWins := 0
	for i := 1; i < 37; i++ {
		for j := i + 1; j < 37; j++ {
			peterWins += colin[i] * peter[j]
		}
	}
	fmt.Printf("%.7f\n", float64(peterWins)/float64(ps*cs))
}
