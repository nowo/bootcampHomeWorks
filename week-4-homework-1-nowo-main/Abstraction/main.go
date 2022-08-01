package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var initialTemp float64
	var scale string
	fmt.Println("Enter the temperature in Celsius:")
	fmt.Scanf("%f", &initialTemp)
	fmt.Println("Enter the scale to convert to:")
	fmt.Scanf("%s", &scale)

	selectedScale, err := selectScale(scale)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("The converted temperature  is:", selectedScale.ConvertToAnotherScale(selectedScale.GetTemperature(), scale))

}

//SelectScale is a function that returns the scale type
func selectScale(scale string) (Temperature, error) {
	switch scale {
	case "C":
		return &Celsius{}, nil
	case "F":
		return &Fahrenheit{}, nil
	case "K":
		return &Kelvin{}, nil
	default:
		return nil, errors.New("no such scale exists:" + scale)
	}
}

//we create out interface
type Temperature interface {
	GetTemperature() float64
	Settemperature(float64)
	ConvertToAnotherScale(float64, string) float64
}

type Celsius struct {
	temp float64
}

//we get our temp
func (c Celsius) GetTemperature() float64 {
	return c.temp
}

//set temperature
func (c *Celsius) Settemperature(temp float64) {
	c.temp = temp
}

//convert to another scale
func (c Celsius) ConvertToAnotherScale(temp float64, scale string) float64 {
	if scale == "F" {
		return temp*9/5 + 32
	} else if scale == "K" {
		return temp + 273.15
	} else {
		return temp
	}
}

type Fahrenheit struct {
	temp float64
}

func (f Fahrenheit) GetTemperature() float64 {
	return f.temp
}
func (f *Fahrenheit) Settemperature(temp float64) {
	f.temp = temp
}
func (f Fahrenheit) ConvertToAnotherScale(temp float64, scale string) float64 {
	if scale == "C" {
		return (temp - 32) * 5 / 9
	} else if scale == "K" {
		return (temp + 459.67) * 5 / 9
	} else {
		return temp
	}
}

type Kelvin struct {
	temp float64
}

func (k Kelvin) GetTemperature() float64 {
	return k.temp
}
func (k *Kelvin) Settemperature(temp float64) {
	k.temp = temp
}
func (k Kelvin) ConvertToAnotherScale(temp float64, scale string) float64 {
	if scale == "C" {
		return temp - 273.15
	} else if scale == "F" {
		return temp*9/5 - 459.67
	} else {
		return temp
	}
}
