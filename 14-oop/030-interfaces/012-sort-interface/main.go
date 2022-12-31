package main

import (
	"fmt"
	"sort"
)

func main() {
	l := list{
		{title: "Moby Dick", price: 10, released: toTimestamp(1672299339)},
		{title: "Little prince", price: 11, released: toTimestamp("1671953739")},
		{title: "Star Wars II", price: 40},
		{title: "Lego", price: 60},
	}
	sort.Sort(l)               // list has to satisfy the sort.Interface
	sort.Sort(sort.Reverse(l)) // reverse order of the sort
	sort.Sort(byReleaseDate(l))
	l.discount(0.5)
	fmt.Println(l)

}
