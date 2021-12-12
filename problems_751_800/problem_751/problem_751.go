package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func b(x string) string {
	f, _ := strconv.ParseFloat(x, 64)
	res := math.Floor(f) * (f - math.Floor(f) + 1)
	return strconv.FormatFloat(res, 'f', -1, 64)
}

func a(b string) int {
	s := strings.Split(b, ".")[0]
	i, _ := strconv.Atoi(s)
	return i
}

func sequence(theta string, l int) string {
	f, _ := strconv.ParseFloat(theta, 64)
	floor := int(math.Floor(f))
	tau := fmt.Sprintf("%d.", floor)
	bn := theta
	for len(tau) < l {
		bn = b(bn)
		an := a(bn)
		tau = tau + fmt.Sprintf("%d", an)
	}
	return tau[:l]
}

func main() {
	theta := "2."
	for len(theta) < 2+24 {
		for i := 0; i < 10; i++ {
			t0 := theta + fmt.Sprintf("%d", i)
			tau := sequence(t0, len(t0))
			if t0 == tau {
				theta = theta + fmt.Sprintf("%d", i)
				fmt.Printf("%d - %s\n", len(theta), theta)
				break
			}
		}
	}
}
