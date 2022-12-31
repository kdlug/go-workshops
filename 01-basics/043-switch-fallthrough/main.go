package main

import "fmt"

func main() {
	i := 110

	switch {
	case i > 100: // the fist condition is checked and has to be fullfiled
		fmt.Printf("big ")
		fallthrough
	case i < 10: // next confitions are ignored, but blocks inside are executed
		fmt.Printf("positive ")
		fallthrough
	default:
		fmt.Printf("number")
	}

	// OUTPUT:
	// big positive number
}
