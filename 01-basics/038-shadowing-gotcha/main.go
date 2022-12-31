package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int // define n variable, default 0

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil { // n := - creates a nee local variable
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		fmt.Printf("%d * 2 = %d\n", n, 2*n) // n and a are visible because there are in the if block
	}

	// n has been shadowed, because
	fmt.Printf("n is %d - 😱😱😱 shadowing gotcha!\n", n)
}
