package main

func main() {
	var (
		mobydick     = book{product{"Moby Dick", 10}, 1672299339}
		littleprince = book{product{"Little prince", 11}, "1671953739"}
		starwars     = game{product{"Star Wars II", 40}}
		lego         = toy{product{"Lego", 60}}
	)
	var store list

	store = append(store, &mobydick, &littleprince, &starwars, &lego)
	store.discount(0.5)
	store.print()

}
