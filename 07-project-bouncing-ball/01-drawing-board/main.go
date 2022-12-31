package main

import "fmt"

func main() {
	width, height := 5, 3
	var (
		cBall  = '🥎'
		cEmpty = ' '
	)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}
	board[0][0] = true
	// len(board) -> width -> 5
	// len(board[0]) -> height -> 3
	//fmt.Println(len(board[0]))
	for x := range board {
		for y := range board[x] {
			cell := cEmpty
			if board[x][y] {
				cell = cBall
			}
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}
