package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> list   : lists all the games
> id #	 : query by id
> quit   : quits
`)
		if !in.Scan() {
			break
		}

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}
			id, err := strconv.Atoi(cmd[1])

			if err != nil {
				fmt.Println("wrong id")
				continue
			}
			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
			break
		case "quit":
			fmt.Println("bye!")
			return
		}
	}

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}

}
