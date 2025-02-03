package weather_test

import (
	"errors"
	"strings"
	"testing"
	"weather_api/geo"
	"weather_api/weather"
)

func TestGetWeather(t *testing.T) {
	expectedCity := "Moscow"
	geolocationData := geo.GeolocationData{
		City: expectedCity,
	}

	format := 3
	result, err := weather.GetWeather(geolocationData, format)
	if err != nil {
		t.Errorf("Error while getting weather data: %v", err)
	}
	if !strings.Contains(result, expectedCity) {
		t.Errorf("Failed test, expected %v, got %v", expectedCity, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "147", format: 147},
	{name: "0", format: 0},
	{name: "-10", format: -10},
}

func TestGetWeatherWithInvalidFormat(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			expectedCity := "Barcelona"
			geolocationData := geo.GeolocationData{
				City: expectedCity,
			}

			_, err := weather.GetWeather(geolocationData, testCase.format)
			if !errors.Is(err, weather.ErrorInvalidFormat) {
				t.Errorf("Failed test, error expected %v, got %v", weather.ErrorInvalidFormat, err)
			}
		})
	}
}
