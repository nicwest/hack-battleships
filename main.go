package main

import (
	"encoding/json"
	"io"
	"net/http"
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

func main() {
	http.HandleFunc("/", TakingFire)
	http.ListenAndServe(":8000", nil)
}
