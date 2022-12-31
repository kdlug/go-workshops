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
	"runtime"
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

	// var buf = make([]byte, 0, len(text)) //optimize, we know size of slice
	// or
	buf := []byte(text[:0]) // creates a new byte slice with text and rewind it
	// buf = append(buf, '!')  // this will append ! to the end and may create a new underlying array
	//buf = buf[:0]          // this rewinds buff, and will append from the beggining
	// fmt.Println(string(buf[:len(text)])) // we can get the slice back
	// we don't use runes because urls doens't containt special chars
	for i := 0; i < len(text); i++ {

		if len(text[i:]) >= nLink && text[i:i+nLink] == link {
			// fmt.Println(text[i:])
			isLink = true
			// append http:/
			buf = append(buf, link...)
			// skip http for masking
			i += nLink
		}
		c := text[i] // ptr // text[i] -> byte

		if c == '\n' || c == ' ' || c == '\t' {
			isLink = false
		}

		if isLink {
			c = mask
		}

		buf = append(buf, c)
	}
	fmt.Println(string(buf))
	Report()
}

func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
