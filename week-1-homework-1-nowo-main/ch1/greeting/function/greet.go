package main

import (
	"fmt"
)

func main() {
	greet("Erdal")
}

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}
