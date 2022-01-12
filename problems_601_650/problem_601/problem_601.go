package main

import "fmt"

func explore() {
	streak := func(n int) int {
		for k := 1; ; k++ {
			if (n+k)%(k+1) != 0 {
				return k
			}
		}
	}
	//first members to search is OEIS
	for i := 2; i < 50; i++ {
		fmt.Printf("%d,", streak(i))
	}
	fmt.Println()
	fmt.Println(streak(13))
	fmt.Println(streak(120))
	// explore pattern
	// cnt := 0
	// for n := 2; cnt < 20; n++ {
	// 	if streak(n) == 6 {
	// 		fmt.Printf("%d,", n)
	// 		cnt++
	// 	}
	// }
}

func lcm(a, b int64) int64 {
	return a * (b / gcd(a, b))
}

func lcmArray(a []int64) int64 {
	if len(a) == 1 {
		return a[0]
	}
	return lcm(a[0], lcmArray(a[1:]))
}

func gcd(x, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func p(s int64, N int64) int64 {
	a := []int64{}
	for i := int64(1); i <= s; i++ {
		a = append(a, i)
	}
	return (N - 2) / lcmArray(a)
}

func main() {
	// explore()
	sum := int64(0)
	fourI := int64(1)
	for i := int64(1); i <= 31; i++ {
		fourI *= 4
		sum += (p(i, fourI) - p(i+1, fourI))
	}
	fmt.Println(sum)
}
