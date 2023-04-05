package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

type triangle struct {
	side_a, side_b, side_c float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) area() float64 {
	s := t.perim() / 2
	return math.Sqrt(s*(s-t.side_a)*(s-t.side_b)*(s-t.side_c))
}

func (t triangle) perim() float64 {
	return t.side_a + t.side_b + t.side_c
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{height: 4, width: 5}
	c := circle{radius: 5}
	t := triangle{side_a: 10, side_b: 10, side_c: 10}

	measure(r)
	measure(c)
	measure(t)
}
