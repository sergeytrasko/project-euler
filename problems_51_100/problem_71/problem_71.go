package main

import "fmt"

const n = 1_000_000

/*
a   c
- < -
b   d

a*d < c*b
*/
func less(a int, b int, c int, d int) bool {
	return a*d < c*b
}

//https://en.wikipedia.org/wiki/Farey_sequence#Next_term
func main() {
	a, b := 0, 1
	c, d := 1, 1
	p, q := 0, 0
	for b+d <= n {
		p, q = a+c, b+d
		// fmt.Printf("%d/%d - %d/%d - %d/%d\n", a, b, p, q, c, d)
		if (p == 3 && q == 7) || !less(p, q, 3, 7) {
			c, d = p, q
		} else {
			a, b = p, q
		}
	}
	if p == 3 && q == 7 {
		fmt.Printf("%d/%d\n", a, b)
		fmt.Println(a)
	} else {
		fmt.Printf("%d/%d\n", p, q)
		fmt.Println(p)
	}

	/*
		a, b := int64(0), int64(1)
		c, d := int64(1), int64(n)
		for {
			// fmt.Printf("%d/%d\n", c, d)
			p := int64(math.Floor(float64(n+b)/float64(d)))*c - a
			q := int64(math.Floor(float64(n+b)/float64(d)))*d - b
			a, b, c, d = c, d, p, q
			if c == 3 && d == 7 {
				break
			}
		}
		fmt.Printf("%d/%d\n", a, b)
	*/
}
