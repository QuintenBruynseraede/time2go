package weather

import (
	"math"

	"github.com/QuintenBruynseraede/time2go/internal/utils"
)

type ScoredForecast struct {
	TemperatureScores   []int
	CloudCoverScores    []int
	PrecipitationScores []int
	BestIndices         []int
}

func ScoreForecast(forecast Forecast, duration int) ScoredForecast {
	temperatureScores := utils.Map(forecast.Temperatures, scoreTemperatureValue)
	cloudCoverScores := utils.Map(forecast.CloudCover, scoreCloudCoverValue)
	precipitationScores := utils.Map(forecast.Precipitation, scorePrecipitation)
	scores := getScoreTotal(temperatureScores, cloudCoverScores, precipitationScores)

	return ScoredForecast{
		TemperatureScores:   temperatureScores,
		CloudCoverScores:    cloudCoverScores,
		PrecipitationScores: precipitationScores,
		BestIndices:         findHighestConsecutiveScores(scores, duration),
	}
}

func findHighestConsecutiveScores(scores []int, length int) []int {
	// Return a slice of consecutive indices whose sum of scores is the highest
	max := -1
	bestIndices := []int{}
	for i := range scores {
		if i+length > len(scores) {
			break
		}
		sum := 0
		for _, score := range scores[i : i+length] {
			sum += score
		}
		if sum > max {
			max = sum
			bestIndices = utils.Between(i, i+length)
		}
	}
	return bestIndices
}

func getScoreTotal(temps []int, cloudcover []int, prec []int) []int {
	scores := []int{}
	for i := range temps {
		scores = append(scores, temps[i]+cloudcover[i]+prec[i])
	}
	return scores
}

func scoreTemperatureValue(temp int) int {
	if temp > 25 {
		return 10 - (temp - 25)
	}
	if temp < 15 {
		return int(math.Max(0, float64(10-(15-temp))))
	} else {
		return 10
	}
}

func scoreCloudCoverValue(cloudCover int) int {
	if cloudCover < 33 {
		return 5
	}
	if cloudCover < 66 {
		return 3
	}
	return 0
}

func scorePrecipitation(precipitation int) int {
	return int(math.Max(0, float64(10-precipitation*3)))
}
