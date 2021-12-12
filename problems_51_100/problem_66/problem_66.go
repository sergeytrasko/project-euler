package main

import (
	"fmt"
	"math"
	"math/big"
)

func sqrt(x int64) (bool, int64) {
	s := int64(math.Sqrt(float64(x)))
	return s*s == x, s
}

//https://en.wikipedia.org/wiki/Periodic_continued_fraction#Canonical_form_and_repetend
//https://en.wikipedia.org/wiki/Continued_fraction#Infinite_continued_fractions_and_convergents
//https://en.wikipedia.org/wiki/Continued_fraction#Pell's_equation
func solvePellEquation(n int64) (big.Int, big.Int) {
	a0 := int64(math.Round(math.Floor(math.Sqrt(float64(n)))))
	m := int64(0)
	d := int64(1)
	a := a0
	x := []big.Int{*big.NewInt(a0)}
	y := []big.Int{*big.NewInt(1)}
	i := 0
	for {
		// fmt.Printf("%d, %d\n", &x[i], &y[i])
		if new(big.Int).Sub(
			new(big.Int).Exp(new(big.Int).Set(&x[i]), big.NewInt(2), nil),
			new(big.Int).Mul(big.NewInt(n), new(big.Int).Exp(new(big.Int).Set(&y[i]), big.NewInt(2), nil)),
		).Cmp(big.NewInt(1)) == 0 {
			return x[i], y[i]
		}
		m = d*a - m
		d = (n - m*m) / d
		a = (a0 + m) / d
		if i == 0 {
			x = append(x, *new(big.Int).Add(new(big.Int).Mul(big.NewInt(a), big.NewInt(a0)), big.NewInt(1)))
			y = append(y, *big.NewInt(a))
		} else {
			x = append(x, *new(big.Int).Add(new(big.Int).Mul(big.NewInt(a), new(big.Int).Set(&x[i])), new(big.Int).Set(&x[i-1])))
			y = append(y, *new(big.Int).Add(new(big.Int).Mul(big.NewInt(a), new(big.Int).Set(&y[i])), new(big.Int).Set(&y[i-1])))
		}
		i++
	}
}

func main() {
	bestD := int64(-1)
	maxX := big.NewInt(0)
	for d := int64(1); d <= 1000; d++ {
		isFullSquare, _ := sqrt(d)
		if !isFullSquare {
			x, y := solvePellEquation(d)
			fmt.Printf("D=%d => x=%d, y=%d\n", d, &x, &y)
			if x.Cmp(maxX) > 0 {
				maxX = &x
				bestD = d
			}
		}
	}
	fmt.Println(bestD)
}
