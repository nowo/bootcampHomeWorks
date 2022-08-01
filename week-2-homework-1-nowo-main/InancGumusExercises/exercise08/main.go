package main

import "fmt"

func main() {
	//exercise 1
	//exercise1()
	//fixPrintStatements()
	printwindowsPath()
}
func exercise1() {
	// exercise 1
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-5 + 2)
}

func fixPrintStatements() {
	fmt.Println(10 + 5 - (5 - 10))
	fmt.Println(-10 + (0.5 - (1 + 5.5)))
}

//stringExercises

func printwindowsPath() {
	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"
	fmt.Println(path)
}
