package main

import "fmt"

func f(n int64, mod int64) int64 {
	res := int64(1)
	for i := int64(1); i <= n; i++ {
		res *= i
		for res%10 == 0 {
			res /= 10
		}
		res = res % (mod * mod)
	}

	return res % mod
}

func main() {
	// for n := int64(100); n <= 5000; n += 100 {
	// 	fmt.Printf("f(%d!)=%d\n", n, f(n, 100))
	// }
	// for n := int64(100_000); n <= 5_000_000; n += 100_000 {
	// 	fmt.Printf("f(%d!)=%d\n", n, f(n, 100_000))
	// }
	fmt.Println(f(2560000, 100_000))
}
