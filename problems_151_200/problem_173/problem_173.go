package main

import "fmt"

const maxTiles = 1_000_000

func tilesForRow(n int) int {
	return 4 * n
}

func main() {
	cnt := 0
	holeSize := 1
	for {
		used := 0
		options := 0
		for i := holeSize + 1; ; i += 2 {
			row := tilesForRow(i)
			if used+row > maxTiles {
				break
			}
			used += row
			options++
			// fmt.Printf("hole size=%d, option %d has %d rows, total used %d\n", holeSize, options, i-holeSize, used)
			cnt++
		}
		if options == 0 {
			break
		}
		holeSize++
	}
	fmt.Println(cnt)
}
