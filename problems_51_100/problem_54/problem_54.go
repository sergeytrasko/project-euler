package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var suits = []string{"H", "C", "S", "D"}

type fn = func([]string) (bool, []int)

func cardRank(c string) int {
	for i := 0; i < len(ranks); i++ {
		if c[0:1] == ranks[i] {
			return i
		}
	}
	return -1
}

func cardSuit(c string) int {
	suit := c[1:2]
	for i := 0; i < len(suits); i++ {
		if suit == suits[i] {
			return i
		}
	}
	return -1
}

func findSame(cards []string, n int, expectedCount int) (bool, []int) {
	r := [13]int{}
	for i := 0; i < len(cards); i++ {
		r[cardRank(cards[i])]++
	}
	rr := []int{}
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] == n {
			rr = append(rr, i)
		}
	}
	found := len(rr) == expectedCount
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] != n && r[i] > 0 {
			rr = append(rr, i)
		}
	}
	return found, rr
}

func allSameSuit(cards []string) bool {
	s := [4]int{}
	for i := 0; i < len(cards); i++ {
		s[cardSuit(cards[i])]++
	}
	for i := 0; i < len(s); i++ {
		if s[i] == len(cards) {
			return true
		}
	}
	return false
}

func highestCard(cards []string) []int {
	r := [13]int{}
	for i := 0; i < len(cards); i++ {
		r[cardRank(cards[i])]++
	}
	rr := []int{}
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] > 0 {
			rr = append(rr, i)
		}
	}

	return rr
}

func isHighestCard(cards []string) (bool, []int) {
	return true, highestCard(cards)
}

func isPair(cards []string) (bool, []int) {
	return findSame(cards, 2, 1)
}

func isTwoPairs(cards []string) (bool, []int) {
	return findSame(cards, 2, 2)
}

func isThreeOfAKind(cards []string) (bool, []int) {
	return findSame(cards, 3, 1)
}

func isStraignt(cards []string) (bool, []int) {
	r := [13]int{}
	for i := 0; i < len(cards); i++ {
		r[cardRank(cards[i])]++
	}
	for i := 0; i < len(r)-len(cards); i++ {
		ok := true
		for j := 0; j < len(cards); j++ {
			if r[i+j] != 1 {
				ok = false
				break
			}
		}
		if ok {
			return true, []int{i}
		}
	}
	return false, []int{}
}

func isFlush(cards []string) (bool, []int) {
	return allSameSuit(cards), highestCard(cards)
}

func isFullHouse(cards []string) (bool, []int) {
	p, _ := isPair(cards)
	t, rank := isThreeOfAKind(cards)
	return p && t, rank
}

func isFourOfAKind(cards []string) (bool, []int) {
	return findSame(cards, 4, 1)
}

func isStraightFlush(cards []string) (bool, []int) {
	s, r := isStraignt(cards)
	f, _ := isFlush(cards)
	return s && f, r
}

func isRoyalFlush(cards []string) (bool, []int) {
	f, r := isStraightFlush(cards)
	return f && r[0] == cardRank("TX"), r
}

func bestHand(cards []string) []int {
	compare := []fn{
		isHighestCard,
		isPair,
		isTwoPairs,
		isThreeOfAKind,
		isStraignt,
		isFlush,
		isFullHouse,
		isFourOfAKind,
		isStraightFlush,
		isRoyalFlush,
	}
	for i := len(compare) - 1; i >= 0; i-- {
		res, rank := compare[i](cards)
		if res {
			return append([]int{i}, rank...)
		}
	}
	return []int{}
}

func beats(c1 []string, c2 []string) int {
	// combinations := []string{
	// 	"isHighestCard",
	// 	"isPair",
	// 	"isTwoPairs",
	// 	"isThreeOfAKind",
	// 	"isStraignt",
	// 	"isFlush",
	// 	"isFullHouse",
	// 	"isFourOfAKind",
	// 	"isStraightFlush",
	// 	"isRoyalFlush",
	// }
	b1, b2 := bestHand(c1), bestHand(c2)
	// fmt.Printf("%s = %d (%s), %s = %d (%s)\n", c1, b1, combinations[b1[0]], c2, b2, combinations[b2[0]])
	for i := 0; i < len(b1) && i < len(b2); i++ {
		if b1[i] > b2[i] {
			// fmt.Println("Player 1 wins")
			return 1
		} else if b2[i] > b1[i] {
			// fmt.Println("Player 2 wins")
			return 2
		}
	}
	fmt.Println("Draw")
	return 0
}

func parse(s string) ([]string, []string) {
	x := strings.Split(s, " ")
	return x[0:5], x[5:]
}

func main() {
	file, _ := os.Open("in")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	player1Wins := 0
	for scanner.Scan() {
		line := scanner.Text()
		if beats(parse(line)) == 1 {
			player1Wins++
		}
	}
	fmt.Println(player1Wins)
}
