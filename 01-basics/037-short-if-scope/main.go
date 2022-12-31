package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		fmt.Printf("%d * 2 = %d\n", n, 2*n) // n and a are visible because there are in the if block
	}

	// ERROR
	// fmt.Printf("%s * 2 %d\n", a[1], n) // a and n are undefined, because they are outside if {}
}
