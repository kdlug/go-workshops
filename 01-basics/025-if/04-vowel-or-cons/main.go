package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

const vowels string = "aeiouąę"
const semiVowels string = "aeiou"

func main() {
	n := len(os.Args) - 1

	if n == 0 || utf8.RuneCountInString(os.Args[1]) != 1 {
		fmt.Printf("Give me a letter\n")

		return
	}

	l := os.Args[1]

	if strings.ContainsAny(l, vowels) {
		fmt.Printf("%q is a vowel.\n", l)
	} else if strings.ContainsAny(l, semiVowels) {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", l)
	} else {
		fmt.Printf("%q is a consonant.\n", l)
	}
}
