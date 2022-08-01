package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rotate(s, 2)

}

//this func take slice and rotate the numbers to left or right base on the given number
func rotate(slice []int, rotateRange int) {
	//if rotateRange is greater than length of slice then thake rotateRange % length of slice
	if rotateRange > len(slice) {
		rotateRange = rotateRange % len(slice)
	}
	if rotateRange < 0 {
		rotateRange = len(slice) + rotateRange
	}
	slice = append(slice[rotateRange:], slice[:rotateRange]...)
	fmt.Println(slice)
}
