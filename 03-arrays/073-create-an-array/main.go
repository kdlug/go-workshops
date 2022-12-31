package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [4]string
	var books [yearly]string // we can use constants

	fmt.Printf("books : %T\n", books)
	fmt.Println("books : ", books)
	fmt.Printf("books : %q\n", books)
	fmt.Printf("books : %#v\n", books)

	books[0] = "Cars"
	books[1] = "Frozen"
	books[2] = "Lion King"
	books[3] = books[0] + " Part 3"
	fmt.Printf("books : %#v\n", books)

	// CONST intex
	// books[4] = books[0] + " Part 2" // compile time error, index is constant value
	// output
	// ./main.go:26:8: invalid argument: array index 4 out of bounds [0:4]

	// VARIABLE index
	i := 4 // index is a variable, not const
	// it will be runtime error, because go has to calculate it when program runs
	books[i] = books[0] + " Part 2"
	// output
	// panic: runtime error: index out of range [4] with length 4
}
