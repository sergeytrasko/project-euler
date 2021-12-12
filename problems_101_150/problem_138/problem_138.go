package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func tripples() {
	for m := 2; m < 100; m++ {
		for n := 1; n < m; n++ {
			if (n+m)%2 == 1 && gcd(n, m) == 1 {
				h, b, L := m*m-n*n, 2*2*m*n, m*m+n*n
				// need h-b = +-1 or m^2-n^2-4mn +-1 = 0
				if h-b == 1 || h-b == -1 {
					fmt.Printf("h=%d, b=%d, L=%d, m=%d, n=%d\n", h, b, L, m, n)
					//m=4, n=1 is a solution
				}
				/*
					a, b, c := m*m-n*n, 2*m*n, m*m+n*n
					//need 2*b - c = +-1
					//or 4mn - m^2-n^2 = +-1
					if (2*b-c == 1) || (2*b-c == -1) {
						fmt.Printf("(%d,%d,%d), m=%d, n=%d\n", a, b, c, m, n)
						//m=4, n=1 is a solution
					}
				*/
			}
		}
	}
}

func main() {
	// tripples()
	/*
		solve Diophantene equations (https://www.alpertron.com.ar/QUAD.HTM):
		L^2 = (b/2)^2 + h^2 = (b/2)^2 + (b+-1)^2
		5b^2 +- 8b + 4 - 4L^2 = 0

		One of solutions (for +8) gives
		b(n+1)=-9b(n)+8L(n)-8
		L(n+1)=10b(n)-9L(n)+8
		For -8 the solution is same, but with minus sign.

		So for sum need to take the abs value

		m^2-n^2-4mn+1 = 0
		m^2-n^2-4mn-1 = 0

		Both equations give alternating solutions in form
		m(i+1) = 17*m(i) + 4*n(i)
		n(i+1) = 4*m(i) +n(i)
		From tripples() call we know that (m, n) = (4, 1) is a solution
	*/
	b, L := int64(0), int64(1)
	b, L = -9*b+8*L-8, 10*b-9*L+8
	sumL := int64(0)
	for i := 1; i <= 12; i++ {
		b, L = -9*b+8*L-8, 10*b-9*L+8
		fmt.Printf("b=%d, L=%d\n", b, L)
		if L < 0 {
			sumL -= L
		} else {
			sumL += L
		}
	}
	fmt.Println(sumL)
	/*
		sumL := int64(0)
		m, n := int64(4), int64(1)
		for i := 1; i <= 12; i++ {
			a, b, c := m*m-n*n, 2*m*n, m*m+n*n
			fmt.Printf("(%d,%d,%d), m=%d, n=%d\n", a, b, c, m, n)
			// if 2*b-c != -1 {
			// fmt.Println("error!!!")
			// }
			sumL += c
			m, n = 17*m+4*n, 4*m+n
		}
		fmt.Println(sumL)
	*/
}
