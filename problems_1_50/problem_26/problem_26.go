package main

import "fmt"

func findCycle(n int) int {
	reminders := []int{}
	x := 1
	for x != 0 {
		reminders = append(reminders, x)
		for i := 0; i < len(reminders)-1; i++ {
			if reminders[i] == x {
				return len(reminders) - i - 1
			}
		}
		if x < n {
			x = x * 10
		} else {
			x = (x % n) * 10
		}
	}
	return 0
}

func main() {
	maxCycle, best := 0, 0
	for i := 2; i < 1000; i++ {
		cycle := findCycle(i)
		if cycle > maxCycle {
			best = i
			maxCycle = cycle
		}
	}
	fmt.Println(best)
}
