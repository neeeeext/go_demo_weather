package weather_test

import (
	"app/geo"
	"app/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {

	city := "Paris"
	expected := geo.GeoData{
		City: city,
	}
	format := 3

	result, _ := weather.GetWeather(expected, format)
	if !strings.Contains(result, expected.City) {
		t.Errorf("Ожидалось %v, полученно %v", expected, result)
	}

}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 128},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherNotCorrectFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}

			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrNoCorrectFormat {
				t.Errorf("Ожидалось %v, полученно %v", weather.ErrNoCorrectFormat, err)
			}
		})
	}
}
