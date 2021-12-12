package main

import "fmt"

func count(m [16]int, pos int, expectedSum int, cnt *int) {
	if pos == 4 {
		//1-st row
		// expectedSum = m[0] + m[1] + m[2] + m[3]
		if m[0]+m[1]+m[2]+m[3] != expectedSum {
			return
		}
	}
	if pos == 8 {
		//2-nd row
		if m[4]+m[5]+m[6]+m[7] != expectedSum {
			return
		}
	}
	if pos == 12 {
		//3-rd row
		if m[8]+m[9]+m[10]+m[11] != expectedSum {
			return
		}
	}
	if pos == 13 {
		//1-st column + 1-st diagonal
		if m[0]+m[4]+m[8]+m[12] != expectedSum {
			return
		}
		if m[3]+m[6]+m[10]+m[12] != expectedSum {
			return
		}
	}
	if pos == 14 {
		//2-nd column
		if m[1]+m[5]+m[9]+m[13] != expectedSum {
			return
		}
	}
	if pos == 15 {
		//3-rd column
		if m[2]+m[6]+m[10]+m[14] != expectedSum {
			return
		}
	}
	if pos == 16 {
		//4-th column + 2-nd diagonal
		if m[3]+m[7]+m[11]+m[15] != expectedSum {
			return
		}
		if m[0]+m[5]+m[10]+m[15] != expectedSum {
			return
		}
		*cnt = *cnt + 1
		return
	}
	for i := 0; i <= 9; i++ {
		m[pos] = i
		count(m, pos+1, expectedSum, cnt)
	}
}

func count2() {
	/*
		a b c d
		e f g h
		i j k l
		m n o p
	*/
	cnt := 0
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					s := a + b + c + d
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							for g := 0; g <= 9; g++ {
								h := s - e - f - g
								if h < 0 || h > 9 {
									continue
								}
								for i := 0; i <= 9; i++ {
									m := s - a - e - i
									if m < 0 || m > 9 {
										continue
									}
									for j := 0; j <= 9; j++ {
										n := s - b - f - j
										if n < 0 || n > 9 {
											continue
										}
										if s-d-g-j != m {
											continue
										}
										for k := 0; k <= 9; k++ {
											l := s - i - j - k
											if l < 0 || l > 9 {
												continue
											}
											o := s - c - g - k
											if o < 0 || o > 9 {
												continue
											}
											p := s - d - h - l
											if p < 0 || p > 9 {
												continue
											}
											if s-m-n-o != p {
												continue
											}
											if s-a-f-k != p {
												continue
											}
											cnt++
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(cnt)
}

func main() {
	// cnt := 0
	// for i := 0; i <= 18; i++ {
	// 	x := 0
	// 	count([16]int{}, 0, i, &x)
	// 	fmt.Printf("%d (and also %d) - %d\n", i, 36-i, x)
	// 	cnt += x * 2
	// }
	// count([16]int{}, 0, 17, &cnt)
	// fmt.Println(cnt)
	count2()
}
