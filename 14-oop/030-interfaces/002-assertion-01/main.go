package main

func main() {
	var (
		mobydick     = book{title: "Moby Dick", price: 10}
		littleprince = book{title: "Little prince", price: 11}
		starwars     = game{title: "Star Wars II", price: 40}
		lego         = toy{title: "Lego", price: 60}
	)
	var store list

	store = append(store, mobydick, littleprince, &starwars, &lego)

	store.print()
	store.discount(0.5)

}
