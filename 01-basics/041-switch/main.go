package main

import (
	"fmt"
	"os"
)

func main() {
	country := os.Args[1]

	switch country { // the switch type should be the same as case condiiton type
	case "Tokyo", "Kyoto": // if country == "Tokyo" || country == "Kyoto"
		fmt.Println("Tokyo")
		// break // break is not needed, go adds it by default, which is different than f.ex. in Java
	case "Warsaw":
		fmt.Println("Poland")
	case "Malaga":
		fmt.Println("Spain")
	default:
		fmt.Println("Provide the city")
	}
}
