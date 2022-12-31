// go run main.go < file.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // set up to read words instead of lines

	for in.Scan() {
		fmt.Println(in.Text())
	}
}
