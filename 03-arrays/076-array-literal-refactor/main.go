// refctor 073 to array literal
package main

import "fmt"

func main() {
	// var books [4]string
	books := [...]string{
		"Cars",
		"Frozen",
		"Lion King",
		"Cars Part 3",
	}

	fmt.Printf("books : %#v\n", books)
}
