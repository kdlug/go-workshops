package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string

	books[0] = "Cars"
	books[1] = "Frozen"
	books[2] = "Lion King"
	books[3] = books[0] + " Part 3"

	fmt.Printf("books : %#v\n", books)
	var (
		wBooks [winter]string // winter books
		sBooks [summer]string // summer
	)

	// 1 method
	wBooks[0] = books[0]
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	// 2 method
	// for i := 0; i < len(sBooks); i++ {
	// 	sBooks[i] = books[i+1]
	// }

	// 3 method
	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	// we want to add additional string to each book title
	for _, v := range sBooks { // because v is a copy of a value - the original array doesn't change
		v += "won't effect"
	}

	fmt.Printf("winter : %#v\n", wBooks)
	fmt.Printf("summer : %#v\n", sBooks)
}
