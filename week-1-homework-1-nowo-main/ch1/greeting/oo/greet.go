package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam " + p.name + " :)"
}

func main() {
	var greeter Person = Person{"Erdal"}
	greeting := greeter.greet()
	fmt.Printf("%s\n", greeting)
}
