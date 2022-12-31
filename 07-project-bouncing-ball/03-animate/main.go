package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	steps := 1000
	width, height := 50, 10

	xv, yv := 1, 1
	x, y := 0, 0
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

	//board[0][0] = true
	// len(board) -> width -> 5
	// len(board[0]) -> height -> 3
	//fmt.Println(len(board[0]))

	// Clear all the characters on the screen
	screen.Clear()

	for i := 0; i < steps; i++ {
		// width x height = 4x4, if true -> revert vertical direction
		// Moves the cursor to the top-left position of the screen

		// calculate the next position of the ball
		// Velocity means: Speed and Direction
		// X velocity = 1 -> ball moves right
		// X velocity = -1 -> ball moves left
		// Y velocity = 1 -> ball moves down
		// Y velocity = -1 -> ball moves up

		if y == height-1 {
			yv = -1
		}

		if y == 0 {
			yv = 1
		}

		if x == width-1 {
			xv = -1
		}

		if x == 0 {
			xv = 1
		}
		// remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// set a next ball position
		x += xv
		y += yv
		board[x][y] = true

		// Animate the time always in the same position
		// draw the board into []rune buffer

		// print []rune buffer: string(buffer)
		// rewind the buffer so that the program reuses it
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell := cEmpty
				if board[x][y] {
					cell = cBall
				}
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		// print the buffer
		fmt.Print(string(buf))
		//wait before drawing the next step
		time.Sleep(time.Second / 10)
	}
}
