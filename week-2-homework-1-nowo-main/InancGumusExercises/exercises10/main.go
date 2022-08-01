package main

import "fmt"

func main() {
	//minutesInWeek()
	//removesTheMagic()
	//constantLength()
	tau()
}

func minutesInWeek() {
	misPerDay := 60 * 24
	weekDay := 7
	fmt.Println(2 * misPerDay * weekDay)
}

func removesTheMagic() {
	const hoursInDay int = 24
	const daysInWeek int = 7
	const hoursPerWeek int = hoursInDay * daysInWeek
	fmt.Printf("There are %d hours in %d weeks.\n", hoursPerWeek*5, 5)
}

func constantLength() {
	const home string = "Milky Way Galaxy"
	const length int = len(home)
	fmt.Printf("There are %d characters inside %s.\n", length, home)
}

func tau() {
	const pi = 3.14159265358979323846264
	const tau = pi * 2
	fmt.Println(tau)
}

func area() {
	const (
		width  int16 = 25
		height int32 = int32(width) * 2
	)

	fmt.Printf("area = %d\n", width*int16(height))
}
