package main

import (
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
type OurBoard [][]*Ship

var b1 Board // Our Board
var b2 Board // Opponents Board

// Display the baord
// b1 our board = b2 their board
func Display(b1, b2, writer io.Writer) {

	// Clear the terminal
	termbox.Clear(termbox.ColorBlack, termbox.ColorWhite)

	xStart := 10
	yStart := 10

	print_tb(5, 5, "Battle Ships")

	printBoard(b1, xStart, yStart)
	printBoard(b2, xStart+30, yStart)

}

func printBoard(board Board, xStart, yStart int) {

	print_tb(xStart, 11, "/----------\\", termbox.ColorBlack, termbox.ColorWhite)

	// X then Y
	x := xStart
	for y := yStart; y < 10; y++ {
		for ; x < 10; x++ {
			termbox.SetCell(xStart-1, y, "|", termbox.ColorBlack, termbox.ColorWhite)
			switch b1[i] {
			case 1:
				termbox.SetCell(x, y, "X", termbox.ColorRed, termbox.ColorBlack)
			default:
				termbox.SetCell(x, y, "~", termbox.ColorBlack, termbox.ColorCyan)
			}
			termbox.SetCell(x+1, y, "|", termbox.ColorBlack, termbox.ColorWhite)
		}
	}

	print_tb(xStart, y+1, "\\----------/", termbox.ColorBlack, termbox.ColorWhite)
}

func print_tb(x, y int, msg string) {
	for _, c := range []rune(msg) {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		x += runewidth.RuneWidth(c)
	}
	termbox.Flush()
}
