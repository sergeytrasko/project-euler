package main

import (
	"fmt"
	"math/big"
)

const min = 724637681
const max = 729927007

func isPrime(p int64) bool {
	return big.NewInt(p).ProbablyPrime(1)
}

func cycleDigitSum(n int) int {
	digitSum := 0
	x := 1
	cycleLength := 0
	for {
		digitSum += (x * 10) / n
		x = (x * 10) % n
		cycleLength++
		if x == 1 || cycleLength == n {
			break
		}
	}
	if cycleLength == n-1 {
		return digitSum
	}
	return 0
}

func main() {
	for p := min; p <= max; p++ {
		if isPrime(int64(p)) && ((p*56789)%100_000) == 99999 {
			fmt.Printf("N: %d\n", p)
			digitSum := cycleDigitSum(p)
			if digitSum > 0 {
				fmt.Printf("Digit sum: %d\n", digitSum)
			}
		}
	}
}
