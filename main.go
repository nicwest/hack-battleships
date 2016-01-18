package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"flag"
	"net/http"
	"log"
	"bytes"
	"os"
	"time"
	"github.com/nsf/termbox-go"
)

type Shot struct {
	X int `json:"x"`
	Y int `json:"x"`
}

var ourTurn = flag.Bool("starts", false, "Choose to start the game and you win the coin toss")
var ourBoard OurBoard
var theirBoard TheirBoard

const (
	AircraftCarrier = 0
	BattleShip = 1
	Cruiser = 2
	Destroyer1 = 3
	Destroyer2 = 4
	Submarine1 = 5
	Submarine2 = 6
)

type Response struct {
	Hit      bool `json:"hit"`
	ShipSunk int  `json:"ship,omitempty"`
}

func TakingFire(w http.ResponseWriter, r *http.Request) {
	log.Print(*ourTurn)
	if (*ourTurn) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Please wait your turn!")
	} else {
		response := Response{
			Hit:      true,
			ShipSunk: BattleShip,
		}
		j, _ := json.Marshal(response)
		io.WriteString(w, string(j))
	}
}

func PewPew(w http.ResponseWriter, r *http.Request) {
	if (!*ourTurn) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Please wait your turn!")
	} else {
		var shot Shot
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &shot)
		if (theirBoard[shot.X][shot.Y] == Nothing) {
			j, _ := json.Marshal(shot)
			res, err := http.Post("localhost:8001/", "application/json", bytes.NewBuffer(j))
			var response Response
			json.Unmarshal(ioutil.ReadAll(res.Body), &response)
			if response.Hit {

			}
		}
		response := Response{
			Hit:      true,
			ShipSunk: BattleShip,
		}
		j, _ := json.Marshal(response)
		io.WriteString(w, string(j))
	}
}

func newBoard() (tb TheirBoard) {
	tb = make(TheirBoard, 10)
	for i := 0; i < 10; i++ {
		tb[i] = make([]int, 10)
	}

	return
}

func main() {
	flag.Parse()

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
