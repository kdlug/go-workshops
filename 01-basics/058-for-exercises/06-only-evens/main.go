package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------
var (
	min, max int
	sum      int
	err      error
	usage    = "[min_number] [max_number]"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	if min, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(usage)
		return
	}

	if max, err = strconv.Atoi(os.Args[2]); err != nil {
		fmt.Println(usage)
		return
	}

	if max <= min {
		fmt.Println("max should be higher than min")
		fmt.Println(usage)
		return
	}

	var s string
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}

		s += strconv.Itoa(i)

		sum += i

		if i < max-1 {
			s += " + "
		}
	}
	s += " = " + strconv.Itoa(sum)
	fmt.Println(s)
}
