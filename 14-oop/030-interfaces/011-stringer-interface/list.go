package main

import (
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry, currently there are no books\n"
	}
	var str strings.Builder

	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, item := range l {
		item.discount(ratio)
	}
}
