package main

import "fmt"

func main() {

	// 3
	type bookcase [5]int
	type cabinet [5]int

	red := cabinet{10, 4, 2, 6, 7}
	blue := bookcase{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	// we can't compare, because variables have different types
	if red == blue { // ERROR: invalid operation: red == blue (mismatched types cabinet and bookcase)
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
