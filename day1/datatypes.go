package main

import "fmt"

// Glabal access
var data int = 50
// short notation cannot here - non declaration statement outside function body
// data2 := 150

func main() {
	// declare a variable with var - can specify exact data type
	var id int16 = 16
	var name string = "Ali"
	const pi float32 = 3.14

		fmt.Println(id, name)
	fmt.Printf("Data types: %T, %T\n", id, name)

	// derivative - with default types - short notation
	id2 := 32
	name2 := "Bob"
	fmt.Println(id2, name2)
	fmt.Printf("Data types: %T, %T\n", id2, name2)

	fmt.Println("Global access ", data)

	fmt.Println("I am a constant ", pi)
	// cannot assign to pi declared const
	// pi = 3

	// shorter version of declaration
	student_name, email := "Doe" , "john@doe.com"
	fmt.Println(student_name, email)
} 