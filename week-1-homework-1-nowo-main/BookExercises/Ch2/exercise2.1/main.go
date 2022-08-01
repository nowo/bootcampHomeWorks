package main

import (
	"example.com/tempConv"
	"fmt"
)

func main() {
	var kelvin tempConv.Kelvin = 100
	fmt.Println(tempConv.KToC(kelvin).String())
}
