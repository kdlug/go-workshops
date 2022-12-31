package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
func newGame(id int, price int, name, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

//
//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
//
func addGame(games []game, g game) []game {
	return append(games, g)
}

//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
func load() (games []game) {
	//var games []game // nil
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))

	return

}

//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//
func indexByID(games []game) (byID map[int]game) {
	byID = make(map[int]game) // we used named output so = instead of :=
	for _, g := range games {
		byID[g.id] = g
	}
	return
}

//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
