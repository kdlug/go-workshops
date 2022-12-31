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

	// dict := map[string]string{} // map initialization -> empty map literal {}
	dict := map[string]string{
		"cat":       "gato",
		"dog":       "perro",
		"duck":      "pato",
		"pretty":    "bonita",
		"beautiful": "bonita",
	}

	delete(dict, "beautiful")  // delete a key
	delete(dict, "beautiful")  // we can call it multiple times
	delete(dict, "noexisting") // we doesn't have to check if key exists

	// copy a map
	// dict2 := dict // copies only a reference to dict
	dict2 := make(map[string]string)
	dict2["test"] = "fsd"
	for k, v := range dict {
		dict2[v] = k
	}

	value, ok := dict[word]

	if !ok {
		fmt.Printf("The word %s doesn't exist in the dictionary\n", word)
		return
	}
	fmt.Printf("Key: %s, value: %s\n", word, value)

	for k, v := range dict2 {
		fmt.Printf("%q means %#v\n", k, v) // the order sometimes changes
	}

	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v) // the order sometimes changes
	}
	// for testing purposes fmt package print the map always sorted by key
	fmt.Printf("%#v\n", dict)
	fmt.Printf("%#v\n", dict2)

	// Clear / destroy a map
	// dict2 = nil // actually it doesnt clear the memory, values of a map are still there
	// because map store a pointer, so this removes only a reference

	// we have to iterate and clear all keys
	for k := range dict2 {
		delete(dict2, k)
	}

	fmt.Printf("%#v\n", dict2)

}
