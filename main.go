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

type MainWeather struct {
	Temp     float32 `json:"temp"`
	TempMax  float32 `json:"temp_max"`
	TempMin  float32 `json:"temp_min"`
	Humidity float32 `json:"humidity"`
}

type Weather struct {
	Name string      `json:"name"`
	DT   int         `json:"dt"`
	Main MainWeather `json:"main"`
}

func getWeather(w http.ResponseWriter, req *http.Request) {
	zipCode := req.URL.Query().Get("zipcode")
	log.Printf("Got weather request from %s for %s\n", req.RemoteAddr, zipCode)
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

	var weather Weather
	json.Unmarshal(body, &weather)

	var output string
	output += fmt.Sprintf("weather_temp_f{location=\"%s\"} %f\n", weather.Name, weather.Main.Temp)
	output += fmt.Sprintf("weather_temp_max_f{location=\"%s\"} %f\n", weather.Name, weather.Main.TempMax)
	output += fmt.Sprintf("weather_temp_min_f{location=\"%s\"} %f\n", weather.Name, weather.Main.TempMin)
	output += fmt.Sprintf("weather_humidity{location=\"%s\"} %f", weather.Name, weather.Main.Humidity)

	log.Printf("Sending %s\n", output)
	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/metrics", getWeather)
	log.Println("Serving on port 9163")
	log.Fatal(http.ListenAndServe(":9163", nil))
}
