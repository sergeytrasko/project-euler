package main

import "fmt"

func ok(n int) bool {
	d := []int{}
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	if len(d) < 3 {
		s := 0
		for i := 0; i < len(d); i++ {
			s += d[i]
		}
		return s <= 9
	}
	for i := 0; i < len(d)-2; i++ {
		s := d[i] + d[i+1] + d[i+2]
		if s > 9 {
			return false
		}
	}
	return true
}

func main() {
	// cnt := 0
	// x := 100000
	// for i := x; i <= x*10-1; i++ {
	// 	if ok(i) {
	// 		// fmt.Println(i)
	// 		cnt++
	// 	}
	// }
	// fmt.Println(cnt)
	m := [21][1000]int64{}
	for i := 0; i < 1000; i++ {
		if ok(i) {
			m[3][i] = 1
		}
	}
	for n := 4; n <= 20; n++ {
		for i := 0; i < 1000; i++ {
			f := i / 10
			for j := 0; j <= 9; j++ {
				sum := f/10 + f%10 + j
				if sum > 9 {
					continue
				}
				p := j*100 + f
				m[n][p] += m[n-1][i]
			}
		}
	}
	s := int64(0)
	// fmt.Println(m[0])
	// fmt.Println(m[1])
	// fmt.Println(m[2])
	for i := 100; i < 1000; i++ {
		s += m[20][i]
	}
	fmt.Println(s)
}
