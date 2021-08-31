package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	base   float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.base * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.base + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Interface function
func printShape(s Shape) {
	fmt.Printf("The Area of %T is: %0.2f \n", s, s.Area())
	fmt.Printf("The Perimeter of %T is: %0.2f \n", s, s.Perimeter())
}

func main() {
	shapes := []Shape{
		Rectangle{base: 30, height: 12},
		Rectangle{base: 100, height: 13},
		Rectangle{base: 36, height: 27},
		Circle{radius: 30},
		Circle{radius: 10},
		Circle{radius: 150},
	}

	for _, v := range shapes {
		printShape(v)
		println("**********")
	}
}
