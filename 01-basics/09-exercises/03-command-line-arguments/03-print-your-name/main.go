package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func main() {
	// get your name from the command-line
	// and print it
	name := os.Args[1]
	fmt.Println(name)

	fmt.Println("Hi", name)
	fmt.Println("How are you?")

	//RUN go run *.go Kate
}
