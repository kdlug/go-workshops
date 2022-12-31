package main

import "fmt"

type list []item

type item interface {
	print()
	discount(float64)
}

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry, currently there are no books")
		return
	}

	for _, item := range l {
		item.print()
	}
}

func (l list) discount(ratio float64) {
	for _, item := range l {
		item.discount(ratio)
	}
}
