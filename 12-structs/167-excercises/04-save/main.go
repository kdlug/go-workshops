package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
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

type jGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
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
> save   : exports the data to json and quits
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

		case "save":
			var jgames []jGame

			for _, g := range games {
				jgame := jGame{
					ID:    g.id,
					Name:  g.name,
					Genre: g.genre,
					Price: g.price,
				}
				jgames = append(jgames, jgame)
			}
			out, err := json.MarshalIndent(jgames, "", "\t")
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(string(out))
			return
		}
	}

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}

}
