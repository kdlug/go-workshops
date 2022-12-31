package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := `Litwo, Ojczyzno moja! ty jesteś jak zdrowie;
Ile cię trzeba cenić, ten tylko się dowie,
Kto cię stracił. Dziś piękność twą w całej ozdobie
Widzę i opisuję, bo tęsknię po tobie.`

	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i]) // jesteÅak -> broken encoding
	}

	fmt.Print("\n", strings.Repeat("-", 45), "\n")

	// fix encoding
	for i := 0; i < len(text); {
		r, s := utf8.DecodeRuneInString(text[i:]) // returns the size of a rune
		fmt.Printf("%c", r)
		i += s
	}

	fmt.Print("\n", strings.Repeat("-", 45), "\n")

	// range loop jumps over the runes automatically
	for _, r := range text {
		fmt.Printf("%c", r)
	}

	fmt.Println()

}
