package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide your age")
		return
	}

	input := os.Args[1]
	age, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("SUCCESS! Converted %q to %d\n", input, age)
}
