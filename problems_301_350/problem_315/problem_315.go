package main

import "fmt"

/*
  0
1   2
  3
4   5
  6
*/
var digits = [][]int{
	/*     0  1  2  3  4  5  6*/
	/*0*/ {1, 1, 1, 0, 1, 1, 1},
	/*1*/ {0, 0, 1, 0, 0, 1, 0},
	/*2*/ {1, 0, 1, 1, 1, 0, 1},
	/*3*/ {1, 0, 1, 1, 0, 1, 1},
	/*4*/ {0, 1, 1, 1, 0, 1, 0},
	/*5*/ {1, 1, 0, 1, 0, 1, 1},
	/*6*/ {1, 1, 0, 1, 1, 1, 1},
	/*7*/ {1, 1, 1, 0, 0, 1, 0},
	/*8*/ {1, 1, 1, 1, 1, 1, 1},
	/*9*/ {1, 1, 1, 1, 0, 1, 1},
	/*10*/ {0, 0, 0, 0, 0, 0, 0},
}

func parseToDigits(n int) []int {
	d := []int{}
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	return d
}

func samClock(from, to int) int {
	df := parseToDigits(from)
	dt := parseToDigits(to)
	fSum := 0
	for _, d := range df {
		for _, x := range digits[d] {
			fSum += x
		}
	}
	tSum := 0
	for _, d := range dt {
		for _, x := range digits[d] {
			tSum += x
		}
	}
	return fSum + tSum
}

func samClockSequence(s []int) int {
	r := 0
	for i := 0; i < len(s); i++ {
		r += samClock(-1, s[i]) + samClock(s[i], -1)
	}

	return r
}

func maxClock(from, to int) int {
	df := parseToDigits(from)
	dt := parseToDigits(to)
	for len(df) < len(dt) {
		df = append(df, 10)
	}
	for len(dt) < len(df) {
		dt = append(dt, 10)
	}
	sum := 0
	for i := 0; i < len(dt); i++ {
		digitF, digitT := df[i], dt[i]
		changes := 0
		for j := 0; j < 7; j++ {
			if digits[digitF][j] != digits[digitT][j] {
				changes++
			}
		}
		sum += changes
	}
	return sum
}

func maxClockSequence(s []int) int {
	r := maxClock(-1, s[0])
	for i := 1; i < len(s); i++ {
		r += maxClock(s[i-1], s[i])
	}
	r += maxClock(s[len(s)-1], -1)
	return r
}

func digitalRootSequence(n int) []int {
	r := []int{n}
	for n >= 10 {
		d := parseToDigits(n)
		s := 0
		for _, r := range d {
			s += r
		}
		n = s
		r = append(r, n)
	}
	r = append(r, -1)
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

func main() {
	primes := sieve(2 * 1e7)
	dif := 0
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if p > 1e7 {
			s := digitalRootSequence(p)
			dif += (samClockSequence(s) - maxClockSequence(s))
		}
	}
	fmt.Println(dif)
}
