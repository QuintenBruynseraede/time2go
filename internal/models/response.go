package models

import "html/template"

type Response struct {
	TimeRangeList template.JS

	// Actual values
	PrecipitationList template.JS
	TemperatureList   template.JS
	CloudCoverList    template.JS
	// Scores
	PrecipitationScoreList template.JS
	TemperatureScoreList   template.JS
	CloudCoverScoreList    template.JS
	// Recommended time
	RecommendedMoment           string
	RecommendedTimeRangeIndices template.JS
	Location                    string
}
