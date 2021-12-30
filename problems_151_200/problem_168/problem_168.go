package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func naive(len int) int {
	sum := 0
	for n := 10; n < int(math.Pow10(len)); n++ {
		d := n % 10
		if d == 0 {
			continue
		}
		x, _ := strconv.Atoi(fmt.Sprintf("%d%d", d, n/10))
		if x%n == 0 {
			sum = (sum + n) % 100_000
			fmt.Println(n)
		}
	}
	return sum
}

func count(len int) int {
	sum := 0
	for n := 2; n <= len; n++ {
		for d := 1; d <= 9; d++ {
			for k := 1; k <= 9; k++ {
				a := d*int(math.Pow10(n-1)) - k*d
				b := k*10 - 1
				if a%b == 0 {
					num := a/b*10 + d
					if num >= int(math.Pow10(n-1)) {
						sum = (sum + num) % 100_000
						fmt.Println(num)
					}
				}
			}
		}
	}
	return sum
}

func countBig(l int) int64 {
	sum := big.NewInt(0)
	for n := 2; n <= l; n++ {
		for d := 1; d <= 9; d++ {
			for k := 1; k <= 9; k++ {
				a := new(big.Int).Add(
					new(big.Int).Mul(big.NewInt(int64(d)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n-1)), nil)),
					big.NewInt(int64(-k*d)),
				)
				b := big.NewInt(int64(k*10 - 1))
				// a := d*int(math.Pow10(n-1)) - k*d
				// b := k*10 - 1
				if big.NewInt(0).Cmp(new(big.Int).Mod(a, b)) == 0 {
					num := new(big.Int).Add(
						new(big.Int).Mul(
							new(big.Int).Div(new(big.Int).Set(a), new(big.Int).Set(b)),
							big.NewInt(10),
						),
						big.NewInt(int64(d)),
					)
					s := num.String()
					if len(s) >= n {
						sum = sum.Add(sum, num)
						// fmt.Println(num)
					}
				}
			}
		}
	}
	return sum.Mod(sum, big.NewInt(100_000)).Int64()
}

func main() {
	// fmt.Println(naive(6))
	// fmt.Println(count(6))
	fmt.Println(countBig(100))
}
