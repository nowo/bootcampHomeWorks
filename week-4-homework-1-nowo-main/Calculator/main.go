package main

import (
	"errors"
	"fmt"
	mathFunction "mathFunction/math"
	"os"
	"strings"
)

var flag bool = true

type Calculator struct {
	functions []mathFunction.MathFunction
}

func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		//Check if the name of the function is same as the name entered by user and make lowercase
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {
	startCalculator()
}

func startCalculator() {
	myCalculator := Calculator{}

	//Define your own functions here
	myCalculator.addMathFunction(mathFunction.Sin{"Sin"})
	myCalculator.addMathFunction(mathFunction.Cos{"Cos"})
	myCalculator.addMathFunction(mathFunction.Log{"Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	//Start the calculator for checking flag
	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter q to exit:")
		fmt.Print("Select one of the functions: ")
		for _, v := range myCalculator.functions {
			fmt.Printf("%s ", v.GetName())
		}
		fmt.Println("") //new line
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		//Check if user wants to exit
		if fName == "q" {
			flag = false
			fmt.Println("\nCalculator stopped.")
			return
		}
		fmt.Println("> Enter a value for the calculation:")
		_, err = fmt.Scanf("%f", &arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		value, err := myCalculator.doCalculation(fName, arg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
		}

	}
}
