package main

import (
	"fmt"
	"math"
)

/*
Shape interface with Area and Perimeter functions
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

/*
Rectangle
*/
type Rectangle struct {
	length, breadth float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

/*
Circle
*/
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var shapes [2]Shape
	shapes[0] = Circle{7}
	shapes[1] = Rectangle{10, 20}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter: ", shape.Perimeter())
	}
}
