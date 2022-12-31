package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Append and Sort Numbers
//
//  We'll have a []string that should contain numbers.
//
//  Your task is to convert the []string to an int slice.
//
//  1. Get the numbers from the command-line
//
//  2. Append them to an []int
//
//  3. Sort the numbers
//
//  4. Print them
//
//  5. Handle the error cases
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    provide a few numbers
//
//  go run main.go 4 6 1 3 0 9 2
//    [0 1 2 3 4 6 9]
//
//  go run main.go a b c
//    []
//
//  go run main.go 4 a 1 c d 9
//    [1 4 9]
//
// ---------------------------------------------------------

func main() {
	var numbers []int

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide a few numbers")
		return
	}

	for _, v := range args {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		numbers = append(numbers, n)

	}
	sort.Ints(numbers)
	fmt.Printf("numbers: %v\n", numbers)

}
