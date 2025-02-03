package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeolocationData struct {
	City string `json:"city"`
}

type CityResponse struct {
	Error bool `json:"error"`
}

var ErrorNoCity = errors.New("city not found")
var ErrorBadRequest = errors.New("bad request")

func GetMyLocation(city string) (*GeolocationData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrorNoCity
		}
		return &GeolocationData{City: city}, nil
	}

	response, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, ErrorBadRequest

	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	var geo GeolocationData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}

func checkCity(city string) bool {

	body, err := json.Marshal(map[string]string{"city": city})
	if err != nil {
		return false
	}

	response, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	var cityResponse CityResponse
	err = json.Unmarshal(body, &cityResponse)
	if err != nil {
		return false
	}
	return !cityResponse.Error
}
