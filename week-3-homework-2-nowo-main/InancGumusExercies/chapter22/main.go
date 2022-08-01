package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	warmUp()
	populate()
	students()
}

//exercise 1
func warmUp() {
	var (
		phones map[string]string
		// Key        : Last name
		// Element    : Phone number

		// Key        : Product ID
		// Element    : Available / Unavailable
		products map[int]bool

		multiPhones map[string][]string
		// Key        : Last name
		// Element    : Phone numbers

		basket map[int]map[int]int
		// Key        : Customer ID
		// Element Key:
		//   Key: Product ID Element: Quantity
	)

	fmt.Printf("phones     : %#v\n", phones)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", multiPhones)
	fmt.Printf("basket     : %#v\n", basket)
}

//exercise 2

func populate() {
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}

	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	// Print dulin's phone number.
	who, phone := "dulin", "N/A"
	if v, ok := phones[who]; ok {
		phone = v
	}
	fmt.Printf("%s's phone number: %s\n", who, phone)

	// Is Product ID 879401371 available?
	id, status := 879401371, "available"
	if !products[id] {
		status = "not " + status
	}
	fmt.Printf("Product ID #%d is %s\n", id, status)

	// What is Greco's second phone number?
	who, phone = "greco", "N/A"
	if phones := multiPhones[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's 2nd phone number: %s\n", who, phone)

	// How many of 576872813 the customer 101 is going to buy?
	cid, pid := 101, 576872813
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, basket[cid][pid], pid)
}

//exercise 3
func students() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// remove the "bobo" house
	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house, students := args[0], houses[args[0]]
	if students == nil {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	// only sort the clone
	clone := append([]string(nil), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}
}
