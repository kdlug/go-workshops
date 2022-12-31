package main

import "fmt"

func main() {
	// var books [4]string
	prev := [...]string{
		"Cars",
		"Frozen",
		"Lion King",
		"Cars Part 3",
	}

	// cannot assign - different types
	// we can only assign to a variable which is the same type
	// var books [3]string
	// books = prev

	books := prev // values are copied; it allocates a new memory space

	for i := range prev {
		books[i] += " 2 edition" // it updates only books array
	}

	fmt.Printf("books : %#v\n", books)
	fmt.Printf("prev : %#v\n", prev)

	// OUTPUT:
	// books : [4]string{"Cars 2 edition", "Frozen 2 edition", "Lion King 2 edition", "Cars Part 3 2 edition"}
	// prev : [4]string{"Cars", "Frozen", "Lion King", "Cars Part 3"}
}
