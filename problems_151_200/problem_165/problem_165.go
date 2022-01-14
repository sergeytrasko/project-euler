package main

import (
	"fmt"
)

type point = struct {
	x, y float64
}

type segment = struct {
	a, b point
}

func intercects(s1, s2 segment) (bool, point) {
	s1_x := s1.b.x - s1.a.x
	s1_y := s1.b.y - s1.a.y
	s2_x := s2.b.x - s2.a.x
	s2_y := s2.b.y - s2.a.y

	s := (-s1_y*(s1.a.x-s2.a.x) + s1_x*(s1.a.y-s2.a.y)) / (-s2_x*s1_y + s1_x*s2_y)
	t := (s2_x*(s1.a.y-s2.a.y) - s2_y*(s1.a.x-s2.a.x)) / (-s2_x*s1_y + s1_x*s2_y)

	if s > 0 && s < 1 && t > 0 && t < 1 {
		// Collision detected
		return true, point{x: s1.a.x + (t * s1_x), y: s1.a.y + (t * s1_y)}
	}

	return false, point{} // No collision
}

func nextS(s int64) int64 {
	return (s * s) % 50515093
}

func main() {
	seg := []segment{}
	s := int64(290797)
	for i := 1; i <= 5000; i++ {
		s = nextS(s)
		t1 := s % 500
		s = nextS(s)
		t2 := s % 500
		s = nextS(s)
		t3 := s % 500
		s = nextS(s)
		t4 := s % 500
		seg = append(seg, segment{
			a: point{x: float64(t1), y: float64(t2)},
			b: point{x: float64(t3), y: float64(t4)},
		})
	}
	uniqueIntercections := map[string]bool{}
	for i := 0; i < len(seg); i++ {
		s1 := seg[i]
		for j := i + 1; j < len(seg); j++ {
			s2 := seg[j]
			ok, p := intercects(s1, s2)
			if ok {
				r := fmt.Sprintf("%0.12f_%0.12f", p.x, p.y)
				uniqueIntercections[r] = true
			}
		}
	}
	fmt.Println(len(uniqueIntercections))
}
