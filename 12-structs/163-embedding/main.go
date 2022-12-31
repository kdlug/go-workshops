package main

import "fmt"

type text struct {
	title string
	words int
}

type book struct {
	title string
	words int
	isbn  string
}

func main() {
	moby := book{
		title: "moby dick",
		words: 203412,
		isbn:  "123123",
	}
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)
}
