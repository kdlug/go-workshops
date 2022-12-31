// ctrl+D to stop the program
// go run main.go < file.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var lines int
	for in.Scan() {
		lines++
		fmt.Println("Scanned Text: ", in.Text())
		// in.Text() print only the last line
	}

	fmt.Printf("There are %d line(s)\n", lines)

	// we can check errors only once at the end, because bufio has Err function
	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
