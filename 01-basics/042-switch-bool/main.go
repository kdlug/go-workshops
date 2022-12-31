package main

import "fmt"

func main() {
	i := 10

	// switch true { // true is the default condition, so we can omit it
	switch {
	case i > 1:
		fmt.Println("positive")
	case i < 1:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
