package main

import "fmt"

func main() {
	fmt.Println(primeFactors(315))
}

//this func take int argument to find prime factors till that number
func primeFactors(n int) []int {
	var primeFactors []int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n = n / i
		}
	}
	return primeFactors
}
