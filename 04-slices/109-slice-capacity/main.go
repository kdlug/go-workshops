package main

import "fmt"

func main() {
	ages := []int{10, 20, 30}

	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", ages, ages, cap(ages), len(ages))

	//capacity determines how much you can extend a slice
	ages = ages[0:0] //  a new slice; capacity is still 3
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", ages, ages, cap(ages), len(ages))

	ages = ages[0:3] // so we can reslice it again and get access to all it's elelements
	// ages = ages[0:cap(ages)]
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", ages, ages, cap(ages), len(ages))

	ages = ages[1:3] // now capacity is 2, and the pointer has changed, we lost info about element 0
	// a slice cannot see the 1st element of it's backing array anymore
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", ages, ages, cap(ages), len(ages))

	// ages = ages[0:3]
	// panic: runtime error: slice bounds out of range [:3] with capacity 2
	// ages = ages[1:3] // we can't do it because cap is 2 now
	ages = ages[0:cap(ages)] // slice: []int{20, 30}, pointer: 0x14000114018, cap: 2, len: 2
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", ages, ages, cap(ages), len(ages))

	ages = ages[2:2] // cap: 0, len: 0 and pointer still points to the backing array, but we can't access elements
	fmt.Printf("slice: %#v, ptr: %p, cap: %d, len: %d\n", ages, ages, cap(ages), len(ages))

}
