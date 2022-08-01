package main

import (
	"fmt"
	"strings"
)

func main() {
	greetPrinter(createGreetInTurkish, "Erdal")
	greetPrinter(createGreetInEnglish, "Erdal")
	greetPrinter(convertToUppercase, "erdal")

	greetCreator := createGreetInTurkish
	greetPrinter(greetCreator, "Erdal")

	func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}("Erdal")

	closure := func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}
	closure("Erdal")
	anotherGreetPrinter(closure, "Burak")
}

func createGreetInTurkish(name string) string {
	return "Selam " + name + " :)"
}

func createGreetInEnglish(name string) string {
	return "Hi " + name + " :)"
}

func convertToUppercase(arg string) string {
	return strings.ToUpper(arg)
}

func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}

func anotherGreetPrinter(function func(it string), name string) {
	function(name)
}
