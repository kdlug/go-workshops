package main

import "fmt"

func main() {

	// inner array provides the type
	// students := [2][3]int{
	// 	[3]int{2, 9, 7},
	// 	[3]int{1, 6, 4},
	// }

	// we can ommit inner type
	// students := [2][3]int{
	// 	{2, 9, 7},
	// 	{1, 6, 4},
	// }

	// ellipsis
	students := [...][3]float64{
		{2, 9, 7},
		{1, 6, 4},
	}

	// i := "Bla"
	// const total = len(i) // ERROR len(i) (value of type int) is not constant

	// it works because array has fixed length, so len(students) returns a constant
	// const total int = len(students) * len(students[0])
	// we can omit type like in :=, but for const we can''t use :=
	const total = len(students) * len(students[0])

	//avg
	// sum_of_all_grades / total
	// sum := students[0][0] + students[0][1] + students[0][2] +
	// 	students[1][0] + students[1][1] + students[1][2]
	var sum float64
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}
	avg := sum / float64(total)

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Avg: %.2f\n", avg)
}
