package main

import "fmt"

func sumDivisors(n int) int {
	sum := 1
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			sum += i
			d := n / i
			if d != i {
				sum += d
			}
		}
	}
	return sum
}

const max = 1_000_000

func amicable(i int, divisors []int) (bool, []int, int) {
	visited := [max]int{}
	x := i
	len := 1
	smalles := i
	chain := []int{i}
	for {
		visited[x] = len
		d := divisors[x]
		if d < i {
			return false, chain, 0
		}
		if d > max {
			return false, chain, 0
		}
		if d == i {
			return true, chain, smalles
		}
		if visited[d] > 0 || d == 1 {
			return false, chain, 0
		}
		if d < smalles {
			smalles = d
		}
		x = d
		chain = append(chain, d)
		len++
	}
}

func main() {
	divisors := [max]int{}
	for i := 1; i < max; i++ {
		divisors[i] = sumDivisors(i)
	}
	fmt.Println("calculated divisors")
	bestLen := 0
	bestSmallest := 0
	for i := 1; i < max; i++ {
		if i%5000 == 0 {
			fmt.Println(i)
		}
		ok, chain, min := amicable(i, divisors[:])
		if ok {
			if len(chain) > bestLen {
				bestLen = len(chain)
				bestSmallest = min
				fmt.Printf("%d has chain of %d elements, smalles of them is %d\n", i, len(chain), min)
			}
		}
	}
	fmt.Println(bestLen)
	fmt.Println(bestSmallest)
}
