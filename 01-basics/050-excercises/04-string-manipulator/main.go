package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("[command] [string]")
		fmt.Println()
		fmt.Println("Available commands: lower, upper and title")
		return
	}

	s := args[2]
	var t string
	switch command := args[1]; command {
	case "lower":
		// lower
		t = strings.ToLower(s)
	case "upper":
		// upper
		t = strings.ToUpper(s)
	case "title":
		t = strings.Title(s)
	default:
		fmt.Printf("Unknown command: %q", command)
		return
	}
	fmt.Println(t)
}
