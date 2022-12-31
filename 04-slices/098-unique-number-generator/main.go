package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])
	var uniques []int

parent:
	for found := 0; found < max; {
		n := rand.Intn(max + 1)
		fmt.Print(n, " ")

		for _, v := range uniques {
			if v == n {
				continue parent
			}
		}
		uniques = append(uniques, n)
		found++
	}
	fmt.Println()
	fmt.Printf("uniques: %v\n", uniques)

	//sort
	sort.Ints(uniques)
	fmt.Printf("sorted: %v\n", uniques)

	// array
	a := [5]int{10, 2, 0, 11, 15}
	// sort.Ints(a) // this function takes a slice, so we need to convert array to a slice
	sort.Ints(a[:])
	// fmt.Printf("a: %#v\n", a[:]) // a: []int{0, 2, 10, 11, 15}
	// fmt.Printf("a: %#v\n", a) // a: [5]int{0, 2, 10, 11, 15}
	fmt.Printf("sorted: %v\n", a)
}
