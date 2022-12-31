package main

import "fmt"

type text struct {
	title string
	words int
}

// anonymous fields
type book struct {
	text // it doesn't have a name, anonymous field
	// go use a field name automatically using the type of anonymous field
	isbn string
}

func main() {
	moby := book{
		text: text{title: "moby dick", words: 203412}, // literal looks the same
		isbn: "123123",
	}

	// now you can access fields more easily
	moby.words++
	moby.text.words++ // this will work too
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)
}
