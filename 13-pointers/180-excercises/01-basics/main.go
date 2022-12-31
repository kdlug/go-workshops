package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var c *computer

	// compare the pointer variable to a nil value, and say it's nil
	if c == nil {
		fmt.Println("c is nil", c)
	}

	// create an apple computer by putting its address to a pointer variable
	apple := &computer{"Apple"}

	// put the apple into a new pointer variable
	newApple := apple
	// compare the apples: if they are equal say so and print their addresses

	if newApple == apple {
		fmt.Println("apple == newApple", &apple)
	}

	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{"Sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses

	if sony != apple {
		fmt.Println("sony != apple", &apple, &sony)
	}
	// put apple's value into a new ordinary variable
	newVar := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("&apple %p, apple %p, &newApple %p\n", &apple, apple, &newVar)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple == newVar {
		fmt.Println("*apple == newVar")

		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Printf("*apple val=%v\n", *apple)
		fmt.Printf("newVar val=%v\n", newVar)
	}

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Println("Sony changed to hp:", *sony)

	// and call the constructor func 3 times and print the returned values' addresses
	fmt.Printf("Constructor 1x %p\n", newComputer("apple"))
	fmt.Printf("Constructor 2x %p\n", newComputer("hp"))
	fmt.Printf("Constructor 3x %p\n", newComputer("asus"))
}

// create a new function: change
// the func can change the given computer's brand to another brand
func change(c *computer, b string) {
	// c is a copy of a value passed to the function
	// but this value is a pointer and contains the address of variable
	// exactly the same, so pointer points the same address
	c.brand = b
}

// write a func that returns the value that is pointed by the given *computer
// print the returned value
func valueOf(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer

func newComputer(b string) *computer {
	return &computer{brand: b}
}
