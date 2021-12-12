package main

import (
	"fmt"
	"sort"
)

var primes = []int{}

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
	fmt.Printf("Found %d primes below %d\n", len(primes), max)
	return primes
}

func isPrime(p int, primes []int) bool {
	if p == 1 {
		return false
	}
	i := sort.SearchInts(primes, p)
	return i < len(primes) && primes[i] == p
}

func permutations(min int, max int, result *[][]int) {
	var p func(cur []int, available []int, result *[][]int)
	p = func(cur []int, available []int, result *[][]int) {
		if len(available) == 0 {
			(*result) = append((*result), cur)
		}
		for i := 0; i < len(available); i++ {
			tmp := append([]int{}, available...)
			p(append(cur, available[i]), append(tmp[:i], tmp[i+1:]...), result)
		}
	}
	avail := []int{}
	for i := min; i <= max; i++ {
		avail = append(avail, i)
	}
	p([]int{}, avail, result)
}

func allPrimes(x []int) bool {
	for _, p := range x {
		if !isPrime(p, primes) {
			return false
		}
	}
	return true
}

func isIncreasing(x []int) bool {
	for i := 1; i < len(x); i++ {
		if x[i-1] > x[i] {
			return false
		}
	}
	return true
}

func countPartitions(d []int, pos int, cur int, numbers []int) int {
	// fmt.Printf("pos=%d, cur=%d, numbers=%d\n", pos, cur, numbers)
	if !allPrimes(numbers) {
		return 0
	}
	if pos == len(d) {
		n := append(numbers, cur)
		if allPrimes(n) && isIncreasing(n) {
			// fmt.Println(n)
			return 1
		}
		return 0
	}
	dn := d[pos]
	cnt := 0
	// append digit to current number
	cnt = cnt + countPartitions(d, pos+1, cur*10+dn, numbers)
	// start new number
	if isPrime(cur, primes) {
		cnt = cnt + countPartitions(d, pos+1, dn, append(numbers, cur))
	}
	return cnt
}

func main() {
	primes = sieve(1e8)
	perms := [][]int{}
	permutations(1, 9, &perms)
	cnt := 0
	for _, s := range perms {
		cnt += countPartitions(s, 0, 0, []int{})
		// if isPrimeSet(s) {
		// fmt.Println(s)
		// cnt++
		// }
	}
	fmt.Println(cnt)
}
