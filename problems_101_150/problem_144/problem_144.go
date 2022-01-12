package main

import (
	"fmt"
	"math"
)

func main() {
	x1, y1, x2, y2 := 0.0, 10.1, 1.4, -9.6

	reflections := 0
	for math.Abs(x2) > 0.01 || y2 < 0 {
		fmt.Printf("%d - Beam goes from (%0.5f,%0.5f) to (%0.5f,%0.5f)\n", reflections, x1, y1, x2, y2)
		// A - angle of the beam
		// B - angle of tangent line
		tanA := -(y2 - y1) / (x2 - x1)
		tanB := -4 * x2 / y2
		tanAB := (tanA + tanB) / (1 - tanA*tanB)
		fmt.Printf("tan(A)=%0.5f, tan(B)=%0.5f, tan(A+B)=%0.5f\n", tanA, tanB, tanAB)
		fmt.Printf("A=%0.5f, B=%0.5f, (A+B)=%0.5f\n", math.Atan(tanA)*180/math.Pi, math.Atan(tanB)*180/math.Pi, math.Atan(tanAB)*180/math.Pi)
		r := (tanAB + tanB) / (1 - tanAB*tanB)
		bb := y2 - x2*r
		fmt.Printf("line y = %0.5f*x + %0.5f\n", r, bb)
		// fmt.Printf("%0.5f = y = %0.5f*%0.5f + %0.5f\n", y2, r, x2, bb)
		// solve y = r*x + b and 4x^2 + y^2 = 100
		// 4x^2 + (r*x + b) = 100
		// 4x^2 + r^2*x^2 + 2*r*x*b + b^2 - 100 = 0
		// (4+r^2)x^2 + (2r*b)*x + (b^2 - 100)
		a, b, c := 4+r*r, 2*r*bb, bb*bb-100
		d := b*b - 4*a*c

		xx1, xx2 := (-b+math.Sqrt(d))/(2*a), (-b-math.Sqrt(d))/(2*a)
		fmt.Printf("x1 = %0.5f, x2 = %0.5f\n", xx1, xx2)
		nextX := 0.0
		if math.Abs(x2-xx1) < math.Abs(x2-xx2) {
			nextX = xx2
		} else {
			nextX = xx1
		}
		x1, y1 = x2, y2
		x2, y2 = nextX, nextX*r+bb
		fmt.Printf("new reflection to (%0.5f,%0.5f)\n", x2, y2)
		fmt.Println()
		reflections++
	}
	fmt.Println(reflections)
}
