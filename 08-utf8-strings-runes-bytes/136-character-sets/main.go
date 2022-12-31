package main

import (
	"fmt"
	"strings"
)

// 1 bit: 0-127
func main() {
	//start, stop := 'A', 'Z'
	start, stop := 255, 500
	fmt.Printf("%-10s %-10s %-10s % -12s \n%s\n", "literal", "dec", "hex", "encoded", strings.Repeat("-", 45))

	for i := start; i < stop; i++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", i, string(i))
	}
}
