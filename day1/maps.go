package main

import "fmt"

func main() {
	// Declare : key:string, value: string
	var emails = make(map[string]string)
	// Assign
	emails["john"] = "john@email.com"
	emails["jane"] = "jane@email.com"
	fmt.Println(emails)


	// Declare and initialize
	grades := map[string]int{"john": 85, "jane": 91}
	fmt.Println(grades)
}