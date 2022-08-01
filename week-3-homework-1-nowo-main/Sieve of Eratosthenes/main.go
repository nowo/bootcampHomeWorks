package main

import "fmt"

func main() {
	fmt.Println(findPrimeNumbers(10))
}

//this func take int argument to find prime numbers till that number
func findPrimeNumbers(n int) []int {
	var primeNumbers []int
	var isPrime []bool
	isPrime = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}
