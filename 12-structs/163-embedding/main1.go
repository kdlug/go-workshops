package main

import "fmt"

type text struct {
	title string
	words int
}

type book struct {
	text text // embedded - composition
	isbn string
}

func main() {
	moby := book{
		text: text{title: "moby dick", words: 203412},
		isbn: "123123",
	}
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.text.title, moby.text.words, moby.isbn)
}
