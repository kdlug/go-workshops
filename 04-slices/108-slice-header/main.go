package main

import "fmt"

// For array
// type collection [4]string
//
// OUTPUT:
// change's data: main.collection{"slices", "are", "brilliant", "period"}
// main's data: main.collection{"slices", "are", "awesome", "period"}
// because arguments are copied in functions

type collection []string

// OUTPUT:
// change's data: main.collection{"slices", "are", "brilliant", "period"}
// main's data: main.collection{"slices", "are", "brilliant", "period"}
// because slice is copied in functions
// so it describes the same underlying array

func main() {
	data := collection{"slices", "are", "awesome", "period"}

	change(data)

	fmt.Printf("main's data: %#v\n", data)

}

func change(data collection) {
	data[2] = "brilliant"
	fmt.Printf("change's data: %#v\n", data)
}
