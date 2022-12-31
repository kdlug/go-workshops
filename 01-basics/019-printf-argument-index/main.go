package main

import "fmt"

func main() {
	a, b := "Temperature", 29
	fmt.Printf("Today weather: %v, %v,  %[2]v is the %[1]v",
		a, b)
}
