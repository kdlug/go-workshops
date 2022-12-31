package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide a directory name")
		return
	}

	files, err := ioutil.ReadDir(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	var filenames []byte
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name() // name is converted to the byte slice on the fly
			filenames = append(filenames, name...)
			filenames = append(filenames, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", filenames, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
