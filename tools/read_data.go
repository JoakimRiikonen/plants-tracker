package main

import (
	"encoding/json"
	"fmt"
	"jr/plants-tracker/pkg/readings"
	"net/http"
	"os"
	"strconv"
)

func main() {
	url := os.Args[1]
	fmt.Println("Reading data from url " + url)

	store := readings.NewReadingDbStore("../db.db")
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	var data LiveDataResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	for _, chSoil := range data.SoilData {
		moistureReading, err := strconv.ParseFloat(chSoil.Humidity[:len(chSoil.Humidity)-1], 64)
		if err != nil {
			// Data is missing (likely "--")
			moistureReading = -1
		}

		reading := readings.AddReading{
			SensorId: chSoil.Channel,
			Moisture: moistureReading,
		}

		fmt.Println(reading)

		store.Add(reading)
	}

}

type LiveDataResponse struct {
	IndoorData []Wh25Data   `json:"wh25"`
	SoilData   []ChSoilData `json:"ch_soil"`
}

type Wh25Data struct {
	Intemp string `json:"intemp"`
	Unit   string `json:"unit"`
	Inhumi string `json:"inhumi"`
	Abs    string `json:"abs"`
	Rel    string `json:"rel"`
}

type ChSoilData struct {
	Channel  string `json:"channel"`
	Name     string `json:"name"`
	Battery  string `json:"battery"`
	Voltage  string `json:"voltage"`
	Humidity string `json:"humidity"`
}
