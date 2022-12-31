package main

import "fmt"

func main() {
	data := []float64{10, 20, 30, 40}

	fmt.Printf("Data: %v\n", data)
	// Data: [10 20 30 40]

	newData := []float64{1, 5}

	// it changes only fist two elements
	// because newData has len=2
	// for i := range newData {
	// 	data[i] = newData[i]
	// }

	// this is equivalent
	copy(data, newData)
	fmt.Printf("Data: %v\n", data)
	// Data: [1 5 30 40]

	copy(data, []float64{100, 200, 250, 300, 400})
	fmt.Printf("Data: %v\n", data)
	// Data: [100 200 250 300]

	// Copy doesn't extend a slice, it uses the len of the smallest slice

	// if we want to copy the whole slice and extend it
	// we can append, data[:0] - set len to 0, because we want to start from 0 index
	data = append(data[:0], []float64{100, 200, 250, 300, 400}...)
	fmt.Printf("Data: %v\n", data)

	// or we can create a new slice and copy it
	saved := make([]float64, len(data))
	copy(saved, data)
	fmt.Printf("Saved: %v\n", saved)

	// or do the same with append which is shorter
	saved2 := append([]float64(nil), data...)
	fmt.Printf("Saved2: %v\n", saved2)

}
