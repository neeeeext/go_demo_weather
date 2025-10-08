package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GeoData struct {
	City string `json:"city"`
}

type isCityResponse struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("такого города нет")

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrNoCity
		}
		return &GeoData{
			City: city,
		}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	escapeCity := url.QueryEscape(city)
	apiUrl := "https://countriesnow.space/api/v0.1/countries/population/cities/q?city=" + escapeCity
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var cityTrue isCityResponse
	json.Unmarshal(body, &cityTrue)
	if err := json.Unmarshal(body, &cityTrue); err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return false
	}
	return !cityTrue.Error
}
