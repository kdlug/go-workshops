package main

import "fmt"

func main() {

	// 5
	type bookcase [5]int
	type cabinet [10]int // different underlying type

	red := cabinet{10, 4, 2, 6, 7}
	blue := bookcase{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	// convert
	if red == cabinet(blue) { // ERROR: cannot convert blue (variable of type bookcase) to type cabinet
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
