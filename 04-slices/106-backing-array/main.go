package main

import (
	"fmt"
	"sort"
)

func main() {
	// grades := [...]float64{60, 40, 10, 25, 30, 36}
	grades := []float64{60, 40, 10, 25, 30, 36} // with the slice, the output will be the same

	fmt.Printf("grades: %#v\n", grades)
	front := grades[0:3]
	fmt.Printf("front: %#v\n", front)

	sort.Float64s(front)
	fmt.Println("after sorting:")

	// slicing doen't create a new underlying array
	// if we sliced an array - it will become an underlying array for the slice
	fmt.Printf("grades: %#v\n", grades) // [6]float64{10, 40, 60, 25, 30, 36}
	fmt.Printf("front: %#v\n", front)   // []float64{10, 40, 60}

	fmt.Println("===Create a new backing array===")
	// if we want to create a new underlying array
	var newGrades []float64

	newGrades = append(newGrades, grades...)
	front = newGrades[0:3]

	fmt.Printf("grades: %#v\n", grades)
	fmt.Printf("new grades: %#v\n", newGrades)
	fmt.Printf("front: %#v\n", front)
	front[0] = 999
	fmt.Println("Changing front[0] = 999")
	fmt.Printf("grades: %#v\n", grades)        // []float64{10, 40, 60, 25, 30, 36}
	fmt.Printf("new grades: %#v\n", newGrades) // []float64{999, 40, 60, 25, 30, 36}
	fmt.Printf("front: %#v\n", front)          // []float64{999, 40, 60}

}
