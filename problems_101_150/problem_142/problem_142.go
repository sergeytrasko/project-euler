package main

import (
	"fmt"
	"math"
)

func isSquare(n int) bool {
	s := int(math.Sqrt(float64(n)))
	return s*s == n
}

const limit = 100000

func main() {
	minSum := math.MaxInt
	for a := 1; ; a++ {
		// fmt.Printf("a=%d\n", a)
		for b := 1; b < a; b++ {
			aa, bb := a*a, b*b
			if (aa+bb)%2 == 0 && (aa-bb)%2 == 0 {
				x, y := (aa+bb)/2, (aa-bb)/2 //x+y, x-y
				for c := 1; c < a; c++ {
					cc := c * c
					dd := aa + bb - cc
					if dd > cc || !isSquare(dd) {
						continue
					}
					z := (cc - dd) / 2
					if y > z && isSquare(y+z) && isSquare(y-z) {
						sum := x + y + z
						// fmt.Printf("%d, %d, %d, %d\n", x, y, z, sum)
						if sum < minSum {
							minSum = sum
							fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)
							fmt.Printf("%d+%d=(%f)^2\n", x, y, math.Sqrt(float64(x+y)))
							fmt.Printf("%d-%d=(%f)^2\n", x, y, math.Sqrt(float64(x-y)))
							fmt.Printf("%d+%d=(%f)^2\n", x, z, math.Sqrt(float64(x+z)))
							fmt.Printf("%d-%d=(%f)^2\n", x, z, math.Sqrt(float64(x-z)))
							fmt.Printf("%d+%d=(%f)^2\n", y, z, math.Sqrt(float64(y+z)))
							fmt.Printf("%d-%d=(%f)^2\n", y, z, math.Sqrt(float64(y-z)))
							fmt.Println(sum)
							return
						}
					}

					/*
						for d := 1; d < b; d++ {
							cc, dd := c*c, d*d //x+z, x-z
							if aa+bb == cc+dd && (cc-dd)%2 == 0 {
								z := (cc - dd) / 2
								if y > z && isSquare(y+z) && isSquare(y-z) {
									sum := x + y + z
									fmt.Printf("%d, %d, %d, %d\n", x, y, z, sum)
									// if sum < minSum {
									// minSum = sum
									fmt.Printf("%d, %d, %d, %d\n", x, y, z, sum)
									// }
								}
							}
						}
					*/
				}
			}
		}
	}

	/*
		for x := 1; ; x++ {
			for y := 1; y < x; y++ {
				if isSquare(x+y) && isSquare(x-y) {
					for z := 1; z < y; z++ {
						if isSquare(x+z) && isSquare(x-z) && isSquare(y+z) && isSquare(y-z) {
							fmt.Printf("%d, %d, %d\n", x, y, z)
						}
					}
				}
			}
		}
	*/
	/*
		squares := []int{}
		for i := 1; i < 1000; i++ {
			squares = append(squares, i*i)
		}
		minSum := math.MaxInt
		for a := 0; a < len(squares); a++ {
			for b := 0; b < a; b++ {
				aa, bb := squares[a], squares[b]
				if (aa+bb)%2 == 0 && (aa-bb)%2 == 0 {
					x, y := (aa+bb)/2, (aa-bb)/2 //x+y, x-y
					for c := 0; c < a; c++ {
						for d := 0; d < b; d++ {
							cc, dd := squares[c], squares[d] //x+z, x-z
							if aa+bb == cc+dd && (cc-dd)%2 == 0 {
								z := (cc - dd) / 2
								if y > z && isSquare(y+z) && isSquare(y-z) {
									sum := x + y + z
									fmt.Printf("%d, %d, %d, %d\n", x, y, z, sum)
									if sum < minSum {
										minSum = sum
										fmt.Printf("%d, %d, %d, %d\n", x, y, z, minSum)
									}
								}
							}
						}
					}
				}
			}
		}
	*/
}
