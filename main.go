package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type cursorPosition struct {
	x, y int
}

type Shot struct {
	X int `json:"x"`
	Y int `json:"x"`
}

const (
	AircraftCarrier = 0
	BattleShip      = 1
	Cruiser         = 2
	Destroyer1      = 3
	Destroyer2      = 4
	Submarine1      = 5
	Submarine2      = 6
)

type Response struct {
	Hit      bool `json:"hit"`
	ShipSunk int  `json:"ship,omitempty"`
}

func TakingFire(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Hit:      true,
		ShipSunk: BattleShip,
	}
	j, _ := json.Marshal(response)
	io.WriteString(w, string(j))
}

func newBoard() (tb TheirBoard) {
	tb = make(TheirBoard, 10)
	for i := 0; i < 10; i++ {
		tb[i] = make([]int, 10)
	}

	return
}

func main() {

	termbox.Init()
	defer termbox.Close()

	// Create a fake board
	ourBoard := newBoard()
	theirBoard := newBoard()

	// Initial Cursor Position
	cursor := cursorPosition{0, 0}

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	http.HandleFunc("/", TakingFire)
	go http.ListenAndServe(":8000", nil)

	for {
		Display(ourBoard, theirBoard, os.Stdout)

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowRight:
				cursor.x++
				if cursor.x > 9 {
					cursor.x = 9
				}
			case termbox.KeyArrowLeft:
				cursor.x--
				if cursor.x < 0 {
					cursor.x = 0
				}
			case termbox.KeyArrowUp:
				cursor.y--
				if cursor.y < 0 {
					cursor.y = 0
				}
			case termbox.KeyArrowDown:
				cursor.y++
				if cursor.y > 9 {
					cursor.y = 9
				}

			}

			termbox.SetCursor(cursor.x, cursor.y)

			if ev.Key == termbox.KeyArrowRight {
				os.Exit(0)
			}
		}

		time.Sleep(200)
	}
}
