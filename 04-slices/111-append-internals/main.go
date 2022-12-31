package main

import "fmt"

func main() {

	var numbers []int // creates a  nil slice
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", numbers, numbers, cap(numbers), len(numbers))
	// []int(nil), ptr: 0x0, cap: 0, len: 0

	// extending a slice
	numbers = append(numbers, 20, 30) // append creates a new backing array under a new memory location with bigger capacity
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", numbers, numbers, cap(numbers), len(numbers))
	// []int{20, 30}, ptr: 0x1400012c020, cap: 2, len: 2

	numbers = append(numbers, 40) // still the same location
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", numbers, numbers, cap(numbers), len(numbers))
	// []int{20, 30, 30}, ptr: 0x14000132020, cap: 4, len: 3

	numbers = append(numbers, 1)

	// now we want to add 2 numbers in the middle
	numbers = append(numbers, numbers[2:]...) // copy 2 last numbers to the end of the slice
	numbers = append(numbers[0:2], 5, 6)      // append new numbers in the middle
	numbers = numbers[:6]                     // extend a slice
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", numbers, numbers, cap(numbers), len(numbers))

}
