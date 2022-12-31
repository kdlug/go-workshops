package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("wheels on the bus go round and round all thorough the town")

	for i := 0; i < len(words); i++ {
		fmt.Printf("#%-2d: %s\n", i+1, words[i])
	}
}
