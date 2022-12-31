package main

func main() {
	var (
		mobydick     = product{title: "Moby Dick", price: 10, released: toTimestamp(1672299339)}
		littleprince = product{title: "Little prince", price: 11, released: toTimestamp("1671953739")}
		starwars     = product{title: "Star Wars II", price: 40}
		lego         = product{title: "Lego", price: 60}
	)
	var store list

	store = append(store, &mobydick, &littleprince, &starwars, &lego)
	store.discount(0.5)
	store.print()

}
