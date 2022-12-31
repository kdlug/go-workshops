package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func main() {
	var (
		min int = 1
		max int = 10
		res int
	)

	for i := min; i <= max; i++ {
		res += i
	}

	fmt.Printf("Sum: %d\n", res)
}
