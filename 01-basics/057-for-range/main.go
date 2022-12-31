package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("wheels on the bus go round and round all thorough the town")

	// for i := 0; i < len(words); i++ {
	// 	fmt.Printf("#%-2d: %s\n", i+1, words[i])
	// }

	for i, v := range words {
		fmt.Printf("#%-2d: %s\n", i+1, v) // v = words[i]
	}

	// skipping i
	for _, v := range words {
		fmt.Printf("%s\n", v) // v = words[i]
	}

	// skipping v
	for i, _ := range words {
		fmt.Printf("#%-2d: %s\n", i+1, words[i])
	}

	// skipping v - simpler - omit v value
	for i := range words {
		fmt.Printf("#%-2d: %s\n", i+1, words[i])
	}
}
