// /-> escape character
package main

import "fmt"

func main() {
	// /n is one byte, it's interpreted as a new line
	fmt.Printf("Len: %d", len("\n")) // OUTPUT: Len: 1

	// To escape use \
	fmt.Printf("\nEscape \\n, \"Hi\"\n") // OUTPUT: Escape \n

}
