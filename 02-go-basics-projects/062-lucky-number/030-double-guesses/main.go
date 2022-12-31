package main

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(os.Args) < 1 {
		fmt.Printf(usage, maxTurns)

		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number")

		return
	}

	var guess2 int
	if len(args) == 2 {
		guess2, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Not a number")

			return
		}
	}

	if guess < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number")

		return
	}

	// find the max number
	max := guess

	if guess2 > guess {
		max = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(max + 1)

		if n == guess || n == guess2 {
			fmt.Println("🎉 YOU WON!!!")

			return
		}
	}

	fmt.Println("😢 YOU LOST. TRY AGAIN?!!!")

}
