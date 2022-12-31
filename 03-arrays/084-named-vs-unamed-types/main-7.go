package main

import "fmt"

func main() {

	// 6
	type (
		integer  int
		bookcase [5]int
		cabinet  [5]integer
	)
	red := [5]integer{10, 4, 2, 6, 7}
	blue := [5]int{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	// ERROR: invalid operation: red == blue (mismatched types [5]integer and [5]int)
	if red == blue {
		fmt.Println("✅")
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
