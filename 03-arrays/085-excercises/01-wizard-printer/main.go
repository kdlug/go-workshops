package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {

	scientists := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Marie", "Hawking", "blackhole"},
		{"Stephen", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}
	// headers
	fmt.Printf("%-15s %-15s %-15s\n", "First Name", "Last Name", "Nickname")
	fmt.Printf(strings.Repeat("=", 45))
	fmt.Println()

	for i := range scientists {
		s := scientists[i]
		fmt.Printf("%-15s %-15s %-15s\n", s[0], s[1], s[2])
	}
}
