package main

import "fmt"

func main() {

	// 6
	type (
		integer  int
		bookcase [5]int
		cabinet  [5]integer
	)
	red := cabinet{10, 4, 2, 6, 7}
	blue := bookcase{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	// ERROR: cannot convert blue (variable of type bookcase) to type cabinet
	// different underlying types [5]int!= [5]integer
	if red == cabinet(blue) {
		fmt.Println("✅")
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
