package main

import "fmt"

func countIncreasing(len int) int64 {
	m := make([][10]int64, len+1)
	for i := 1; i <= 9; i++ {
		m[1][i] = 1
	}
	for i := 2; i <= len; i++ {
		for d := 0; d <= 9; d++ {
			for prevD := 0; prevD <= d; prevD++ {
				m[i][d] += m[i-1][prevD]
			}
		}
	}
	sum := int64(0)
	for i := 0; i <= 9; i++ {
		sum += m[len][i]
	}
	return sum
}

func countDecreasing(len int) int64 {
	m := make([][10]int64, len+1)
	for i := 1; i <= 9; i++ {
		m[1][i] = 1
	}
	for i := 2; i <= len; i++ {
		for d := 0; d <= 9; d++ {
			for prevD := d; prevD <= 9; prevD++ {
				m[i][d] += m[i-1][prevD]
			}
		}
	}
	sum := int64(0)
	for i := 0; i <= 9; i++ {
		sum += m[len][i]
	}
	return sum
}

func countRepUnits(len int) int64 {
	return 9
}

func countNonBouncy(len int) int64 {
	return countIncreasing(len) + countDecreasing(len) - countRepUnits(len)
}

func main() {
	max := 100
	n := 1
	sum := int64(0)
	for n <= max {
		sum += countNonBouncy(n)
		n++
	}
	fmt.Println(sum)

}
