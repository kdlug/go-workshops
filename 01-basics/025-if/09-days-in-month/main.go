package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Give me a month name\n")
		return
	}

	m := strings.ToLower(os.Args[1])
	var d int // number of days
	y := time.Now().Year()
	leap := y%400 == 0 || y%100 != 0 && y%4 == 0

	if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		d = 31
	} else if m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		d = 31
	} else if m == "february" {
		if leap {
			d = 28
		} else {
			d = 29
		}
	} else {
		fmt.Printf("Give me a month name\n")
		return
	}

	fmt.Printf("%q has %d days\n", os.Args[1], d)

}
