package main

import (
	"fmt"
	"time"
)

func main() {
	switch hour := time.Now().Hour(); {
	case hour >= 18:
		fmt.Println("Good evening!")
	case hour >= 12:
		fmt.Println("Good afternoon!")
	case hour >= 6:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good night!")
	}

}
