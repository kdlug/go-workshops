package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---")

	// we can move init statement outside the for loop
	var i int // init statement
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---")

	// we can remove ; from for
	i = 0 // init statement
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---")

	// we can move post statement from for
	i = 0 // init statement
	for i < 5 {
		fmt.Println(i)
		i++ // post statement
	}

	// we can remove condition
	// and use condition inside the loop with break
	// break breaks the loop
	i = 0 // init statement
	for { // for true - we can skip true, it's default infinite loop; like do->while
		if i == 5 { // should be break somewhere
			break
		}
		fmt.Println(i)
		i++ // post statement
	}
}
