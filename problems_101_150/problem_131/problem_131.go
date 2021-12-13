package main

import (
	"fmt"
	"math"
	"math/big"
)

var primes = []int64{}

func isPrime(p int64) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	primes = append(primes, p)
	return true
}

func isCube(m int64) bool {
	x := int64(math.Round(math.Pow(float64(m), float64(1.0/3.0))))
	// fmt.Println(x)
	return x*x*x == m
}

func cbrtBinary(i *big.Int) (cbrt *big.Int, rem *big.Int) {
	var (
		n0 = big.NewInt(0)
		n1 = big.NewInt(1)
		n2 = big.NewInt(2)
		n3 = big.NewInt(3)
	)
	var (
		guess = new(big.Int).Div(i, n2)
		dx    = new(big.Int)
		absDx = new(big.Int)
		minDx = new(big.Int).Abs(i)
		step  = new(big.Int).Abs(new(big.Int).Div(guess, n2))
		cube  = new(big.Int)
	)
	for {
		cube.Exp(guess, n3, nil)
		dx.Sub(i, cube)
		cmp := dx.Cmp(n0)
		if cmp == 0 {
			return guess, n0
		}

		absDx.Abs(dx)
		switch absDx.Cmp(minDx) {
		case -1:
			minDx.Set(absDx)
		case 0:
			return guess, dx
		}

		switch cmp {
		case -1:
			guess.Sub(guess, step)
		case +1:
			guess.Add(guess, step)
		}

		step.Div(step, n2)
		if step.Cmp(n0) == 0 {
			step.Set(n1)
		}
	}
}

func isBigCube(m *big.Int) (*big.Int, bool) {
	res, rem := cbrtBinary(m)
	return res, rem.Cmp(big.NewInt(0)) == 0
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

func main() {
	/*
		primes := sieve(1_000_000)
		for i := 0; i < len(primes); i++ {
			p := int64(primes[i])
			for k := int64(1); k < 10; k++ {
				n := k * k * k
				m3 := n*n*n + n*n*p
				if isCube(m3) {
					fmt.Printf("%d^3+%d^*%d=%d, n=(%d)^3\n", n, n, p, m3, int(math.Round(math.Pow(float64(n), 1.0/3.0))))
				}
			}
		}
		fmt.Println("Hypithesis - n is a perfect cube")
	*/

	primes := sieve(1_000_000)
	cnt := 0
	kStart := int64(1)
	for i := 0; i < len(primes); i++ {
		p := int64(primes[i])
		for k := kStart; k < kStart+20; k++ {
			n := new(big.Int).Exp(big.NewInt(k), big.NewInt(int64(3)), nil)
			m1 := new(big.Int).Exp(new(big.Int).Set(n), big.NewInt(int64(3)), nil)
			m2 := new(big.Int).Exp(new(big.Int).Set(n), big.NewInt(int64(2)), nil)
			m2 = m2.Mul(m2, big.NewInt(p))
			m3 := m1.Add(m1, m2)
			r, isCube := isBigCube(m3)
			if isCube {
				fmt.Printf("%d^3+%d^*%d=%d, n=(%d)^3, k=%d\n", n, n, p, m3, r, k)
				kStart = k + 1
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)

	/*
		for p := int64(2); p < 1_000_000; p++ {
			if isPrime(p) {
				for a := int64(1); a < 1000; a++ {
					n := new(big.Int).Exp(big.NewInt(a), big.NewInt(int64(3)), nil)
					m1 := new(big.Int).Exp(new(big.Int).Set(n), big.NewInt(int64(3)), nil)
					m2 := new(big.Int).Exp(new(big.Int).Set(n), big.NewInt(int64(2)), nil)
					m2 = m2.Mul(m2, big.NewInt(p))
					m3 := m1.Add(m1, m2)
					// n := a * a * a
					// m3 := n*n*n + n*n*p
					// if isCube(m3) {
					// 	fmt.Printf("%d^3+%d^*%d=%d, n=(%d)^3\n", n, n, p, m3, a)
					// 	break
					// }
				}
			}
		}
		fmt.Println("Hypithesis - n is a perfect cube")
	*/
	// cnt := 0
	// for p := int64(2); p < 1e9; p++ {
	// 	if isPrime(p) {
	// 		for a := int64(1); a < 100; a++ {
	// 			n := a * a * a
	// 			m3 := n*n*n + n*n*p
	// 			if isCube(m3) {
	// 				fmt.Printf("%d^3+%d^*%d=%d\n", n, n, p, m3)
	// 				a++
	// 				cnt++
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(cnt)
}
