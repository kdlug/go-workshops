package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, w := range words {
		fmt.Printf("%q", w)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\tbyte len: %d, rune len: %d", len(w), utf8.RuneCountInString(w))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\thex bytes: % x", w)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\thex runes: ")
		for _, r := range w {
			fmt.Printf("% x", r)
		}

		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trunes: ")
		for _, r := range w {
			fmt.Printf("%q ", r)
		}

		fmt.Println()
		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		fr, size := utf8.DecodeRuneInString(w)
		fmt.Printf("\tfirst rune: %c, byte size: %d", fr, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		lr, size := utf8.DecodeLastRuneInString(w)
		fmt.Printf("\tlast rune: %c, byte size: %d", lr, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(w)
		_, second := utf8.DecodeRuneInString(w[first:])
		fmt.Printf("\tfirst 2 : %q\n", w[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(w)
		_, last2 := utf8.DecodeLastRuneInString(w[:len(w)-last1])
		fmt.Printf("\tlast 2  : %q\n", w[len(w)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		runes := []rune(w)
		fmt.Printf("\tfist 2  : %q\n", runes[0:2])
		fmt.Printf("\tlast 2  : %q\n", runes[len(runes)-2:])
		fmt.Println()
	}
}
