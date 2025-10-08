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

func TestGetWeatherNotCorrectFormat(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}

	format := 13

	_, err := weather.GetWeather(geoData, format)
	if err != weather.ErrNoCorrectFormat {
		t.Errorf("Ожидалось %v, полученно %v", weather.ErrNoCorrectFormat, err)
	}
}
