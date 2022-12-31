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

	for i, v := range tasks {
		// if we want to add elements from the beggining, we have to use index
		upTasks[i] = strings.ToUpper(v)
	}

	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))
	// OUTPUT:
	// upTasks: [JUMP WRITE READ], len: 3, cap: 3
}
