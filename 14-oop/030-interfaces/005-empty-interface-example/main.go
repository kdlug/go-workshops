package main

func main() {
	var (
		starwars = game{title: "Star Wars II", price: 40}
	)

	var any interface{}

	any = &starwars

	// empty interface hides all the methods
	// any.print undefined (type interface{} has no field or method print)
	// any.print()

}
