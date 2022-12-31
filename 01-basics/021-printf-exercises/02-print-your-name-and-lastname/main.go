package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier
	name := "John"
	surname := "Lennon"

	fmt.Printf("My name is %s and my lastname is %s.\n", name, surname)

	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, name, surname)
}
