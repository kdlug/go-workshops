package main

import (
	"fmt"
	"strings"
)

func main() {
	tasks := []string{"jump", "write", "read"}

	// now we can use upTasks[i] = v
	upTasks := make([]string, len(tasks)) // if we omit the last arg, cap=len

	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))
	// OUTPUT:
	// [  ], len: 3, cap: 3 // 3 elements with "" value

	for _, v := range tasks {
		// appends add an element to the end of the slice, after it's len
		// len is 3, so it will add as 4th element
		upTasks = append(upTasks, strings.ToUpper(v))
	}

	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))
	// OUTPUT:
	// upTasks: [jump write read JUMP WRITE READ], len: 6, cap: 6
}
