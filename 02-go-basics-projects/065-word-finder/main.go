package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps and again and again"

func main() {
	words := strings.Fields(corpus)

	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%d - %s\n", i, w)
				break
			}
		}
	}
}
