package main

func main() {
	mobydick := book{
		title: "Moby Dick",
		price: 10,
	}

	littleprince := book{
		title: "Little prince",
		price: 11,
	}

	starwars := game{
		title: "Star Wars II",
		price: 40,
	}
	mobydick.print() // automatically passes mobnydick to print()
	littleprince.print()
	starwars.print()

	// a different way of calling - by using a type name
	book.print(mobydick) // manually passes mobydicjk to print()

	// change price
	starwars.discount(0.5)
	starwars.print()
}
