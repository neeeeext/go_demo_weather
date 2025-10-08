package weather

import (
	"app/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrNoCorrectFormat = errors.New("указан не верный формат (формат от 1 до 4)")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	city := url.QueryEscape(geo.City)
	if format < 1 || format > 4 {
		return "", ErrNoCorrectFormat
	}
	url := fmt.Sprintf("https://wttr.in/%s?format=%d", city, format)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("NOT 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
