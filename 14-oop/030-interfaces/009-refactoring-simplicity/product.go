package main

import "fmt"

type product struct {
	title    string
	price    float64
	released timestamp
}

func (p *product) print() {
	fmt.Printf("%-15s: %.2f, released: %s\n", p.title, p.price, p.released.string())
}

func (p *product) discount(ratio float64) {
	p.price *= float64(1 - ratio)
}
