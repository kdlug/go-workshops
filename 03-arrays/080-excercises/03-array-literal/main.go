package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {

	names := [3]string{
		"Einstein",
		"Tesla",
		"Shepard",
	}

	distances := [5]int{
		50,
		40,
		75,
		30,
		125,
	}

	data := [5]uint8{
		72,
		69,
		76,
		76,
		79,
	}

	ratios := [1]float64{3.14}
	alives := [4]bool{
		true,
		false,
		true,
		false,
	}
	zero := [0]uint8{}

	fmt.Printf("names	 : %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data	 : %d\n", data)
	fmt.Printf("ratios	 : %.2f\n", ratios)
	fmt.Printf("alives	 : %t\n", alives)
	fmt.Printf("zero	 : %d\n", zero)

}
