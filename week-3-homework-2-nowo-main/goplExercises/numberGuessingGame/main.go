package main

import (
	"fmt"
	"math/rand"
	"time"
)

var isGameFinished bool

func main() {
	var guessedNumber int
	var guessedNumberList []int
	randomNumber := getRandomNumber()

	for isGameFinished == false {
		fmt.Println("Enter your number: ")
		fmt.Scanf("%d", &guessedNumber)
		fmt.Println(compareGuessedNumberWithRandomNumber(guessedNumber, randomNumber))
		guessedNumberList = append(guessedNumberList, guessedNumber)
	}
	fmt.Println("Your guessed numbers are: ", guessedNumberList)
}

func compareGuessedNumberWithRandomNumber(guessedNumber int, randomNumber int) string {

	var guessedCorrectNumberWithoutOrder int
	var guessedCorrectNumberWithOrder int
	var guessedCorrectNumberWithoutOrderString string
	var guessedCorrectNumberWithOrderString string

	//finish game if user guessed the number
	if guessedNumber == randomNumber {
		isGameFinished = true
		return "+4"

	}
	guessedNumberSlice := IntToSlice(guessedNumber, []int{})
	randomNumberSlice := IntToSlice(randomNumber, []int{})

	for index, val := range guessedNumberSlice {
		for index2, val2 := range randomNumberSlice {
			if val == val2 {
				//check if number is in correct order
				if index == index2 {
					guessedCorrectNumberWithOrder++
				} else {
					//check if number is in correct order but not in the same position
					guessedCorrectNumberWithoutOrder++
				}
			}
		}
	}
	if guessedCorrectNumberWithoutOrder != 0 {
		guessedCorrectNumberWithoutOrderString = fmt.Sprintf("-%d", guessedCorrectNumberWithoutOrder)
	}
	if guessedCorrectNumberWithOrder > 0 {
		guessedCorrectNumberWithOrderString = fmt.Sprintf("+%d", guessedCorrectNumberWithOrder)
	}
	return guessedCorrectNumberWithOrderString + guessedCorrectNumberWithoutOrderString
}

//get random number but number should be 4 digits and do not have repeating digits
func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	randomNumer := rand.Intn(9999)
	for checkRandomNumberEquility(randomNumer) == false || randomNumer < 1000 {
		randomNumer = rand.Intn(9999)
	}

	return randomNumer
}

//check if number has repeating digits
func checkRandomNumberEquility(randomNumber int) bool {
	slicedInt := IntToSlice(int(randomNumber), []int{})
	if slicedInt[0] == 0 {
		return false
	}
	for index, _ := range slicedInt {
		if index != len(slicedInt)-1 {
			if slicedInt[index] == slicedInt[index+1] {
				return false
			}
		}
	}
	return true
}

//convert int to slice of int
func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}
