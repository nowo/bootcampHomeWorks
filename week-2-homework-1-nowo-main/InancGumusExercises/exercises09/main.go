package main

import (
	"fmt"
)

func main() {
	typeProblem()
}

func typeProblem() {
	var (
		width  uint16
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	fmt.Printf("are they equal? %t", width == height)
}
