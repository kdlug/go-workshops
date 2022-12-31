package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

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

	if len(os.Args) != 2 {
		fmt.Printf(usage, maxTurns)

		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number")

		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number")

		return
	}

	max := 10
	if guess > 10 {
		max = guess
	}
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(max + 1)

		if n == guess {
			fmt.Println("🎉 YOU WON!!!")
			return
		}
	}

	fmt.Println("😢 YOU LOST!!!")
}
