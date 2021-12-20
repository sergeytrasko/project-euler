package main

import (
	"fmt"
	"math"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func count90Degrees(max int) int {
	cnt := 0
	for m := 2; m*m < max; m++ {
		for n := 1; n < m; n++ {
			if (n+m)%2 == 1 && gcd(n, m) == 1 {
				a, b, c := m*m-n*n, 2*m*n, m*m+n*n
				p := a + b + c
				if a+b+c > max {
					break
				}
				// fmt.Printf("(a, b, c) = (%d, %d, %d), p=%d x %d\n", a, b, c, p, max/p)
				cnt += max / p
			}
		}
	}
	fmt.Printf("Found %d of 90 degree triangles\n", cnt)
	return cnt
}

func count60Degrees(max int) int {
	cnt := 0
	for m := 2; m*m < 3*max; m++ {
		for n := 1; 2*n <= m; n++ {
			if gcd(n, m) == 1 {
				a, b, c := m*m-m*n+n*n, 2*m*n-n*n, m*m-n*n
				// fmt.Printf("(a, b, c) = (%d, %d, %d)\n", a, b, c)
				d := gcd(a, gcd(b, c))
				a, b, c = a/d, b/d, c/d
				p := a + b + c
				if p > max*3 {
					break
				}
				// fmt.Printf("(a, b, c) = (%d, %d, %d), p=%d x %d\n", a, b, c, p, max/p)
				cnt += max / p
			}
		}
	}
	fmt.Printf("Found %d of 60 degree triangles\n", cnt)
	return cnt
}

func count120Degrees(max int) int {
	cnt := 0
	for m := 2; m*m < 3*max; m++ {
		for n := 1; n < m; n++ {
			if gcd(n, m) == 1 {
				a, b, c := m*m+m*n+n*n, 2*m*n+n*n, m*m-n*n
				if b > c {
					break // mirrored
				}
				d := gcd(a, gcd(b, c))
				a, b, c = a/d, b/d, c/d
				p := a + b + c
				if p > max*3 {
					break
				}
				// fmt.Printf("(a, b, c) = (%d, %d, %d), p=%d x %d\n", a, b, c, p, max/p)
				cnt += max / p
			}
		}
	}
	fmt.Printf("Found %d of 120 degree triangles\n", cnt)
	return cnt
}

func naive(max int) int {
	cnt := 0
	for a := 1; a < maxP; a++ {
		for b := a; b < maxP; b++ {
			for _, cos := range []float64{
				0.5,  //60
				0,    //90
				-0.5, //120
			} {
				cc := float64(a*a+b*b) - float64(2*a*b)*cos
				c := math.Sqrt(cc)
				if a+b+int(c) > max {
					break
				}
				if float64(int(c)) == c {
					cnt++
				}
			}
		}
	}
	return cnt
}

const maxP = 1e8

func main() {
	fmt.Println(count90Degrees(maxP) + count60Degrees(maxP) + count120Degrees(maxP))
	// fmt.Println(naive(maxP))
}
