package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// GOAL 3: Animate the Clock
//
// Challenge Steps
//
// 1. Create an infinite loop to update the clock
// 2. Update the clock every second
// 	  time.Sleep(time.Second) will stop the world for 1 second
// 3. Clear the screen before the infinite loop
//    - Get my library for clearing the screen:
//      go get -u github.com/inancgumus/screen
//      github.com/inancgumus/screen@latest
//    - Then, import it and call it in your code like this:
//      screen.Clear()
//    - If you're using Go Playground instead, do this:
//      fmt.Println("\f")

// 4. Move the cursor to the top-left corner of the screen, before each step of the infinite loop
//    - Call this in your code like this:
//      screen.MoveTopLeft()
//    - If you're using Go Playground instead, do this again:
//      fmt.Println("\f")

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

	for i := 0; i < 10; i++ {

		t := time.Now()
		h, m, s := t.Hour(), t.Minute(), t.Second()

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

		time.Sleep(time.Second)
		screen.Clear()
	}

}
