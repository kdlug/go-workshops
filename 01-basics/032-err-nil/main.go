package main

import (
	"fmt"
	"strconv"
)

func main() {
	// valid
	input := "3"
	n, err := strconv.Atoi(input)
	fmt.Printf("Input: %q, Returned value: %d, Error: %v\n", input, n, err)

	//invalid
	input = "three"
	n, err = strconv.Atoi(input)
	// it returns 0 even if there is an error
	fmt.Printf("Input: %q, Returned value: %d, Error: %v\n", input, n, err)
}
