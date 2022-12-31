package main

import "fmt"

func main() {
	var i int
	for {
		if i > 10 {
			break
		}

		if i%2 == 0 {
			i++
			continue
		}
		// prints only odd number
		fmt.Println(i)
		i++
	}
}
