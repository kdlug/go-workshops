package api

var temps = []int{5, 10, 3, 25, 45, 80, 990}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	portion := make([]int, 0, cap(temps))
	portion = append(portion, temps[start:stop]...)
	portion = portion[0:len(portion)]
	//portion := temps[start:stop]

	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// ----------------------------------------

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
