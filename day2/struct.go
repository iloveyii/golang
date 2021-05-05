package main

import "fmt"

type Rectangle struct {
	length int
	width int
}

func main() {
	r := Rectangle{length: 10, width: 20}
	fmt.Println(r)
}