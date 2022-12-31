/*
✅ #1- Get and check the input
✅ #2- Create a byte buffer and use it as the output
✅ #3- Write input to the buffer as it is and print it
✅ #4- Detect the link
#5- Mask the link
#6- Stop masking when whitespace is detected
#7- Put a http:// prefix in front of the masked link
*/

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var (
		link  = "http://"
		nLink = len(link)
	)
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Provide some text to mask urls")
		return
	}
	text := args[0]
	//var buf []byte - buff will create new underl arrays
	buf := make([]byte, 0, len(text)) //optimize, we know size of slice
	// for _, r := range text { // iterates over runes (int32, not bytes)
	// 	buf = append(buf, r)
	// }

	// we don't use runes because urls doens't containt special chars
	for i := 0; i < len(text); i++ {
		buf = append(buf, text[i]) // text[i] -> byte
		// fmt.Printf("len: %d, cap: %d, ptr: %p ", len(buf), cap(buf), buf)
		// fmt.Println(string(buf))

		if len(text[i:]) >= nLink && text[i:i+nLink] == link {
			fmt.Println(text[i:])
		}
	}
	//fmt.Printf("%v", buf)
	// check buffer -> if it allocates a new backing array
	// if yes -> create big enough buffer
	Report()
}

func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
