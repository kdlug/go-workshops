package main

import (
	"fmt"
	"os"
	"strconv"
)

//const max int = 5

func main() {

	args := os.Args

	if len(os.Args) != 2 {
		fmt.Println("Give a max number")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%s is not a correct number\n", args[1])
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}
}
