package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Uppercase first letter

func main() {
	word := []byte("świat")
	fmt.Printf("%s = % [1]x\n", word)

	// we neeed the size of 1st rune
	var size int
	// converting slice byte to a string is not efficient
	// for i := range string(word) { // we have to convert to string, because we need runes, not bytes
	// 	if i > 0 { // i -> index of second rune is the size of 1st rune
	// 		size = i
	// 		break
	// 	}
	// }

	// more efficient way
	_, size = utf8.DecodeRune(word)
	copy(word[:size], strings.ToUpper(string(word[:size])))
	// fmt.Printf("%s%s\n", strings.ToUpper(string(word[:size])), string(word[size:]))
	fmt.Printf("%s = % [1]x\n", word)
}
