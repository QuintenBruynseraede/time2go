package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/QuintenBruynseraede/time2go/internal/models"
	"github.com/QuintenBruynseraede/time2go/internal/utils"
	"github.com/QuintenBruynseraede/time2go/internal/weather"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// location := r.FormValue("location")
	dateRange := r.FormValue("daterange")
	duration, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to case duration %v to int", duration), http.StatusBadRequest)
	}

	timeRange, err := utils.DateRangeInputToTimeRange(dateRange)
	if err != nil {
		http.Error(w, "Invalid date range", http.StatusBadRequest)
	}

	timeLabels := utils.GenerateTimeRangeLabels(timeRange)
	// TODO hardcoded to Paris coordinates
	forecast, err := weather.GetForecast(52.3676, 4.9041, timeRange)
	if err != nil {
		http.Error(w, "Unable to get Weather forecast", http.StatusInternalServerError)
		return
	}

	scores := weather.ScoreForecast(forecast, duration)
	best := fmt.Sprintf("Between %s and %s", timeLabels[scores.BestIndices[0]], timeLabels[scores.BestIndices[len(scores.BestIndices)-1]])

	data := models.Response{
		TimeRangeList: utils.FormatAsJSList(timeLabels, true),

		PrecipitationList: utils.FormatAsJSList(forecast.Precipitation, false),
		TemperatureList:   utils.FormatAsJSList(forecast.Temperatures, false),
		CloudCoverList:    utils.FormatAsJSList(forecast.CloudCover, false),

		PrecipitationScoreList: utils.FormatAsJSList(scores.PrecipitationScores, false),
		TemperatureScoreList:   utils.FormatAsJSList(scores.TemperatureScores, false),
		CloudCoverScoreList:    utils.FormatAsJSList(scores.CloudCoverScores, false),

		RecommendedMoment:           best,
		RecommendedTimeRangeIndices: utils.FormatAsJSList(scores.BestIndices, false),
	}

	err = app.templates.ExecuteTemplate(w, "response", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	locationInput := strings.ToLower(r.FormValue("location"))
	var searchResult []string

	if utf8.RuneCountInString(locationInput) < 3 {
		return
	} else {
		searchResult = app.trie.AutoComplete(locationInput)
	}

	if len(searchResult) > 5 {
		searchResult = searchResult[:5]
	}

	err := app.templates.ExecuteTemplate(w, "search_results", searchResult)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
