package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g *game) print() {
	fmt.Printf("Title: %s, price: %.2f\n", g.title, g.price)
}
