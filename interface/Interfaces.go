package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type recta struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r recta) area() float64 {
	return r.width * r.height
}

func (r recta) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)        // Prints the struct values
	fmt.Println(g.area()) //  Calls the correct method based on the type
	fmt.Println(g.perim())

}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok { // type assertion with two return values:
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {

	r := recta{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) // Calls measure() with a rectangle
	measure(c) // Calls measure() with a circle

	detectCircle(r)
	detectCircle(c)
}
