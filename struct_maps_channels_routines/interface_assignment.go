package main

import "fmt"

type square struct {
	name string
	sideLength float64
}

type triangle struct {
	name string
	base float64
	height float64
}

type shape interface {
	getName() string
	getArea() float64
}

func main() {
	t := triangle{
		name: "Triangle",
		base:   7,
		height: 10,
	}
	printArea(t)

	s := square{
		name: "Square",
		sideLength: 5,
	}
	printArea(s)
}

func printArea(s shape) {
	fmt.Printf("Area of %s is %0.2f\n", s.getName(), s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) getName() string {
	return t.name
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) getName() string {
	return s.name
}
