package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

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

func coundDivisors(max int) []int {
	primes := sieve(max)
	d := make([]int, max+1)
	for i := 1; i <= max; i++ {
		d[i] = 1
	}
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		for j := 1; p*j <= max; j++ {
			x := p * j
			cnt := 0
			for x%p == 0 {
				cnt++
				x /= p
			}
			d[p*j] *= (cnt + 1)
		}
	}
	return d
}

func m(n, k int, d []int) int {
	max := 0
	for j := n; j <= n+k-1; j++ {
		if d[j] > max {
			max = d[j]
		}
	}
	return max
}

func s(u, k int, d []int) int {
	sum := 0
	for n := 1; n <= u-k+1; n++ {
		sum += m(n, k, d)
		// fmt.Printf("%d ", m(n, k, d))
	}
	// fmt.Println()
	return sum
}

func maxSliding(d []int, k int) []int {
	var q deque.Deque
	var i int
	res := []int{}
	for i = 0; i < k; i++ {
		for q.Len() > 0 && d[i] >= d[q.Back().(int)] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	for ; i < len(d); i++ {
		res = append(res, d[q.Front().(int)])
		for q.Len() > 0 && q.Front().(int) <= i-k {
			q.PopFront()
		}
		for q.Len() > 0 && d[i] >= d[q.Back().(int)] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	res = append(res, d[q.Front().(int)])
	return res
}

const U = 100_000_000
const K = 100_000

func main() {
	d := coundDivisors(U)
	// fmt.Println(s(U, K, d))
	m := maxSliding(d[1:], K)
	// fmt.Println(m)
	sum := 0
	for i := 0; i < len(m); i++ {
		sum += m[i]
	}
	fmt.Println(sum)
}
