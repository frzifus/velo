package main

import (
	"fmt"

	"git.klimlive.de/frzifus/velo"
)

func main() {
	stations, err := velo.Stations()
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		fmt.Println(station)
		slots, err := velo.SlotsByStationID(station.ID)
		if err != nil {
			panic(err)
		}
		for _, slot := range slots {
			fmt.Println(slot)
		}
		fmt.Println()
	}
}
