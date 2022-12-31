package main

import "fmt"

func main() {
	games := []string{"Tetris", "Pacman", "Sims", "Heroes", "Age of Empire", "Doom"}

	top3 := games[0:3]
	// which is equivalent to
	// 	top3 := games[:3]

	fmt.Printf("Top 3: %#v\n", top3)

	last3 := games[3:]
	//last3 := games[len(games)-3:]
	fmt.Printf("Last 3: %#v\n", last3)

	mid := games[1 : len(games)-1]
	fmt.Printf("Mid: %#v\n", mid)

}
