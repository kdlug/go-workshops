package main

import "fmt"

func main() {
	var s1 []string

	fmt.Printf("s1: %#v, len: %d\n", s1, len(s1)) // nil slice, because it's not initialized yet
	fmt.Printf("s1 == nil ? %v\n", s1 == nil)

	s2 := []string{}                              // initialized with empty value
	fmt.Printf("s2: %#v, len: %d\n", s2, len(s2)) // empty slice
	fmt.Printf("s2 == nil ? %v\n", s2 == nil)

	// fmt.Printf("s1 == s2 ? %v\n", s1 == s2) // ERROR: cannot compare, slices can be only compared to nil

	// if we want to check if slice is empty, we can do this
	if len(s1) == 0 {
		fmt.Println("Slice s1 is empty")
	}

	games := []string{"Doom 3", "Sims"}
	books := []string{"Lord of the Rings", "Witcher"}

	// Comparing slices
	// if games == books - ERROR not allowed
	// we can iterate
	var compare string
	for i, game := range games {
		if game != books[i] {
			compare = "not "
			break
		}
	}

	fmt.Printf("games and books slice are %sequal\n", compare)

	// assigning slices
	books = games
	fmt.Printf("books: %#v\n", books)
	fmt.Printf("games: %#v\n", games)

	games = nil
	fmt.Printf("books: %#v\n", books)
	fmt.Printf("games: %#v\n", games)

	// compare again
	compare = ""
	for i, game := range games { // games is nil, so range won't break the program, but this loop won't run
		if game != books[i] {
			compare = "not "
			break
		}
	}
	// so we have to compare len
	if len(games) != len(books) {
		compare = "not "
	}

	fmt.Printf("games and books slice are %sequal\n", compare)

	// assigning
}
