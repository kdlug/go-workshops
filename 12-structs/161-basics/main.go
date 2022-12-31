package main

import "fmt"

func main() {
	type person struct {
		name, lastname string
		age            int
	}

	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      91,
	}

	var freud person
	freud.name = "Zigmunt"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("Picasso: %+v\n", picasso)
	fmt.Printf("Freud: %#v\n", freud)
}
