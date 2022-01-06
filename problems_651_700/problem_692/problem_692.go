package main

import "fmt"

func alwaysWins(remaining, canTake int64) bool {
	if remaining <= canTake {
		return true
	}
	opponentAlwaysWins := true
	for move := int64(1); move <= canTake; move++ {
		if opponentAlwaysWins && !alwaysWins(remaining-move, 2*move) {
			opponentAlwaysWins = false
		}
	}
	return !opponentAlwaysWins
}

func h(n int64) int64 {
	for i := int64(1); i <= n; i++ {
		if alwaysWins(n, i) {
			return i
		}
	}
	return n
}

func g(n int64) int64 {
	sum := int64(0)
	for i := int64(1); i <= n; i++ {
		sum += h(i)
	}
	return sum
}

func explore() {
	//first values
	for i := int64(1); i <= 35; i++ {
		fmt.Printf("h(%2d)=%2d\n", i, h(i))
	}
	//is it a know sequence?
	for i := int64(1); i < 10; i++ {
		fmt.Printf("%d,", h(i))
	}
	fmt.Println()

	// calculate some first values of G for n = fib(k)
	fib := []int64{1, 2, 3, 5, 8, 13, 21, 34}
	for i := 0; i < len(fib); i++ {
		f := fib[i]
		fmt.Printf("g(f(%d)) = %d, f(%d)=%d\n", i+1, g(f), i+1, f)
	}
}

func main() {
	// explore()
	N := int64(23416728348467685)
	a, b := int64(1), int64(1)
	fib := []int64{a, b}
	for b != N {
		a, b = b, a+b
		fib = append(fib, b)
	}
	// fmt.Println(fib)
	gg := []int64{1, 1} // g(0) = g(1) = h(1)
	for n := 2; n < len(fib); n++ {
		res := gg[n-1] + gg[n-2] - fib[n-2] + fib[n]
		gg = append(gg, res)
		// fmt.Printf("%d - %d\n", fib[n], res)
	}
	fmt.Println(gg[len(gg)-1])
}
