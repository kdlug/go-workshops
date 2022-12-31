package main

import "fmt"

func main() {
	var any interface{}

	any = []int{1, 2, 3}

	any = map[int]bool{1: true, 2: false}
	any = "string"
	any = 10
	// we can't multiply because any is not a number, it's empty interface value
	// any = any * 2

	// we can use assertion to extract the number
	any = any.(int) * 2
	fmt.Println(any)
}
