package main

import "fmt"

// iota - built in number generator
func main() {
	// const (
	// 	MONDAY    = 0
	// 	TUESDAY   = 1
	// 	WENDESDAY = 2
	// 	THURSDAY  = 3
	// 	FRIDAY    = 4
	// 	SATURDAY  = 5
	// 	SUNDAY    = 6
	// )

	const (
		MONDAY  = iota // 0
		TUESDAY        // 1
		WENDESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY
	)

	fmt.Println(MONDAY, TUESDAY, WENDESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY)
	// OUTPUT:
	// 0 1 2 3 4 5 6

	// const (
	// 	EST = -5
	// 	MST = -7
	// 	PST = -8
	// )

	const (
		EST = -(iota + 5) // -5
		_                 // blank idendifier means that iota is incremented but the value goes to the blackhole
		MST               // -7
		PST               // -8
	)

	fmt.Println(EST, MST, PST)

	// OUTPUT
	// -5 -7 -8
}
