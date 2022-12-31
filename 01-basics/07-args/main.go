// go run *.go one two three - it will show a temporary path so use build
// go build -o greeter main.go
// ./greeter one two three
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st parameter: ", os.Args[1])
	fmt.Println("2nd parameter: ", os.Args[2])
	fmt.Println("3rd parameter: ", os.Args[3])
	fmt.Println("Number of arguments: ", len(os.Args))

	/*
		[]string{"./greeter", "one", "two", "three"}
		Path:  ./greeter
		1st parameter:  one
		2nd parameter:  two
		3rd parameter:  three
		Number of arguments:  4
	*/
}
