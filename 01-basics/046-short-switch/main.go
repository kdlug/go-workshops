package main

import "fmt"

func main() {

	// use ; to separate statement from switch condition
	// 	switch i := 10; true{
	switch i := 10; { // true is the default condition, so we can omit it, but ; should stay
	case i > 1:
		fmt.Println("positive")
	case i < 1:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
