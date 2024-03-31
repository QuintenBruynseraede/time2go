package weather

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/QuintenBruynseraede/time2go/internal/location"
	"github.com/QuintenBruynseraede/time2go/internal/timerange"
	"github.com/QuintenBruynseraede/time2go/internal/utils"
)

func GetForecast(location location.Location, timerange timerange.TimeRange) (Forecast, error) {
	// Returns a forecast for all hours within the timerange
	url := buildRequestUrl(location)
	raw, error := utils.MakeRequest(url)

	if error != nil {
		return Forecast{}, error
	}

	response := APIResponse{}
	err := json.Unmarshal(raw, &response)
	if err != nil {
		return Forecast{}, err
	}

	forecast, error := filterResponseToTimeRange(response.Hourly, timerange)
	if error != nil {
		return Forecast{}, error
	}

	return forecast, nil
}

type APIResponse struct {
	Hourly HourlyResponse `json:"hourly"`
}

type HourlyResponse struct {
	Time          []string  `json:"time"`
	Temperature   []float32 `json:"temperature_2m"`
	Precipitation []float32 `json:"precipitation"`
	CloudCover    []float32 `json:"cloud_cover"`
}

type Forecast struct {
	Temperatures  []int // Temperature in degrees Celsius
	CloudCover    []int // Cloud cover percentage
	Precipitation []int // Precipitation in mm
}

func buildRequestUrl(location location.Location) string {
	BASE_URL := "https://api.open-meteo.com/v1/forecast"
	return fmt.Sprintf(
		`%s?latitude=%v&longitude=%v&hourly=temperature_2m,precipitation,cloud_cover&timezone=auto`,
		BASE_URL,
		location.Latitude,
		location.Longitude,
	)
}

func filterResponseToTimeRange(response HourlyResponse, timerange timerange.TimeRange) (Forecast, error) {
	// Build a forecast by removing all values outside of the timerange
	temperatures := []int{}
	cloudCover := []int{}
	precipitation := []int{}

	for i, hour := range response.Time {
		parsedTime, err := time.Parse("2006-01-02T15:04", hour)
		if err != nil {
			return Forecast{}, err
		}

		if parsedTime.Before(timerange.Start) {
			continue
		} else if parsedTime.After(timerange.End) {
			break
		} else {
			temperatures = append(temperatures, int(response.Temperature[i]))
			cloudCover = append(cloudCover, int(response.CloudCover[i]))
			precipitation = append(precipitation, int(response.Precipitation[i]))
		}
	}

	return Forecast{Temperatures: temperatures, CloudCover: cloudCover, Precipitation: precipitation}, nil
}
