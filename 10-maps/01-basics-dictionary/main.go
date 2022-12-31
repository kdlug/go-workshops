package main

import "fmt"

func main() {
	var dict1 map[string]string // map definition

	fmt.Printf("# of keys: %d\n", len(dict1))
	fmt.Printf("Zero value: %#v\n", dict1) // you can read from an uninitialized map

	// panic: assignment to entry in nil map
	// you cannot write to uninitialized map
	// dict1["up"] = "arriba"
	// dict1["down"] = "abajo"
}
