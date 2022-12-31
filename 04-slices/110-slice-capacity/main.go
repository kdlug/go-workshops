package main

import "fmt"

func main() {
	var books []string // nil slice ptr: 0x0, cap: 0, len: 0
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", books, books, cap(books), len(books))

	// empty slices has the same dummy backing array
	books = []string{} // empty slice; ptr: 0x100fe6fc8, cap: 0, len: 0
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", books, books, cap(books), len(books))

	movies := []string{} // empty slice; ptr: 0x100fe6fc8, cap: 0, len: 0
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", movies, movies, cap(movies), len(movies))

	// the backing array pointer changes when we add some elements
	books = []string{"Mort", "The Fifth Elephant", "Moving Pictures", "Soul Music"}

	// look at pointers
	for cap(books) != 0 {
		books = books[1:cap(books)]
		fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", books, books, cap(books), len(books))
	}

}
