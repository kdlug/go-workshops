package main

import "fmt"

type product struct {
	title string
	price float64
}

func (p *product) print() {
	fmt.Printf("%-15s: %.2f\n", p.title, p.price)
}

func (p *product) discount(ratio float64) {
	p.price *= float64(1 - ratio)
}
