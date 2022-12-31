package main

import "fmt"

type list []printer

type printer interface {
	print()
}

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry, currently there are no books")
		return
	}

	for _, item := range l {
		item.print()

		// you cannot access to the discount method of the game type.
		// `it` is a printer not a game.
		// it.discount(.5)
	}
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		type discounter interface {
			discount(float64)
		}

		// fmt.Printf("%T\n", it)
		// check whether an interface value provide methods which you wants
		// item, ok := it.(interface{ discount(float64) })
		item, ok := it.(discounter)

		if ok {
			item.discount(ratio)
		}
	}
}
