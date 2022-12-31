package main

import "fmt"

func main() {
	f := 10.12300

	fmt.Printf("Number: %f\n", f)   // Number: 10.123000
	fmt.Printf("Number: %.0f\n", f) // Number: 10
	fmt.Printf("Number: %.1f\n", f) // Number: 10.1
	fmt.Printf("Number: %.2f\n", f) // Number: 10.12
	fmt.Printf("Number: %.3f\n", f) // Number: 10.123
}
