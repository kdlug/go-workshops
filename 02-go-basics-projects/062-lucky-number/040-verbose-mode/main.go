package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
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
	var v bool // verbose
	rand.Seed(time.Now().UnixNano())

	// var v bool
	// flag.BoolVar(&v, "v", false, "verbose")
	// flag.Parse()

	args := os.Args[1:]
	if args[0] == "-v" {
		v = true
	}
	nArgs := 1

	if v {
		nArgs = 2
	}

	if len(args) != nArgs {
		fmt.Printf(usage, maxTurns)

		return
	}

	guess, err := strconv.Atoi(args[nArgs-1])

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
		if v {
			fmt.Printf("%d ", n)
		}
		if n == guess {
			fmt.Println("🎉 YOU WON!!!")
			return
		}
	}

	fmt.Println("😢 YOU LOST!!!")
}
