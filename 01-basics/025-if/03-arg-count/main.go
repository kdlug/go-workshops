package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	n := len(os.Args) - 1

	if n == 0 {
		fmt.Printf("Give me args\n")
	} else if n == 1 {
		fmt.Printf("There is one: \"%s\"\n", os.Args[1])
	} else if n == 2 {
		fmt.Printf("There are two: \"%s %s\"\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("There are %d arguments\n", n)
	}
}
