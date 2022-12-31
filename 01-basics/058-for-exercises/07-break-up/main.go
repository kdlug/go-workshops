package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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
	i := min

	for {
		if i > max {
			break
		}
		if i%2 != 0 {
			i++
			continue
		}

		s += strconv.Itoa(i)

		sum += i

		if i < max-1 {
			s += " + "
		}

		i++
	}
	s += " = " + strconv.Itoa(sum)
	fmt.Println(s)
}
