package main

import (
	"fmt"
	"math"
)

func generate(cur int64, p int, n int, d int, sameDigitCount int, result *[]int64) {
	if p == n {
		if sameDigitCount == 0 {
			*result = append(*result, cur)
		}
		return
	}
	if p < n {
		for i := 0; i <= 9; i++ {
			if i == 0 && cur == 0 {
				continue //cannot start with 0
			}
			if int(i) == d && sameDigitCount == 0 {
				continue //cannot add this digit
			}
			if sameDigitCount > n-p {
				continue //cannot fit all digits
			}
			curNext := cur*10 + int64(i)
			if int(i) == d {
				generate(curNext, p+1, n, d, sameDigitCount-1, result)
			} else {
				generate(curNext, p+1, n, d, sameDigitCount, result)
			}
		}
	}
}

func generateNumbers(n int, d int, sameDigitCount int) []int64 {
	r := []int64{}
	generate(0, 0, n, d, sameDigitCount, &r)
	return r
}

func sieve(max int) []int {
	n := make([]bool, max)
	primes := []int{}
	for i := 2; i < max; i++ {
		if !n[i] {
			primes = append(primes, i)
			for j := 1; ; j++ {
				if i*j >= max {
					break
				}
				n[i*j] = true
			}
		}
	}
	return primes
}

func filterPrimes(nums []int64, primes []int) []int64 {
	res := []int64{}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		isPrime := true
		for j := 0; j < len(primes); j++ {
			p := int64(primes[j])
			if n%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			res = append(res, n)
		}
	}
	return res
}

const n = 10

func main() {
	primes := sieve(int(math.Sqrt(math.Pow10(n))))
	sum := int64(0)
	for d := 0; d <= 9; d++ {
		for nn := n - 1; nn >= 0; nn-- {
			numbers := generateNumbers(n, d, nn)
			p := filterPrimes(numbers, primes)
			if len(p) > 0 {
				s := int64(0)
				for i := 0; i < len(p); i++ {
					s += p[i]
				}
				sum += s
				fmt.Printf("M(%d,%d)=%d, N(%d,%d)=%d, S(%d,%d)=%d\n", n, d, nn, n, d, len(p), n, d, s)
				break
			}
		}
	}
	fmt.Println(sum)
}
