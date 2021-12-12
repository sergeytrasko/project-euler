package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func choose(s []int, taken []int, p int) bool {
	if p == len(s) {
		a, b := []int{}, []int{}
		for i := 0; i < len(s); i++ {
			if taken[i] == 1 {
				a = append(a, s[i])
			} else if taken[i] == 2 {
				b = append(b, s[i])
			}
		}
		if len(a) == 0 || len(b) == 0 {
			return true
		}
		sa, sb := sum(a), sum(b)
		// fmt.Printf("A:%d (sum = %d), B:%d (sum=%d)\n", a, sa, b, sb)
		if sa == sb {
			return false
		}
		if len(a) > len(b) {
			return sa > sb
		}
		if len(b) > len(a) {
			return sb > sa
		}
		return true
	}
	if !choose(s, taken, p+1) {
		return false
	}
	taken[p] = 1
	if !choose(s, taken, p+1) {
		return false
	}
	taken[p] = 2
	if !choose(s, taken, p+1) {
		return false
	}
	taken[p] = 0
	return true
}

func isSpecialSumSet(s []int) bool {
	taken := make([]int, len(s))
	return choose(s, taken, 0)
}

func sum(s []int) int {
	ss := 0
	for _, r := range s {
		ss += r
	}
	return ss
}

func main() {
	res := 0
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		s := []int{}
		for j := 0; j < len(tokens); j++ {
			num, _ := strconv.Atoi(tokens[j])
			s = append(s, num)
		}
		if isSpecialSumSet(s) {
			fmt.Println(s)
			res += sum(s)
		}
	}
	fmt.Println(res)
}
