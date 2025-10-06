package weather

import (
	"app/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	city := url.QueryEscape(geo.City)
	url := fmt.Sprintf("https://wttr.in/%s?format=%d", city, format)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(errors.New("NOT 200"))
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
