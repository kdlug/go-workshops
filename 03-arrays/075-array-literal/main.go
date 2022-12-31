package main

import "fmt"

func main() {
	// var books [4]string

	// books[0] = "Cars"
	// books[1] = "Frozen"
	// books[2] = "Lion King"
	// books[3] = books[0] + " Part 3"

	// 1. single line array list
	// var books [4]string = {"Cars", "Frozen", "Lion King", "Lion King Part 3"} // we dont neet a comma , after the last element

	// 2. multiple lines array list
	// var books [4]string = {
	// 	"Cars",
	// 	"Frozen",
	// 	"Lion King",
	// 	"Lion King Part 3", // comma is required
	// 	}

	// 3. shortcut
	// books := [4]string{ // type is needed for arrays
	// 	"Cars",
	// 	"Frozen",
	// 	"Lion King",
	// 	"Lion King Part 3", // comma is required
	// 	}

	// 4. ellipsis [...]
	books := [...]string{ // calculates the len of array automatically based on the array list
		"Cars",
		"Frozen",
		"Lion King",
		"Lion King Part 3", // comma is required
	}

	fmt.Printf("books : %#v\n", books)
}
