package main

import (
	"fmt"
	"io"

	"github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

//
// 0 = empty
// S = ship
// X = hit ship

const (
	Nothing = 0
	Hit     = 1
	Miss    = 2
)

// Board exported
type TheirBoard [][]int

// type OurBoard [][]*Ship

// var b1 Board // Our Board
// var b2 Board // Opponents Board

// Display the baord
// b1 our board = b2 their board
func Display(b1, b2 TheirBoard, writer io.Writer) {

	fmt.Println("Display")

	// Clear the terminal
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	xStart := 10
	yStart := 10

	print_tb(5, 5, "Battle Ships", termbox.ColorMagenta, termbox.ColorWhite)

	printBoard(b1, xStart, yStart)
	printBoard(b2, xStart+30, yStart)

}

func printBoard(board TheirBoard, xStart, yStart int) {

	print_tb(xStart-1, 9, "/ - - - - - - - - - -\\", termbox.ColorBlack, termbox.ColorWhite)

	// X then Y
	// y := yStart
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			termbox.SetCell(xStart-1, yStart+y, '|', termbox.ColorBlack, termbox.ColorWhite)

			switch board[x][y] {
			case 1:
				termbox.SetCell(xStart+(x*2), yStart+y, 'X', termbox.ColorRed, termbox.ColorBlack)
			case 2:
				termbox.SetCell(xStart+(x*2), yStart+y, ' ', 8, termbox.ColorBlack)

			default:
				termbox.SetCell(xStart+(x*2), yStart+y, '~', termbox.ColorGreen, termbox.ColorCyan)
				termbox.SetCell(xStart+(x*2)-1, yStart+y, ' ', termbox.ColorGreen, termbox.ColorCyan)
			}
		}
		termbox.SetCell(xStart+(10*2), yStart+y, '|', termbox.ColorBlack, termbox.ColorWhite)
	}
	termbox.Flush()

	print_tb(xStart-1, yStart+10, "\\ - - - - - - - - - -/", termbox.ColorBlack, termbox.ColorWhite)
}

func print_tb(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, c := range []rune(msg) {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
	termbox.Flush()
}
