package main

func main() {
	var (
		mobydick     = book{title: "Moby Dick", price: 10, published: 1672299339}
		littleprince = book{title: "Little prince", price: 11, published: "1671953739"}
		starwars     = book{title: "Star Wars II", price: 40}
	)
	var store list

	store = append(store, mobydick, littleprince, starwars)

	store.print()

}
