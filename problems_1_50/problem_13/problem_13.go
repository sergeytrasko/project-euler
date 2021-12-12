package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	file, _ := os.Open("in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := big.NewInt(0)
	for scanner.Scan() {
		x, _ := big.NewInt(0).SetString(scanner.Text(), 10)
		sum = sum.Add(sum, x)
	}
	file.Close()
	s := sum.String()
	fmt.Println(s[0:10])
}
