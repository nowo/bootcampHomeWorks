// All semicolons are redundant if they are not used in for!
package main

import "fmt"

func main() {
	var aString string = "Selam"
	fmt.Printf("%s \n", aString)

	anInteger := 5
	fmt.Printf("%d\n", anInteger)
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
	return
}
