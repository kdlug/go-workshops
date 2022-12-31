package main

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉 YOU WON!!!")
			case 1:
				fmt.Println("😎 YOU'RE AWESOME!!!")
			case 2:
				fmt.Println("😍 LOVELY!!!")
			}

			if turn == 0 {
				fmt.Println("In the 1st turn!!!")
			}

			return
		}
	}

	switch rand.Intn(2) {
	case 0:
		fmt.Println("😢 YOU LOST!!!")
	case 1:
		fmt.Println("😢 YOU LOST. TRY AGAIN?!!!")
	}
}
