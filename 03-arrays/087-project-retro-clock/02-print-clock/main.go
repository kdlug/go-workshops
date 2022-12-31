package main

import (
	"fmt"
	"time"
)

// GOAL 2: Printing the Clock
//
// Challenge Steps
//
// 1. Get the current time
// 2. Get the current hour, minute and second from the current time
// 3. Create the clock array by getting the digits from the digits array
// 4. Print the clock by using the clock array
// 5. Create a separator array (it's also a placeholder type)
// 6. Add the separators into the correct positions of the clock array

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

	// get time
	t := time.Now()
	h, m, s := t.Hour(), t.Minute(), t.Second()

	// extract digits
	// firstH, _ := strconv.Atoi(strconv.Itoa(h)[0:1])
	// secondH, _ := strconv.Atoi(strconv.Itoa(h)[1:2])

	// firstM, _ := strconv.Atoi(strconv.Itoa(m)[0:1])
	// secondM, _ := strconv.Atoi(strconv.Itoa(m)[1:2])

	// firstS, _ := strconv.Atoi(strconv.Itoa(s)[0:1])
	// secondS, _ := strconv.Atoi(strconv.Itoa(s)[1:2])

	// // put digits into the clock array
	// clock := [8]placeholder{
	// 	digits[firstH],
	// 	digits[secondH],
	// 	separator,
	// 	digits[firstM],
	// 	digits[secondM],
	// 	separator,
	// 	digits[firstS],
	// 	digits[secondS],
	// }

	// better way
	// put digits into the clock array
	clock := [8]placeholder{
		digits[h/10], digits[h%10],
		separator,
		digits[m/10], digits[m%10],
		separator,
		digits[s/10], digits[s%10],
	}
	// print the clock
	fmt.Println()
	for i := 0; i < lines; i++ { // 0 to 5
		for k := range clock {
			fmt.Print(clock[k][i], " ") // print 1st line of all digits
		}
		fmt.Println()
	}

	fmt.Println()
}
