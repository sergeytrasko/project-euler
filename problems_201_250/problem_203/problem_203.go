package main

import (
	"fmt"
)

const max = 51

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func primeDivisors(n int) [max]int {
	d := [max]int{}
	for n > 1 {
		for i := 2; i <= n; i++ {
			if isPrime(i) {
				for n%i == 0 {
					// fmt.Printf("%d is prime divisor of %d\n", i, n)
					n /= i
					d[i]++
				}
				// fmt.Println(d)
			}
		}
	}
	return d
}

func factorialPrimeDivisors(n int) [max]int {
	divisors := [max]int{}
	for i := 1; i <= n; i++ {
		d := primeDivisors(i)
		for j := 0; j < max; j++ {
			divisors[j] += d[j]
		}
	}
	return divisors
}

func c(n, k int) (bool, int64) {
	nd := factorialPrimeDivisors(n)
	dd1 := factorialPrimeDivisors(k)
	dd2 := factorialPrimeDivisors(n - k)
	// fmt.Println(nd)
	// fmt.Println(dd1)
	// fmt.Println(dd2)
	for i := 0; i < max; i++ {
		nd[i] = nd[i] - dd1[i] - dd2[i]
	}

	// fmt.Println(nd)
	res := int64(1)
	squareFree := true
	for i := 0; i < max; i++ {
		if nd[i] > 1 {
			squareFree = false
		}
		for j := 1; j <= nd[i]; j++ {
			res *= int64(i)
		}
	}
	return squareFree, res
}

func unique(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	numbers := []int64{}
	for n := 0; n < max; n++ {
		for m := 0; m < n; m++ {
			squareFree, v := c(n, m)
			if squareFree {
				numbers = append(numbers, v)
			}
		}
	}
	numbers = unique(numbers)
	sum := int64(0)
	for _, r := range numbers {
		sum += r
	}
	fmt.Println(sum)
}
