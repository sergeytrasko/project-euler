package main

import (
	"fmt"
	"math"
)

func findIntercectionX(n float64) float64 {
	a := 1 + 1/(n*n)
	b := -(1 + 1/n)
	c := 0.25
	return (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
}

func traingeArea(x float64, n float64) float64 {
	h := x / n
	return h * x / 2
}

//https://www.wolframalpha.com/input/?i2d=true&i=Integrate%5BDivide%5B1%2C2%5D-sqrt%5C%2840%29Divide%5B1%2C4%5D-Power%5B%5C%2840%29x-Divide%5B1%2C2%5D%5C%2841%29%2C2%5D%5C%2841%29%2Cx%5D
func integral(x float64) float64 {
	return 0.25 * (-2*math.Sqrt(x-x*x)*x + math.Sqrt(x-x*x) + 2*x + math.Asin(math.Sqrt(1-x)))
	// return -0.25 * (-2*x + math.Sqrt(x*(1-x))*(2*x-1) - math.Asin(math.Sqrt(1-x)))
}

func conclaveTriangleArea(x float64) float64 {
	return integral(0.5) - integral(x)
}

func ratio(n int) float64 {
	fmt.Printf("n=%d\n", n)
	total := (1 - math.Pi/4) / 4
	x := findIntercectionX(float64(n))
	// fmt.Printf("total=%f\n", total)
	// fmt.Printf("integ=%f\n", conclaveTriangleArea(0))
	// fmt.Printf("x=%f\n", x)
	// fmt.Printf("triangleArea=%f\n", traingeArea(x, n))
	// fmt.Printf("conclaveArea=%f\n", conclaveTriangleArea(x))
	area := traingeArea(x, float64(n)) + conclaveTriangleArea(x)
	// fmt.Printf("Area=%f\n", area)
	fmt.Printf("Ratio = %.10f\n", area/total)
	return area / total
}

/*
line: y = x/n
circle: (x-1/2)^2 + (y-1/2)^2 = 1
*/
func main() {
	n := 1
	for ; ; n++ {
		if ratio(n) < 0.001 {
			break
		}
	}
	fmt.Println(n)
}
