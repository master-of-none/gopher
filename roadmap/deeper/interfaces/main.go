package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	length float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (s square) area() float64 {
	return s.side * s.side
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{2.5}
	r1 := rectangle{3, 4}
	s1 := square{5}
	t1 := triangle{10, 5}

	shapes := []shape{c1, r1, s1, t1}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
