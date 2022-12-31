package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func main() {
	// Change this accordingly to produce the expected outputs
	age := 350

	// Type your if statement here.
	if age > 60 {
		fmt.Printf("Getting older\n")
	} else if age > 30 {
		fmt.Printf("Getting wiser\n")
	} else if age > 20 {
		fmt.Printf("Adulthood\n")
	} else if age > 10 {
		fmt.Printf("Young blood\n")
	} else {
		fmt.Printf("Booting up\n")
	}
}
