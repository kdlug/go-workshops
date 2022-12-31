package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func main() {
	l := list{
		{Title: "Moby Dick", Price: 10, Released: toTimestamp(1672299339)},
		{Title: "Little prince", Price: 11, Released: toTimestamp("1671953739")},
		{Title: "Star Wars II", Price: 40},
		{Title: "Lego", Price: 60},
	}
	sort.Sort(l)               // list has to satisfy the sort.Interface
	sort.Sort(sort.Reverse(l)) // reverse order of the sort
	sort.Sort(byReleaseDate(l))
	l.discount(0.5)

	data, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
