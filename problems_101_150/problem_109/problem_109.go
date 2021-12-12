package main

import "fmt"

type Option = struct {
	mul    int
	number int
}

var options = []Option{}

func fillOptions() {
	for i := 1; i <= 20; i++ {
		for m := 1; m <= 3; m++ {
			options = append(options, Option{m, i})
		}
	}
	options = append(options, Option{1, 25})
	options = append(options, Option{2, 25})
}

func checkout(n int) int {
	if n == 1 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(options); i++ {
		last := options[i]
		if last.mul == 2 {
			remaining := n - last.mul*last.number
			if remaining == 0 {
				cnt++
			} else {
				for j := 0; j < len(options); j++ {
					if remaining == options[j].mul*options[j].number {
						cnt++
					}
				}
				for j := 0; j < len(options); j++ {
					for k := j; k < len(options); k++ {
						if remaining == options[j].mul*options[j].number+options[k].mul*options[k].number {
							cnt++
						}
					}
				}
			}
		}
	}
	return cnt
}

func main() {
	fillOptions()
	cnt := 0
	for i := 1; i < 100; i++ {
		cnt += checkout(i)
	}
	fmt.Println(cnt)
}
