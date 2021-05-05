package main

import "fmt"


func main() {

	// Declare 
	var rolls [3]int

	// Assign
	rolls[0] = 10
	rolls[1] = 11
	rolls[2] = 12

	fmt.Println("Rolls :: ", rolls)

	// One liner - can ommit 2 in [2], use len(names)
	names := [2]string{"john", "doe"}
	fmt.Println("Names :: ", names, len(names))

	// Take a slice or subarr . start:0-indexed, stop:1-indexed, inclusive
	fmt.Println("Slice :: ", rolls[1:3])
	
}