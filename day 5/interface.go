package main

import (
	"fmt"
	"math"
)

type geomatry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type cricle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c cricle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c cricle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geomatry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := cricle{radius: 5}

	measure(r)
	measure(c)
}
