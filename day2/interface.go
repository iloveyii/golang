package main

import (
	"fmt"
	"math"
)


// Declare interface
type Shape interface {
	area() float64
}

// Declare struct 1
type Circle struct {
	radius float64
}

// Declare struct 2
type Rectangle struct {
	height, width float64
}

// Bind area to Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Bind area to Rectangle
func (r Rectangle) area() float64 {
	return r.height * r.width
}

// Find area for a Shape 
func getArea(s Shape) float64 {
	return s.area()
}


func main() {
	// Initialize circle
	circle := Circle{radius: 5}
	// Initialize rectangle
	rectangle := Rectangle{height: 3, width: 5}

	// Calculate area
	fmt.Printf("Area of circle :: %.2f\n", getArea(circle))
	fmt.Printf("Area of rectangle :: %.2f\n", getArea(rectangle))
}