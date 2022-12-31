package main

import (
	"fmt"
	"strconv"
)

type product struct {
	Title    string
	Price    money
	Released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("%-15s: %.2f, released: %s", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}

func (ts timestamp) MarshalJSON() (data []byte, _ error) { // _ skip variable name for error
	// ts -> integer ts.Unix()->integer
	// data <- integer strconv.AppendInt(data, integer)
	return strconv.AppendInt(data, ts.Unix(), 10), nil

}

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))

	// ts -> integer ts.Unix()->integer
	// data <- integer strconv.AppendInt(data, integer)
	return nil

}
