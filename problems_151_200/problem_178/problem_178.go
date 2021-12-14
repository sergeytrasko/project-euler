package main

import "fmt"

func main() {
	dp := [42][1024][10]int64{}
	for d := 1; d <= 9; d++ {
		dp[1][1<<d][d] = 1
	}
	sum := int64(0)
	for n := 2; n <= 40; n++ {
		for m := 1; m < 1024; m++ {
			for d := 0; d <= 9; d++ {
				if d > 0 {
					dp[n][m|1<<d][d] += dp[n-1][m][d-1]
				}
				if d < 9 {
					dp[n][m|1<<d][d] += dp[n-1][m][d+1]
				}
			}
		}
		s := int64(0)
		for d := 0; d <= 9; d++ {
			s += dp[n][1023][d]
		}
		sum += s
	}
	fmt.Println(sum)
}
