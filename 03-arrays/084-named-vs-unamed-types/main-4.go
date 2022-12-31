package main

import "fmt"

func main() {

	// 4
	type bookcase [5]int
	type cabinet [5]int

	red := cabinet{10, 4, 2, 6, 7}
	blue := bookcase{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	// convert
	if red == cabinet(blue) {
		fmt.Println("✅") // now they are equal
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
