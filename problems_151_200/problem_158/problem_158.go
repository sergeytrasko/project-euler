package main

import "fmt"

func c(n, k int) int64 {
	if 2*k > n {
		return c(n, n-k)
	}
	p := int64(1)
	q := int64(1)
	for i := 1; i <= k; i++ {
		p *= int64(n + 1 - i)
		q *= int64(i)
	}
	return p / q
}

func pow(a, b int) int64 {
	m := int64(1)
	for i := 0; i < b; i++ {
		m *= int64(a)
	}
	return m
}

func a(n, m int) int64 {
	s := int64(0)
	for k := 0; k <= m+1; k++ {
		s += pow(-1, k) * c(n+1, k) * pow(m+1-k, n)
	}
	return s
}

func p(n int) int64 {
	return c(26, n) * a(n, 1)
}

func main() {
	max := int64(0)
	for n := 1; n <= 26; n++ {
		pn := p(n)
		fmt.Printf("p(%d)=%d\n", n, pn)
		if pn > max {
			max = pn
		}
	}
	fmt.Println(max)
}
