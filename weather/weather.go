package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather_api/geo"
)

var ErrorInvalidFormat = errors.New("INVALID_FORMAT")

func GetWeather(geo geo.GeolocationData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorInvalidFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", errors.New("ERROR_URL")
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		return "", errors.New("ERROR_HTTP")
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", errors.New("ERROR_READ_BODY")
	}
	return string(body), nil
}
