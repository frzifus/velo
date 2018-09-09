package velo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	stationURL = "https://cms.velocity-aachen.de/backend/app/stations/"
	slotURL    = stationURL + "%d/slots"
)

var client struct {
	Get func(string) (*http.Response, error)
}

func init() {
	client.Get = http.Get
}

type Station struct {
	ID                int     `json:"stationId"`
	Name              string  `json:"name"`
	LocationLatitude  float64 `json:"locationLatitude"`
	LocationLongitude float64 `json:"locationLongitude"`
	State             string  `json:"state"`
	Note              string  `json:"note"`
	NumFreeSlots      int     `json:"numFreeSlots"`
	NumAllSlots       int     `json:"numAllSlots"`
}

type Slot struct {
	ID            int     `json:"stationSlotId"`
	Position      int     `json:"stationSlotPosition"`
	State         string  `json:"state"`
	IsOccupied    bool    `json:"isOccupied"`
	StateOfCharge float64 `json:"stateOfCharge"`
}

func SlotsByStationID(stationID int) ([]Slot, error) {
	var wrapper struct {
		Slots []Slot `json:"stationSlots"`
	}

	resp, err := client.Get(fmt.Sprintf(slotURL, stationID))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return wrapper.Slots, nil
}

func Stations() ([]Station, error) {
	var stations []Station
	resp, err := client.Get(stationURL)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return stations, nil
}
