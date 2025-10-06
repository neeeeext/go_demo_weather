package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type isCityResponce struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("Такого города нет !")
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
	postBody, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	resp, err := http.Post("https://countriesnow.space/api/v0.1/tountries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var cityTrue isCityResponce
	json.Unmarshal(body, &cityTrue)
	return !cityTrue.Error
}
