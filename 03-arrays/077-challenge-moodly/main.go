package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[your name]")
		return
	}

	rand.Seed(time.Now().UnixNano())

	moods := [...]string{
		"good 👍",
		"happy 😀",
		"awesome 😎",
		"bad 👎",
		"sad 😞",
		"terrible 😩",
	}

	name := os.Args[1]
	r := rand.Intn(len(moods))

	fmt.Printf("%s feels %s", name, moods[r])
	fmt.Println()
}
