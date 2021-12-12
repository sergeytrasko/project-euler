package main

import (
	"bufio"
	"fmt"
	"os"
)

var letters = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func toRoman(n int) string {
	s := ""
	for i := 0; i < len(letters); i++ {
		l, v := letters[i], values[i]
		for n >= v && n > 1 {
			s = s + l
			n = n - v
		}
	}
	if n == 1 {
		s += "I"
	}
	return s
}

func fromRoman(r string) int {
	n := 0
	for len(r) > 0 {
		for i := 0; i < len(letters); i++ {
			l, ll, v := letters[i], len(letters[i]), values[i]
			if len(r) >= ll && r[0:ll] == l {
				// fmt.Printf("%s - %s - %d\n", r, l, n)
				n += v
				r = r[ll:]
			}
		}
	}
	return n
}

func main() {
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	dif := 0
	for scanner.Scan() {
		line := scanner.Text()
		n := fromRoman(line)
		r := toRoman(n)
		improvement := len(line) - len(r)
		fmt.Printf("%s - %d - %s - %d\n", line, n, r, improvement)
		dif += improvement
	}
	fmt.Println(dif)
}
