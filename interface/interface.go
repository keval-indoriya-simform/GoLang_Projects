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
	sideA, sideB, sideC float64
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
	return math.Sqrt(s * (s - t.sideA) * (s - t.sideB) * (s - t.sideC))
}

func (t triangle) perim() float64 {
	return t.sideA + t.sideB + t.sideC
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{height: 4, width: 5}
	c := circle{radius: 5}
	t := triangle{sideA: 10, sideB: 10, sideC: 10}

	measure(r)
	measure(c)
	measure(t)
}
