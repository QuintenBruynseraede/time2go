package models

type Response struct {
	TimeRangeList string

	// Actual values
	PrecipitationList string
	TemperatureList   string
	CloudCoverList    string
	// Scores
	PrecipitationScoreList string
	TemperatureScoreList   string
	CloudCoverScoreList    string
	// Recommended time
	RecommendedMoment           string
	RecommendedTimeRangeIndices string
}
