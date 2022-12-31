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
	mobydick.print()
	littleprince.print()
	starwars.print()
}
