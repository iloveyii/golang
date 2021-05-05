package main

import "fmt"


func main() {
	// By value
	fmt.Println("Greeting :: ", greeting("Ali"))
	fmt.Println("Add two numbers :: ", add( 2, 3))

	// By ref
	number := 10
	increment(&number)
	fmt.Println("Increment a number :: ", number)
}

func greeting(name string) string {
	return "Welcome back " + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func increment(number *int) {
	(*number)++
	// *number = *number + 1
}