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
		// fmt.Printf("%T\n", it)
		// extract a dynamic value from an interface value
		g, isGame := it.(*game)

		if isGame {
			g.discount(ratio)
		}

		// not very effective when we have check more than 1 type
		t, isToy := it.(*toy)
		if isToy {
			t.discount(ratio)
		}
	}
}
