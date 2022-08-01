package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{}
	// call goroutines
	go appendOne(&numbers)
	go appendTwo(&numbers)
	go appendThree(&numbers)
	go appendTwo(&numbers)
	time.Sleep(30 * time.Millisecond)
}

func appendTwo(numbers *[]int) {
	// add 2 to numbers and print before and after the append
	fmt.Printf("numbers %v\n", numbers)
	*numbers = append(*numbers, 2)
	fmt.Printf("numbers %v\n", numbers)
}

func appendOne(numbers *[]int) {
	// add 1 to numbers
	*numbers = append(*numbers, 1)
}

func appendThree(numbers *[]int) {
	// add 3 to numbers
	*numbers = append(*numbers, 3)
}
