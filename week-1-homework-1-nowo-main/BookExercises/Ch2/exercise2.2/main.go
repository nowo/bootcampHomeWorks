package main

import (
	"fmt"
)

func KToC(k float64) float64 {
	return (k - 273.15)
}
func Ctof(c float64) float64 {
	return (c*9/5 + 32)
}

var generalOptions = map[string]int{
	"temperature": 0,
	"lengh":       1,
	"weight":      2,
}

func printOptions(option map[string]int) {
	for key, value := range option {
		fmt.Println(key, value)
	}
}
func main() {
	var number int

	printOptions(generalOptions)
	fmt.Printf("Enter your option: ")
	fmt.Scanf("%d", &number)

	switch number {
	case generalOptions["temperature"]:
		fmt.Print("Enter your kelvin temperature: ")
		var kelvin int
		fmt.Scanf("%d", &kelvin)
		fmt.Printf("%d kelvin is %f celsius\n", kelvin, KToC(float64(kelvin)))
		fmt.Printf("%d kelvin is %f fahrenheit\n", kelvin, Ctof(KToC(float64(kelvin))))
	default:
		fmt.Println("Not available right now")
	}

}
