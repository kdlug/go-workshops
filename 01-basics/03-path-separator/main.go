package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("css/body.css")

	// or shortcut, without declaring variables
	// dir, file := path.Split("css/body.css")

	// if we don't need dir, use underscore _ variable
	// _, file = path.Split("css/body.css")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
