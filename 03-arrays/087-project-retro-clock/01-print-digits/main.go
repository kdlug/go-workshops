package main

import "fmt"

// GOAL 1: Printing the Digits
// Challenge Steps
// 1. Define a new placeholder type

// 2. Create the digits from "zero" to "nine"

// You can use these characters for the clock:

// Digit character       : █
// Separator character   : ░
//
// 3. Put them into the "digits" array

// Print the digits side-by-side

// Loop for all the lines in a digit

// Print each digit line by line

// Don't forget printing a space after each digit

// Don't forget printing a newline after each line

// EXAMPLE: Let's say you want to print 10.

//       ██   ███ <--- Print a new line after printing a single line from
//        █   █ █     all the digits.
//        █   █ █
//        █   █ █
//       ███  ███
//          ^^
//          ||
//          ++----> Add space between the digits

func main() {
	const lines = 5
	type placeholder [lines]string

	zero := placeholder{
		"███", // 1 line
		"█ █", // 2 line
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [10]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	_ = separator

	for i := 0; i < lines; i++ { // 0 to 5
		for k := range digits {
			fmt.Print(digits[k][i], " ") // print 1st line of all digits
		}
		fmt.Println()

	}
	fmt.Println()

}
