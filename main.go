package main

import (
	"flag"
	"fmt"
	"weather_api/geo"
	"weather_api/weather"
)

func main() {

	city := flag.String("city", "", "User`s city")
	format := flag.Int("format", 1, "User`s weather format")

	flag.Parse()

	geolocationData, err := geo.GetMyLocation(*city)
	if err != nil {
		panic(err)
	}

	w, _ := weather.GetWeather(*geolocationData, *format)
	fmt.Println(w)

}
