package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

const corpus = "lazy cat jumps and again and again"

func main() {
	words := strings.Fields(corpus)

	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%d - %s\n", i, w)
				break
			}
		}
	}

}
