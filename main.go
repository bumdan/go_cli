package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Timezone int `json:"timezone"`
}

const (
	Reset  = "\033[0m"
	Yellow = "\033[33m"
)

func main() {
	q := "Myanmar"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + q + "&units=Metric&appid=d12cf0c3579b00a6bcee8306d2e25d7e")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(Yellow + q + Reset + " Weather API not abailable")
	}

	body, err := io.ReadAll((res.Body))

	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC().Add(time.Duration(weather.Timezone) * time.Second)

	fmt.Printf(
		"Country: %s%s%s | Temp: %s%.0f°C%s | Hum: %s%.0f%s | Win-speed: %s%.0f%s | Weather: %s%s%s\n",
		Yellow, weather.Name, Reset,
		Yellow, weather.Main.Temp, Reset,
		Yellow, weather.Main.Humidity, Reset,
		Yellow, weather.Wind.Speed, Reset,
		Yellow, weather.Weather[0].Description, Reset,
	)

	fmt.Printf("Local Time: %s%s%s\n", Yellow, currentTime.Format("03:04 PM"), Reset)
}
