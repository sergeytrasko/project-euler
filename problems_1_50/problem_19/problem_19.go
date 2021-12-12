package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := 0
	for y := 1901; y <= 2000; y++ {
		for m := time.Month(time.January); m <= time.Month(time.December); m++ {
			date := time.Date(y, m, 1, 12, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Sunday {
				// fmt.Printf("%d-%d-%d was Sunday\n", 1, m, y)
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
