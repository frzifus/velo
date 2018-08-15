package velo

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

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

type Slots []Slot

type slotsWrapper struct {
	Slots `json:"stationSlots"`
}

func SlotsByStationID(stationID int) (Slots, error) {
	var wrapper slotsWrapper
	url := "https://cms.velocity-aachen.de/backend/app/stations/"
	url += strconv.Itoa(stationID) + "/slots"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return wrapper.Slots, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		log.Fatalln(err)
		return wrapper.Slots, err
	}
	return wrapper.Slots, nil
}

func Stations() ([]Station, error) {
	var stations []Station
	resp, err := http.Get("https://cms.velocity-aachen.de/backend/app/stations")
	if err != nil {
		log.Fatalln(err)
		return stations, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
		log.Fatalln(err)
		return stations, err
	}
	return stations, nil
}
