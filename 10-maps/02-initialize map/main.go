package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide an english word to translate")
		return
	}

	word := args[0]
	// dict := map[string]string
	// dict := map[string]string{} // map initialization -> empty map literal {}
	dict := map[string]string{
		"cat":  "gato",
		"dog":  "perro",
		"duck": "pato",
	}

	dict["up"] = "arriba"
	dict["down"] = "abajo"
	dict["cat"] = "kot" // we can override a map value

	value, ok := dict[word]

	if !ok {
		fmt.Printf("The word %s doesn't exist in the dictionary\n", word)
		return
	}
	fmt.Printf("Key: %s, value: %s\n", word, value)

	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v) // the order sometimes changes
	}
	// for testing purposes fmt package print the map always sorted by key
	fmt.Printf("%#v\n", dict)

}
