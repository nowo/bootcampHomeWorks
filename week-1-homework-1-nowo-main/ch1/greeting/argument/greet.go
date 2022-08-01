package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
}

func createGreet(name string) string {
	return "Selam " + name + " :)"
}
