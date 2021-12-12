package main

import (
	"fmt"
	"math"
	"math/big"
)

const precision = 120
const buffer = 5

var modulo = big.NewInt(int64(10)).Exp(big.NewInt(int64(10)), big.NewInt(precision+buffer), nil)

func zeros(n int) string {
	s := ""
	for len(s) < n {
		s = s + "0"
	}
	return s
}

func initBig(n int) *big.Int {
	s := fmt.Sprintf("%d%s", n, zeros(precision+buffer))
	i, _ := big.NewInt(0).SetString(s, 10)
	return i
}

func middle(a *big.Int, b *big.Int) *big.Int {
	m := big.NewInt(0)
	m = m.Add(a, b)
	return m.Div(m, big.NewInt(2))
}

func square(x *big.Int) *big.Int {
	r := big.NewInt(1)
	r = r.Mul(x, x)
	return r.Div(r, modulo)
}

func precisionReached(l *big.Int, r *big.Int) bool {
	d := big.NewInt(0)
	d = d.Sub(r, l)
	p := big.NewInt(0)
	p = p.Exp(big.NewInt(int64(10)), big.NewInt(buffer), nil)
	return d.Cmp(p) <= 0
}

func approx(n int, lower int, upper int) string {
	target := initBig(n)
	l, r, m := initBig(lower), initBig(upper), big.NewInt(0)
	for !precisionReached(l, r) {
		m = middle(l, r)
		// fmt.Printf("%d - %d, %d\n", l, r, m)
		s := square(m)
		// fmt.Println(s)
		// fmt.Println(target)
		if target.Cmp(s) > 0 {
			l = m
			// fmt.Println("Move left")
		} else {
			r = m
			// fmt.Println("Move right")
		}
		// fmt.Println("====")

	}
	// fmt.Printf("sqrt(%d)=%d\n", n, m)
	return m.String()
}

func approxSqrt(n int) (bool, string) {
	lower := int(math.Sqrt(float64(n)))
	if lower*lower == n {
		// fmt.Printf("sqrt(%d) is a whole number", n)
		return false, initBig(lower).String()
	}
	upper := lower + 1
	// fmt.Printf("sqrt(%d) is between %d and %d\n", n, lower, upper)
	return true, approx(n, lower, upper)
}

func digitSum(s string, firstDigits int) int {
	sum := 0
	for i := 0; i < firstDigits; i++ {
		sum += int(s[i] - byte('0'))
	}
	return sum
}

func main() {
	sum := 0
	for i := 1; i < 100; i++ {
		irrational, approx := approxSqrt(i)
		if irrational {
			sum += digitSum(approx, 100)
		}
	}
	fmt.Println(sum)
	/*
		for i := 1; i < 100; i++ {
			irrational, approx := approxSqrt(i)
			if irrational {
				calc := approx[0 : precision-1]
				real := fmt.Sprintf("%f", math.Sqrt(float64(i))*1000000)[0 : precision-1]
				if calc != real {
					fmt.Printf("%s - %s\n", calc, real)
				}
				// digitSum(approx)
				// fmt.Printf("sqrt(%d)=%f\n", i, math.Sqrt(float64(i))*100)
			}
			// fmt.Println()
		}
	*/
}
