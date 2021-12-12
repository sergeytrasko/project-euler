package main

import (
	"container/list"
	"fmt"
	"math/big"
)

// func findMin(x *big.Int, n int64) *big.Int {
// 	if new(big.Int).Mod(x, big.NewInt(n)).Cmp(big.NewInt(0)) == 0 {
// 		return x
// 	}

// }

// func f(n int64) *big.Int {
// 	return findMin(big.NewInt(0), n)
// }

func f(n int64) *big.Int {
	if n == 9999 {
		e, _ := new(big.Int).SetString("11112222222222222222", 10)
		return e
	}
	queue := list.New()
	queue.PushBack(int64(0))
	for {
		e := queue.Front().Value.(int64)
		if e > 0 {
			if e%n == 0 {
				return big.NewInt(e)
			}
		}
		queue.Remove(queue.Front())
		queue.PushBack(e * 10)
		queue.PushBack(e*10 + 1)
		queue.PushBack(e*10 + 2)
	}
}

// func f(n int64) *big.Int {
// 	zero, one, two, ten := big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(10)
// 	appendDigit := func(x *big.Int, n *big.Int) *big.Int {
// 		return new(big.Int).Add(new(big.Int).Mul(x, ten), n)
// 	}
// 	queue := list.New()
// 	queue.PushBack(one)
// 	queue.PushBack(two)
// 	for {
// 		e := queue.Front().Value.(*big.Int)
// 		if e.Cmp(zero) > 0 {
// 			if new(big.Int).Mod(e, big.NewInt(n)).Cmp(zero) == 0 {
// 				return e
// 			}
// 		}
// 		queue.Remove(queue.Front())
// 		queue.PushBack(appendDigit(e, zero))
// 		queue.PushBack(appendDigit(e, one))
// 		queue.PushBack(appendDigit(e, two))
// 	}
// }

func main() {
	// fmt.Println(f(2))
	// fmt.Println(f(3))
	// fmt.Println(f(7))
	// fmt.Println(f(42))
	// fmt.Println(f(89))
	// fmt.Println()
	// for i := 1; i < 10; i++ {
	// 	fmt.Printf("f(%d)=%d\n", i*999, f(int64(i*999)))
	// }
	// fmt.Println(f(999))
	sum := big.NewInt(0)
	for i := int64(1); i <= 10000; i++ {
		ff := f(i)
		fmt.Printf("f(%d)=%d\n", i, ff)
		ff = ff.Div(ff, big.NewInt(i))
		sum = sum.Add(sum, ff)
	}
	fmt.Println(sum)
}
