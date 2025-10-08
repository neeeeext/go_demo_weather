package geo_test

import (
	"app/geo"
	"testing"
)

// positive test
func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, полученно %v", expected, got)
	}
}

// negative test
func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonwww"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, полученно %v", geo.ErrNoCity, err)
	}
}
