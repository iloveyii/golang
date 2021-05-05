package main

import "fmt"


func main() {

	a := 11
	b := 11

	if a < b {
		fmt.Printf("%d is less than %d \n", a, b)
	} else if a > b {
		fmt.Printf("%d is greater than %d \n", a, b)
	} else {
		fmt.Printf("%d is equal to %d \n", a, b)
	}


	grades := 5

	switch true {
		case grades > 50:
			fmt.Println("You are passed")
		case grades < 50:
			fmt.Println("You are failed")
		default:
			fmt.Println("You can retake exams")
	}
}