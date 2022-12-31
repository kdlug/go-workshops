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
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 4,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

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
	// load games from json
	var decoded []jGame
	err := json.Unmarshal([]byte(data), &decoded)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, d := range decoded {
		games = append(games, game{item{d.ID, d.Name, d.Price}, d.Genre})
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
