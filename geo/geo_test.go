package geo_test

import (
	"errors"
	"testing"
	"weather_api/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "Barcelona"
	expected := geo.GeolocationData{
		City: city,
	}
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Error getting location")
	}
	if expected.City != got.City {
		t.Errorf("City mismatch, expected %v, got %v", expected.City, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Barcelonas"

	_, err := geo.GetMyLocation(city)
	if !errors.Is(err, geo.ErrorNoCity) {
		t.Errorf("Bad error returned, expected %v, got %v", geo.ErrorNoCity, err)
	}

}
