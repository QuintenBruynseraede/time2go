package main

import (
	"net/http"

	"github.com/QuintenBruynseraede/time2go/internal/models"
	"github.com/QuintenBruynseraede/time2go/internal/utils"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := app.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) form(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Handling request", "request", r.URL)

	timeRanges := []string{"00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"}
	actualPrecipitation := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	actualTemperatures := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9}
	actualCloudCover := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	precipitationScores := utils.Repeat(2, 24)
	temperatureScores := utils.Repeat(3, 24)
	cloudCoverScores := utils.Repeat(4, 24)

	data := models.Response{
		TimeRangeList: utils.FormatAsJSList(timeRanges, true),

		PrecipitationList: utils.FormatAsJSList(actualPrecipitation, false),
		TemperatureList:   utils.FormatAsJSList(actualTemperatures, false),
		CloudCoverList:    utils.FormatAsJSList(actualCloudCover, false),

		PrecipitationScoreList: utils.FormatAsJSList(precipitationScores, false),
		TemperatureScoreList:   utils.FormatAsJSList(temperatureScores, false),
		CloudCoverScoreList:    utils.FormatAsJSList(cloudCoverScores, false),

		RecommendedMoment:           "Friday between X and Y",
		RecommendedTimeRangeIndices: utils.FormatAsJSList([]int{0, 1, 2, 3, 4}, false),
	}

	err := app.templates.ExecuteTemplate(w, "response", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
