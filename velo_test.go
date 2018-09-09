package velo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
)

const fileFormatString = "testdata/%s"

func TestSlotsByStationID(t *testing.T) {
	tt := []struct {
		name      string
		stationID int
		want      []Slot
		expected  bool
	}{
		// TODO: Add test cases
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SlotsByStationID(tc.stationID)
			if (err != nil) != tc.expected {
				t.Errorf("SlotsByStationID() error = %v, expected %v", err, tc.expected)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SlotsByStationID() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStations(t *testing.T) {
	filename := "stations_valid.json"
	validFile, err := os.Open(fmt.Sprintf(fileFormatString, filename))
	defer validFile.Close()
	if err != nil {
		t.Error(err)
		return
	}
	validStations := []Station{}
	data, err := ioutil.ReadFile(fmt.Sprintf(fileFormatString, filename))
	if err != nil {
		t.Error(err)
		return
	}

	err = json.Unmarshal(data, &validStations)
	if err != nil {
		t.Error(err)
		return
	}

	filename = "stations_empty.json"
	emptyFile, err := os.Open(fmt.Sprintf(fileFormatString, filename))
	defer emptyFile.Close()
	if err != nil {
		t.Error("missing file", filename)
	}

	tt := []struct {
		name     string
		get      func(string) (*http.Response, error)
		want     []Station
		expected bool
	}{
		// TODO: add more test cases
		{
			name: "station response",
			get: func(url string) (*http.Response, error) {
				return &http.Response{
					Body: ioutil.NopCloser(bufio.NewReader(validFile)),
				}, nil
			},
			want:     validStations,
			expected: true,
		},
		{
			name: "station response empty",
			get: func(url string) (*http.Response, error) {
				return &http.Response{
					Body: ioutil.NopCloser(bufio.NewReader(emptyFile)),
				}, nil
			},
			want:     []Station{},
			expected: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client.Get = tc.get
			got, err := Stations()
			if (err != nil) == tc.expected {
				t.Errorf("Stations() error = %v, expected %v", err, tc.expected)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Stations() = %v, want %v", got, tc.want)
			}
		})
	}
}
