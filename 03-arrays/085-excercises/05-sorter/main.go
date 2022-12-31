package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------
const N = 5

func main() {
	args := os.Args[1:]
	var numbers [N]int

	if len(args) < 1 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	if len(args) > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	for i, a := range args {
		n, err := strconv.Atoi(a) // if it's not a number n = 0
		if err != nil {
			continue
		}

		numbers[i] = n
	}

	fmt.Printf("Given numbers: %#v\n", numbers)

	for i := 0; i < (N - 1); i++ {
		for j := 0; j < N-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				// swap
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Printf("Sorted dumbers: %#v\n", numbers)
}
