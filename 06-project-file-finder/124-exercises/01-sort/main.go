package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide some items and I will sort them")
		return
	}

	sort.Strings(args)

	var items []byte

	for _, item := range args {
		items = append(items, item...)
		items = append(items, '\n')
	}

	err := ioutil.WriteFile("out.txt", items, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
