package main

import "fmt"

// Declare
type Rectangle struct {
	length float32
	width float32
}

// Add a function to struct
func (r Rectangle) area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{length: 10, width: 20}
	fmt.Println(" Area of ", r, " is :: ", r.area())
}