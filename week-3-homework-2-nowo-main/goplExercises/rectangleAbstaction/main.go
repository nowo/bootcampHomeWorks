package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func main() {
	rectangle, err := createRectangle(10, 20)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(calculateArea(rectangle))
	fmt.Println(calculateCircumference(rectangle))
}

// calculateArea calculates the area of a rectangle
func calculateArea(rectangle Rectangle) int {
	return rectangle.width * rectangle.height
}

// calculateCircumference calculates the circumference of a rectangle
func calculateCircumference(rectangle Rectangle) int {
	return 2 * (rectangle.width + rectangle.height)
}

// createRectangle creates a rectangle and checks if the width and height are positive
func createRectangle(width int, height int) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, errors.New("Invalid rectangle")
	}
	return Rectangle{width, height}, nil
}
