package main

import (
	"fmt"
	"math"
)

// interfaces : is used define the common behaviour of objects .

// interface is the list of methods that are going to describe behaviour of particular types

type geometry interface {
	area() float64
	peri() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// methods
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) peri() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.peri())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
