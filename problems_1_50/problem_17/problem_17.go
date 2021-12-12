package main

import "fmt"

var oneDigit = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var teens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var decimal = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func spell(n int) string {
	if n == 1000 {
		return "onethousand"
	}
	if n < 10 {
		return oneDigit[n]
	}
	if n < 20 {
		return teens[n-10]
	}
	if n < 100 {
		return decimal[n/10] + oneDigit[n%10]
	}
	if n%100 == 0 {
		return oneDigit[n/100] + "hundred"
	}
	return oneDigit[n/100] + "hundred" + "and" + spell(n%100)
}

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		fmt.Printf("%d - %s - %d\n", i, spell(i), len(spell(i)))
		sum += len(spell(i))
	}
	fmt.Println(sum)
}
