package main

import "fmt"

func main() {
	//exercise1
	printLiterals()
	printHexes()
	printDeclerations()
	typeConversions()
}

func printLiterals() {
	fmt.Println(1)
	fmt.Println(1.0)
	fmt.Println(true)
	fmt.Println("hello")
	fmt.Println("türkçe karakter")
}

func printHexes() {
	//Print 0 to 9 in hexadecimal
	for i := 0; i < 10; i++ {
		fmt.Printf("%x\n", i)
	}
	//Print 10 to 15 in hexadecimal
	for i := 10; i < 16; i++ {
		fmt.Printf("%x\n", i)
	}
	fmt.Printf("%x\n", 17)
	fmt.Printf("%x\n", 25)
	fmt.Printf("%x\n", 50)
	fmt.Printf("%x\n", 100)
}

func printDeclerations() {
	var a int
	var b float64
	var c bool
	var d string
	fmt.Println(a, b, c, d)

	//var 3speed int
	//var !speed bool
	//var spe?ed string
	//var var int
	//var func int
	//var package int
	//fmt.Println(3speed, !speed, spe?ed, var, func, package)
}

func typeConversions() {
	a, b := 10, 5.5
	fmt.Println(a + int(b))
}
