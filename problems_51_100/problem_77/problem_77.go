package main

import "fmt"

var values = []int{0}

const max = 100
const vl = 50

func isPrime(p int) bool {
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func fillPrimes() {
	cnt := 0
	p := 2
	for cnt < vl {
		if isPrime(p) {
			values = append(values, p)
			cnt++
		}
		p++
	}
}

func main() {
	fillPrimes()
	p := [max + 1][vl + 1]int64{}
	for n := 1; n <= max; n++ {
		for k := 1; k < vl; k++ {
			if n == 1 || k == 1 {
				p[n][k] = 1
			}
			if n > values[k] {
				p[n][k] = p[n][k-1] + p[n-values[k]][k]
			} else if n < values[k] {
				p[n][k] = p[n][k-1]
			} else if n == values[k] {
				p[n][k] = p[n][k-1] + 1
			}
		}
	}
	for n := 1; n <= max; n++ {
		for k := 1; k < vl; k++ {
			if p[n][k] > 5000 {
				fmt.Println(n)
				return
			}
		}
	}
	// fmt.Println(p[max][vl-1])
}
