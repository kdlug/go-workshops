package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	args := os.Args
	count := len(args) - 1

	if count < 3 {
		fmt.Println("Please provide 3 arguments")
		os.Exit(1)
	}

	p1, p2, p3 := args[1], args[2], args[3]

	fmt.Println("There are", count, "people!")
	fmt.Println("Hello great", p1, "!")
	fmt.Println("Hello great", p2, "!")
	fmt.Println("Hello great", p3, "!")
	fmt.Println("Nice to meet you all.")
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}
