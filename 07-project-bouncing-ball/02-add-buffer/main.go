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

	// create a drawing buffer
	buf := make([]rune, 0, width*height)

	board[0][0] = true
	// len(board) -> width -> 5
	// len(board[0]) -> height -> 3
	//fmt.Println(len(board[0]))

	// use the loop for measuring the performance difference
	for i := 0; i < 1000; i++ {
		// rewind the buffer so that the program reuses it
		buf = buf[:0]

		for x := range board {
			for y := range board[x] {
				cell := cEmpty
				if board[x][y] {
					cell = cBall
				}
				//fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}

		// print the buffer
		fmt.Print(string(buf))
	}
}
