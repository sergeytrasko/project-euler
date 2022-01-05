package main

import "fmt"

const maxK = 12_000
const limit = 2*maxK + 1

func findK(n int, remaining int, k int, sum int, product int, min *[limit]int) {
	if sum > 2*n || product > 2*n {
		return
	}
	if remaining == 1 && k >= 2 {
		if sum != product {
			ones := product - sum
			k += ones
		}
		if min[k] > n {
			min[k] = n
		}
		return
	}
	for i := 2; i <= remaining; i++ {
		if remaining%i == 0 {
			findK(n, remaining/i, k+1, sum+i, product*i, min)
		}
	}
}

func main() {
	min := [limit]int{}
	for i := 0; i < limit; i++ {
		min[i] = 2*i + 1
	}
	for n := 2; n < limit; n++ {
		findK(n, n, 0, 0, 1, &min)
	}
	distinct := map[int]bool{}
	for k := 2; k <= maxK; k++ {
		// fmt.Printf("mps(%d)=%d\n", k, min[k])
		distinct[min[k]] = true
	}
	sum := 0
	for k := range distinct {
		sum += k
	}
	fmt.Println(sum)
}
