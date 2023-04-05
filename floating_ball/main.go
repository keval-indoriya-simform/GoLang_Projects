package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const (
	width     = 50
	height    = 10
	cellEmpty = ' '
	cellBall  = '*'
	maxFrames = 1200
	speed     = time.Second / 10
)

func main() {
	var cell rune
	var px, py int
	var vx = 1
	var vy = 1

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height)

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}

				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Println(string(buf))
		time.Sleep(speed)
	}
}
