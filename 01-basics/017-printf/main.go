// placeholders:
// %s - strings
// %q - quoted
// %d - decimals
// %t - bool
// %p - chan, pointer
// %g - float32, complex64
// %v - any value
// %T - type of a value
// /n -> escape sequence, a new line
package main

import "fmt"

func main() {
	var brand string

	fmt.Printf("%q\n", brand) // OUTPUT: "" -> quotes
	fmt.Printf("%s\n", brand) // OUTPUT:  -> without quotes

	brand = "Google"

	fmt.Printf("%q\n", brand) // OUTPUT: "Google" -> quotes
	fmt.Printf("%s\n", brand) // OUTPUT: Google -> without quotes

	n := 100

	fmt.Printf("%v\n", n) // OUTPUT: 100
	fmt.Printf("%d\n", n) // OUTPUT: 100
	fmt.Printf("%T\n", n) // OUTPUT: int
}
