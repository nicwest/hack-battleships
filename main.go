package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"flag"
	"net/http"
	"log"
	"bytes"
)

type Shot struct {
	X int `json:"x"`
	Y int `json:"x"`
}

var ourTurn = flag.Bool("starts", false, "Choose to start the game and you win the coin toss")

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
		if (b2[shot.X][shot.Y] == Nothing) {
			j, _ := json.Marshal(shot)
			res, err := http.Post("localhost:8001/", "application/json", bytes.NewBuffer(j))
		}
		response := Response{
			Hit:      true,
			ShipSunk: BattleShip,
		}
		j, _ := json.Marshal(response)
		io.WriteString(w, string(j))
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", TakingFire)
	http.ListenAndServe(":8000", nil)
}
