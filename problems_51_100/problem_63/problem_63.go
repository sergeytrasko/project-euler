package main

import (
	"fmt"
	"math/big"
)

/*
10^(n-1) <= a^n < 10^n
take lg on all sides

n-1 <= n*lga < n
divide by n > 0

1 - 1/n <= lga < 1
raise to exp 10

10^(1-1/n) <= a < 10

Clearly a < 10 => a belongs to [1; 9] interval
Need to find upper bound for n or find such n so that 10^(1-1/n) > 9

10^(1-/1n) > 9
1-1/n > lg9
1-lg9 > 1/n
n > 1/(1-lg9)

fmt.Println(1.0 / (1.0 - math.Log10(9)))
give us 21.854345326782884
so 1 <= n < 22
*/
func main() {
	cnt := 0
	for n := 1; n < 22; n++ {
		for a := 1; a <= 9; a++ {
			x := big.NewInt(0)
			s := x.Exp(big.NewInt(int64(a)), big.NewInt(int64(n)), nil).String()
			if len(s) == n {
				fmt.Printf("%d^%d=%s\n", a, n, s)
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
