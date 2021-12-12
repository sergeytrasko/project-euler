package main

import "fmt"

func hasDigit(n int, digit int) bool {
	for n > 0 {
		d := n % 10
		if d == digit {
			return true
		}
		n = n / 10
	}
	return false
}

func gcd(a int, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(b-a, a)
	}
}

func simplify(p int, q int) (int, int) {
	if p == 0 || q == 0 {
		return p, q
	}
	d := gcd(p, q)
	return p / d, q / d
}

func removeDigit(n int, digit int) int {
	if n%10 == digit {
		return n / 10
	} else {
		return n % 10
	}
}

func main() {
	mulP, mulQ := 1, 1
	for p := 10; p < 100; p++ {
		for q := p + 1; q < 100; q++ {
			sp, sq := simplify(p, q)
			for d := 1; d <= 9; d++ {
				if hasDigit(p, d) && hasDigit(q, d) {
					pp, qq := simplify(removeDigit(p, d), removeDigit(q, d))
					if sp == pp && sq == qq {
						fmt.Printf("%d/%d=%d/%d\n", p, q, removeDigit(p, d), removeDigit(q, d))
						mulP, mulQ = mulP*p, mulQ*q
					}
				}
			}
		}
	}
	mulP, mulQ = simplify(mulP, mulQ)
	fmt.Printf("%d/%d\n", mulP, mulQ)
}
