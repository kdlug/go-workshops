package main

import "fmt"

var N int // package scope variables are visible within functions
// it's not a good practice

func main() {
	n, m := 2, 3 // local variables are invisible in funcions
	res := multiply(n, m)
	fmt.Printf("%d * %d = %d\n", n, m, res)

	//
	incN()
	incN()
	showN()
	incN()
	showN()
}

func multiply(a, b int) int {
	// a,b // local variables
	a *= b
	return a
}

func incN() {
	N++
}

func showN() {
	if N == 0 {
		return // breaks the function
	}

	fmt.Println(N)
}
