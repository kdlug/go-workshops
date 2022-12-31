package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Should take an argument, capitalize it and add some !
// f.ex hey -> !!! HEY !!! (3 characters -> 3x!)
func main() {
	s := os.Args[1]
	n := utf8.RuneCountInString(s)
	u := strings.ToUpper(s)
	ext := strings.Repeat("!", n)

	fmt.Println(ext, u, ext)
}
