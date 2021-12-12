package main

import "fmt"

func countDivisors(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAbundand(n int) bool {
	return n < countDivisors(n)
}

const max = 28123

func main() {
	abundand := []int{}
	for i := 1; i < max; i++ {
		if isAbundand(i) {
			abundand = append(abundand, i)
		}
	}
	sum := 0
	for n := max; n > 0; n-- {
		can := false
		for i := 0; i < len(abundand); i++ {
			for j := i; j < len(abundand); j++ {
				if abundand[i]+abundand[j] == n {
					can = true
					break
				}
			}
			if can {
				break
			}
		}
		if !can {
			fmt.Println(n)
			sum += n
		}
	}
	fmt.Printf("Sum=%d\n", sum)
}
