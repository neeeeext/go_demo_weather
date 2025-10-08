package main

import (
	"app/geo"
	"app/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("*** Прогноз погоды ***")
	fmt.Println("")

	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(weatherData)
}
