package main

import (
	"fmt"
	"strconv"
)

var primes = []int{}
var isPrimeNumber = map[int]bool{}

func isPrime(p int) bool {
	for i := 2; i*i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func replace(p int, positions []int, n int) int {
	s := fmt.Sprintf("%d", p)
	for i := 0; i < len(positions); i++ {
		s = s[0:positions[i]] + string(byte('0')+byte(n)) + s[positions[i]+1:]
	}
	x, _ := strconv.Atoi(s)
	return x
}

func numberLength(n int) int {
	return len(fmt.Sprintf("%d", n))
}

func countPrimeFamily(p int, positions []int) []int {
	res := []int{}
	for i := 0; i <= 9; i++ {
		x := replace(p, positions, i)
		if isPrimeNumber[x] && numberLength(p) == numberLength(x) {
			res = append(res, x)
		}
	}
	return res
}

func try(p int) []int {
	best := []int{}
	digits := len(fmt.Sprintf("%d", p))
	result := [][]int{}
	choose(digits, 0, []int{}, &result)
	for i := 0; i < len(result); i++ {
		if len(result[i]) != 0 && len(result[i]) != digits {
			family := countPrimeFamily(p, result[i])
			if len(family) > len(best) {
				best = family
			}
		}
	}
	return best
}

func choose(size int, n int, current []int, result *[][]int) {
	if n == size {
		*result = append(*result, current)
		return
	}
	choose(size, n+1, current, result)
	choose(size, n+1, append(current, n), result)
}

const max = 1_000_000

func main() {
	for p := 2; p < max; p++ {
		if isPrime(p) {
			primes = append(primes, p)
			isPrimeNumber[p] = true
		}
	}
	for n := 55_000; ; n++ {
		if isPrimeNumber[n] {
			family := try(n)
			if len(family) == 8 {
				fmt.Println(family)
				fmt.Println(len(family))
				break
			}
		}
	}
}
