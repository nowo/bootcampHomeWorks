package main

//exercise 1
import (
	_ "packageExample.com/package"
	//exercise 3
	differantPackage "packageExample.com/package"
)

func main() {
	//exercise 1
	Greet()
	bye()
	differantPackage.SayErdal()
}
