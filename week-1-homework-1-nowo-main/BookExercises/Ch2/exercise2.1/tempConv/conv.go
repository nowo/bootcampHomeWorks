package tempConv

import "fmt"

func Ctof(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func Ftoc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func KToC(k Kelvin) Celsius {
	return Celsius((k - 273.15))
}

func main() {
	var c Celsius = 100
	fmt.Println(c.String())
}
