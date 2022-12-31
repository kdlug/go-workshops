package main

import "fmt"

func main() {

	// 2
	type bookcase [5]int
	red := [5]int{10, 4, 2, 6, 7}
	blue := bookcase{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	if red == blue {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
