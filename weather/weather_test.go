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

	result := weather.GetWeather(expected, format)
	if !strings.Contains(result, expected.City) {
		t.Errorf("Ожидалось %v, полученно %v", expected, result)
	}

}
