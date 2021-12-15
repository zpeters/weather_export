package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")
var zipCode = "39110"

type MainWeather struct {
	TempF    float32 `json:"temp"`
	TempFMax float32 `json:"temp_max"`
	TempFMin float32 `json:"temp_min"`
	Humidity float32 `json:"humidity"`
}

type Weather struct {
	Name string      `json:"name"`
	DT   int         `json:"dt"`
	Main MainWeather `json:"main"`
}

func main() {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?units=imperial&zip=%s,us&appid=%s", zipCode, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", body)

	var weather Weather
	json.Unmarshal(body, &weather)
	fmt.Printf("%#v\n", weather)
}
