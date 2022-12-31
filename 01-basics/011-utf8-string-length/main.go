package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Dzień dobry"

	fmt.Println("Bytes:", len(s))
	fmt.Println("Chars:", utf8.RuneCountInString(s))
}
