package main

import "fmt"

func main() {
	var n []string
	friends := []string{"Julia", "Kate"}

	fmt.Printf("n	: %#v\n", n)
	fmt.Printf("friends	: %#v\n", friends)

	n = append(n, "Leon")
	fmt.Printf("n	: %#v\n", n)
	n = append(n, "Leon", "John", "Vincent") // we can append multiple elements
	fmt.Printf("n	: %#v\n", n)
	n = append(n, friends...) // or append a slice of the same type using variadic operator ...
	fmt.Printf("n	: %#v\n", n)
}
