package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

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

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	http.HandleFunc("/", TakingFire)
	go http.ListenAndServe(":8000", nil)

	for {
		Display(ourBoard, theirBoard, os.Stdout)

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeySpace {
				os.Exit(0)
			}
		}

		time.Sleep(1000)
	}
}
