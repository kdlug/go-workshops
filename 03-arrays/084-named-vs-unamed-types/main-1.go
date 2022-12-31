package main

import "fmt"

func main() {

	// 1
	red := [5]int{10, 4, 2, 6, 7}
	blue := [5]int{10, 4, 2, 6, 7}

	fmt.Print("Are they equal? ")

	if red == blue {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("red: %v\n", red)
	fmt.Printf("blue: %v\n", blue)
}
