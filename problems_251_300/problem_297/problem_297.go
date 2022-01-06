package main

import "fmt"

func z(n int64, fib []int64) int64 {
	k := len(fib) - 1
	l := int64(0)
	for n > 0 {
		for fib[k] > n {
			k--
		}
		n -= fib[k]
		l++
		k--
	}
	return l
}

func sumZ(from, to int64, fib []int64) int64 {
	sum := int64(0)
	for i := from; i <= to; i++ {
		sum += z(i, fib)
	}
	return sum
}

func findK(n int64, fib []int64) int {
	k := len(fib) - 1
	for fib[k] > n {
		k--
		if k < 0 {
			return -1
		}
	}
	return k
}

func sumRec(n int64, fib []int64, memo map[int64]int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] > 0 {
		return memo[n]
	}
	k := findK(n, fib)
	if n == fib[k] {
		s := sumRec(n-1, fib, memo) + 1
		memo[n] = s
		return s
	}
	sum := sumRec(fib[k], fib, memo) + (n - fib[k]) + sumRec(n-fib[k], fib, memo)
	memo[n] = sum
	return sum
}

func main() {
	a, b := int64(0), int64(1)
	fib := []int64{a, b}
	for b < 1e17 {
		a, b = b, a+b
		fib = append(fib, b)
	}
	memo := map[int64]int64{}
	fmt.Println(sumRec(1e17-1, fib, memo))
}
