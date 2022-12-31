package main

import (
	"fmt"
	"math/rand"
)

// for computers it's hard to generate a random number
// numbers generated by coputers are pseudo random

func main() {
	guess := 10

	// each time when you run the program output will be exactly the same
	// we can change it using seed
	rand.Seed(10) // rand.Intn depends on seed
	// if the seed number is the same - rand will generate the same output
	// so we have to make sure that seed is changing every time when program runs

	for n := 0; n != guess; {
		//
		n = rand.Intn(guess + 1) // rand generates pseuo random numbers
		// rand[0, n) - if we want to include n number, we have to add + 1
		fmt.Printf("%d ", n)
	}

	fmt.Println()
	// output without seed
	// 1 1 9 3 2 4 7 6 6 6 4 0 10
	// with seed
	// 6 8 1 0 2 4 6 0 10
}
