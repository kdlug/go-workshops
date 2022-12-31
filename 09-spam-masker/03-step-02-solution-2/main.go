/*
✅ #1- Get and check the input
✅ #2- Create a byte buffer and use it as the output
✅ #3- Write input to the buffer as it is and print it
✅ #4- Detect the link
✅ #5- Mask the link
✅ #6- Stop masking when whitespace is detected
✅ #7- Put a http:// prefix in front of the masked link
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		link  = "http://"
		nLink = len(link)
		mask  = '*'
	)
	var isLink bool

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Provide some text to mask urls")
		return
	}
	text := args[0]

	buf := []byte(text) // creates a new byte slice with text

	for i := 0; i < len(text); i++ {

		if len(text[i:]) >= nLink && text[i:i+nLink] == link {
			// fmt.Println(text[i:])
			isLink = true

			// skip http for masking
			i += nLink
		}
		b := buf[i] // buf[i] -> byte (copy)

		if b == '\n' || b == ' ' || b == '\t' {
			isLink = false
		}

		if isLink {
			buf[i] = mask // we can't do b = mask, because b is a copy
		}

	}
	fmt.Println(string(buf))
}
