package main

import "fmt"

func play(red, blue, redTaken, blueTaken, turnsRemaining int) float64 {
	if turnsRemaining == 0 {
		if blueTaken > redTaken {
			return 1.0
		} else {
			return 0.0
		}
	}
	prob := 0.0
	total := float64(red + blue)
	prob += float64(red) / total * play(red+1, blue, redTaken+1, blueTaken, turnsRemaining-1)
	prob += float64(blue) / total * play(red+1, blue, redTaken, blueTaken+1, turnsRemaining-1)
	return prob
}

func main() {
	probability := play(1, 1, 0, 0, 15)
	fmt.Println(int(1 / probability))
}
