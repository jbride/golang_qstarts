package main

// https://go.dev/tour/methods/4
// https://gobyexample.com/interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// implement geometry as a "rect" type
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implement geometry as a "circle" type
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// provide measurements on "geometry" types
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	fmt.Printf("main()")
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
