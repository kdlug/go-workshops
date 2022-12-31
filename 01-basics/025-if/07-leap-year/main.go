package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// What is a leap year? To be a leap year, the year number must be divisible by four
// – except for end-of-century years, which must be divisible by 400.
// This means that the year 2000 was a leap year, although 1900 was not. 2020, 2024 and 2028 are all leap years.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Give me a year number\n")
		return
	}

	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%s is not a valid year.\n", os.Args[1])
		return
	}

	if year%400 == 0 || year%100 != 0 && year%4 == 0 {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}
}
