package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	names := [...]string{
		"Einstein",
		"Tesla",
		"Shepard",
	}

	distances := [...]int{
		50,
		40,
		75,
		30,
		125,
	}

	data := [...]uint8{
		72,
		69,
		76,
		76,
		79,
	}

	ratios := [...]float64{3.14}
	alives := [...]bool{
		true,
		false,
		true,
		false,
	}
	zero := [...]uint8{}

	fmt.Printf("names	 : %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data	 : %d\n", data)
	fmt.Printf("ratios	 : %.2f\n", ratios)
	fmt.Printf("alives	 : %t\n", alives)
	fmt.Printf("zero	 : %d\n", zero)
}
