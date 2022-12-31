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
	}
}
