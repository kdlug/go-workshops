package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceable := []byte{'f', 'u', 'l', 'l'} // len 4, capacity 4
	fmt.Printf("sliceable before: ptr -> %p, cap: %d, len: %d, %s\n", sliceable, cap(sliceable), len(sliceable), sliceable)

	newSlice := sliceable[0:3] // len 3, capacity 4 → capacity of a backing array
	fmt.Printf("newSlice before:  ptr -> %p, cap: %d, len: %d, %s\n", newSlice, cap(newSlice), len(newSlice), newSlice)

	newSlice = append(newSlice, '!')
	// sliceable -> ful! - we lost the original element of backing array
	// newSlice -> ful!
	fmt.Printf("newSlice:    	  ptr -> %p, cap: %d, len: %d, %s\n", newSlice, cap(newSlice), len(newSlice), newSlice)
	fmt.Printf("sliceable after:  ptr -> %p, cap: %d, len: %d, %s\n", sliceable, cap(sliceable), len(sliceable), sliceable)
	fmt.Println(strings.Repeat("=", 75))

	sliceable = []byte{'f', 'u', 'l', 'l'}
	newSlice = sliceable[0:3:3] //  len 3, capacity 3
	fmt.Printf("sliceable before: ptr -> %p, cap: %d, len: %d, %s\n", sliceable, cap(sliceable), len(sliceable), sliceable)
	fmt.Printf("newSlice before:  ptr -> %p, cap: %d, len: %d, %s\n", newSlice, cap(newSlice), len(newSlice), newSlice)

	newSlice = append(newSlice, '!') // len 3
	// it's not enouth capacity to add a new element, so go will create a new backing array
	// which prevents to override it
	// sliceable -> full
	// newSlice -> ful!

	fmt.Printf("newSlice:    	  ptr -> %p, cap: %d, len: %d, %s\n", newSlice, cap(newSlice), len(newSlice), newSlice)
	fmt.Printf("sliceable after:  ptr -> %p, cap: %d, len: %d, %s\n", sliceable, cap(sliceable), len(sliceable), sliceable)

}
