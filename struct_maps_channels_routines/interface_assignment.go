package main

import "fmt"

type square struct {
	name       string
	sideLength float64
}

type triangle struct {
	name   string
	base   float64
	height float64
}

type shape interface {
	getArea() float64
	printArea()
}

func main() {
	t := triangle{
		name:   "Triangle",
		base:   7,
		height: 10,
	}

	s := square{
		name:       "Square",
		sideLength: 5,
	}

	shapes := []shape{t, s}
	for _, shape := range shapes {
		shape.printArea()
	}
}


func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) printArea() {
	fmt.Printf("Area of %s is %0.2f\n", t.name, t.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) printArea() {
	fmt.Printf("Area of %s is %0.2f\n", s.name, s.getArea())
}
