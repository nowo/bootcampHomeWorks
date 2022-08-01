package main

import (
	"fmt"
	"strconv"
)

func main() {
	//arithmeticOperations()
	//basicTypes()
	iotaFunction()
}

func arithmeticOperations() {
	fmt.Println(8 + 6)
	fmt.Println(8 - 6)
	fmt.Println(8 * 6)
	fmt.Println(8 / 6)
	fmt.Println(8 % 6)
	fmt.Println(8 * 6)
	fmt.Println(8 % 6)

	var a int = 8
	a++
	fmt.Println(a)
}

func basicTypes() {
	var a int = 8
	var b float64 = 8.5
	var c string = "hello"
	var d bool = true
	fmt.Println(a, b, c, d)
}

func conversion() {
	var a int = 8
	var b float64 = 8.5
	var c string = "8"
	var d bool = true
	fmt.Println(strconv.Atoi(c))
	fmt.Println(strconv.FormatFloat(b, 'f', 2, 64))
	fmt.Println(strconv.FormatBool(d))
	fmt.Println(strconv.Itoa(a))
}

func iotaFunction() {
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)
}
