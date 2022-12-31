package main

import "fmt"

func main() {
	const pageSize int = 4

	games := []string{
		"Tetris", "Pacman", "Sims", "Heroes",
		"Age of Empire", "Doom", "Zybex", "Sim City",
		"Super Mario",
	}

	total := len(games)
	for from := 0; from < total; from += pageSize {
		to := from + pageSize

		if to > total {
			to = total
		}

		fmt.Printf("Page: %d ", from/pageSize+1)

		currentPage := games[from:to]
		fmt.Printf("%#v\n", currentPage)
	}
}
