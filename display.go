package main

import (
	"fmt"
	"io"

	termbox "github.com/nsf/termbox-go"
)

//
// 0 = empty
// S = ship
// X = hit ship

// Board exported
type Board [][]int

var b1 Board // Our Board
var b2 Board // Opponents Board

// Display the baord
// b1 our board = b2 their board
func Display(b1, b2, writer io.Writer) {

	// Clear the terminal
	termbox.Clear(termbox.ColorBlack, termbox.ColorWhite)

    xStart := 10

	print_tb(xStart, 10, "Battle Ships")
    print_tb(xStart, 11, "/----------\\")

	// X then Y
	for i := xStart; i < 10; i++ {
        switch b1
	}
}

func print_tb(x, y int, msg string) {
    for _, c := range []rune(msg) {
        termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
        x += runewidth.RuneWidth(c)
    }
    termbox.Flush()
}