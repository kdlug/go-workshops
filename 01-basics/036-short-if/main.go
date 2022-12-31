package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("13")

	if err == nil {
		fmt.Println("There is no error, n is", n)
	}

	// short if
	if n, err := strconv.Atoi("13"); err == nil { // n, err are shadowed here
		fmt.Println("There is no error, n is", n) // n, err are available nly within a block { }
	}
}
