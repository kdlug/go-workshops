// go run main1.go < file.txt
// go run main1.go beautiful
//today is a beautiful day
//ctrl+D
// curl -s https://www.rfc-editor.org/rfc/rfc8493.txt | go run main1.go  full
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	rx := regexp.MustCompile(`[^a-z]+`) // everything except letters
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")
		if len(word) > 2 {
			words[word] = true
		}
	}

	fmt.Printf("Map: %v\n", words)
	//fmt.Println("first:", words["first"], "last:", words["last"])
	for word := range words {
		fmt.Println(word)
	}
	query := args[0]
	//if _, ok := words[query]; ok { // existing keys already return true
	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}

	fmt.Printf("The input does not contain %q.\n", query)
}
