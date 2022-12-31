package main

import (
	"sort"
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

// Len is the number of elements in the collection.
func (l list) Len() int {
	return len(l)
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title // compare by title
}

// Swap swaps the elements with indexes i and j.
func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byRelease struct {
	list
}

func (br byRelease) Less(i, j int) bool {
	return br.list[i].Released.Before(br.list[j].Released.Time) // compare by released
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}
