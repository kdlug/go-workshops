package main

import (
	"fmt"
	"strings"
)

func main() {
	tasks := []string{"jump", "write", "read"}

	var upTasks []string

	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))
	// OUTPUT:
	//  [], len: 0, cap: 0
	for _, v := range tasks {
		// we can't use in range loop index like: upTasks[i] = v, only append
		// because it's empty now
		// upTasks[i] = v ->. panic: runtime error: index out of range [0] with length 0
		upTasks = append(upTasks, strings.ToUpper(v))
	}

	fmt.Printf("upTasks: %v, len: %d, cap: %d\n", upTasks, len(upTasks), cap(upTasks))
	// OUTPUT:
	// [JUMP WRITE READ], len: 3, cap: 4
}
