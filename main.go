package main

import (
        "fmt"
        "os"
)

var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")

func main() {
        fmt.Println(apiKey)
}
