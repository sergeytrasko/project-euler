package main

import (
	"fmt"
	"math/big"
	"sort"
)

// func digitsSum(x, y int64) int64 {
// 	r := new(big.Int).Exp(big.NewInt(x), big.NewInt(y), nil).String()
// 	sum := int64(0)
// 	for i := 0; i < len(r); i++ {
// 		sum += int(r[i] - byte('0'))
// 	}
// 	return sum
// }

func sameDigitSum(r string, x int64) bool {
	sum := int64(0)
	for i := 0; i < len(r); i++ {
		sum += int64(r[i] - byte('0'))
	}
	return sum == x
}

func main() {
	res := []*big.Int{}
	for x := int64(2); x < 100; x++ {
		for y := int64(2); y < 100; y++ {
			// d := digitsSum(x, y)
			s := new(big.Int).Exp(big.NewInt(x), big.NewInt(y), nil)
			if sameDigitSum(s.String(), x) {
				fmt.Printf("%d^%d=%s\n", x, y, s)
				res = append(res, s)
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Cmp(res[j]) < 0
	})
	fmt.Println(res[29])
}
