package main

import (
	"fmt"
	"strings"
)

func main() {
	tasks := []string{"jump", "write", "read"}

	// now we can't use upTasks[i] = v, we can append, because len=0
	upTasks := make([]string, 0, len(tasks)) // we specified len=0, if we omit the last arg, cap=len
	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))

	// OUTPUT:
	// [], len: 0, cap: 3 // 0 elements
	for _, v := range tasks {
		upTasks = append(upTasks, strings.ToUpper(v))
	}

	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))
	// OUTPUT:
	// upTasks: [JUMP WRITE READ], len: 3, cap: 3
}
