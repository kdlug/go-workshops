package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	s := strings.Split(data, " ")
	var numbers []int

	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		numbers = append(numbers, n)
	}

	fmt.Printf("numbers		: %#v\n", numbers)

	evens, odds := numbers[:3], numbers[3:]

	fmt.Printf("even		: %#v\n", evens)
	fmt.Printf("odds		: %#v\n", odds)

	mid := numbers[2:4]

	fmt.Printf("mid		: %#v\n", mid)

	first2 := numbers[0:2]

	fmt.Printf("first2		: %#v\n", first2)

	last2 := numbers[len(numbers)-2:]

	fmt.Printf("last2		: %#v\n", last2)

	elast := evens[len(evens)-1:]

	fmt.Printf("evens last	: %#v\n", elast)

	olast2 := odds[len(odds)-2:]

	fmt.Printf("odds last two	: %#v\n", olast2)

}
