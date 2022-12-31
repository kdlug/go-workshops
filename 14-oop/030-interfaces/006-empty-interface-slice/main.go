package main

import "fmt"

func main() {
	var any interface{}

	// we can assign a slice to an empty interface{} value
	nums := []int{1, 2, 3, 4, 5}
	any = nums

	// but we can't use len
	// len(any) // invalid argument: any (variable of type interface{}) for len
	// we have to do assertion
	_ = len(any.([]int))

	fmt.Printf("%#v\n", any)

	// slice of empty interfaces
	var many []interface{}
	// many = nums different types []int != []interface{}
	// we can use a loop

	for _, n := range nums {
		many = append(many, n)
	}

	fmt.Printf("%#v\n", many)

}
