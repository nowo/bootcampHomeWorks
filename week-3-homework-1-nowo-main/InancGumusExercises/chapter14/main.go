package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//wizardPrinter()
	//currencyCovnerter()
	//hipsterLoveSearch()
	//avarange()
	//sorter()
	wordFinder()
}
func wizardPrinter() {
	names := [][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i := range names {
		n := names[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}
}
func currencyCovnerter() {
	const (
		EUR = iota
		GBP
		JPY
	)

	rates := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
	fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP]*amount)
	fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY]*amount)
}

func hipsterLoveSearch() {
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Tell me a book title")
		return
	}
	query := strings.ToLower(args[0])

	fmt.Println("Search Results:")

	var found bool
	for _, v := range books {
		if strings.Contains(strings.ToLower(v), query) {
			fmt.Println("+", v)
			found = true
		}
	}

	if !found {
		fmt.Printf("We don't have the book: %q\n", query)
	}
}

func avarange() {
	args := os.Args[1:]
	if l := len(args); l == 0 || l > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		sum   float64
		nums  [5]float64
		total float64
	)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		total++
		nums[i] = n
		sum += n
	}

	fmt.Println("Your numbers:", nums)
	fmt.Println("Average:", sum/total)
}

func sorter() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 0:
		fmt.Println("Please give me up to 5 numbers.")
		return
	case l > 5:
		fmt.Println("Sorry. Go arrays are fixed.",
			"So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	var nums [5]float64

	// fill the array with the numbers
	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			// skip if it's not a valid number
			continue
		}

		nums[i] = n
	}

	/*
		check whether it's the last element or not:
		  i < len(nums)-1
		check whether the next number is greater than the current one, if so, swap it:
		  v > nums[i+1]
	*/
	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	fmt.Println(nums)
}

func wordFinder() {
	corpus := "lazy cat jumps again and again and again since the beginning this was very important"

	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}

	filter := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	words := strings.Fields(corpus)

queries:
	for _, q := range query {
		q = strings.ToLower(q)

		for _, v := range filter {
			if q == v {
				continue queries
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}
