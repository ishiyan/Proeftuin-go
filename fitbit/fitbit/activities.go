package fitbit

import (
	"encoding/json"
	"fitbit/fitbit/activities"
	"time"
)

// https://dev.fitbit.com/build/reference/web-api/activity/

// ActivitiesDaySummary contains a summary of activities of a requested day.
type ActivitiesDaySummary struct {
	Activities []struct {
		ActivityID           int       `json:"activityId"`
		ActivityParentID     int       `json:"activityParentId"`
		ActivityParentName   string    `json:"activityParentName"`
		Calories             int       `json:"calories"`
		Description          string    `json:"description"`
		DetailsLink          string    `json:"detailsLink,omitempty"`
		Distance             float64   `json:"distance"`
		Duration             int       `json:"duration"`
		HasActiveZoneMinutes bool      `json:"hasActiveZoneMinutes"`
		HasStartTime         bool      `json:"hasStartTime"`
		IsFavorite           bool      `json:"isFavorite"`
		LastModified         time.Time `json:"lastModified"`
		LogID                int64     `json:"logId"`
		Name                 string    `json:"name"`
		StartDate            string    `json:"startDate"`
		StartTime            string    `json:"startTime"`
		Steps                int       `json:"steps"`
	} `json:"activities"`
	Goals struct {
		ActiveMinutes int     `json:"activeMinutes"`
		CaloriesOut   int     `json:"caloriesOut"`
		Distance      float64 `json:"distance"`
		Floors        int     `json:"floors"`
		Steps         int     `json:"steps"`
	} `json:"goals"`
	Summary struct {
		ActiveScore            int `json:"activeScore"`
		ActivityCalories       int `json:"activityCalories"`
		CalorieEstimationMu    int `json:"calorieEstimationMu"`
		CaloriesBMR            int `json:"caloriesBMR"`
		CaloriesOut            int `json:"caloriesOut"`
		CaloriesOutUnestimated int `json:"caloriesOutUnestimated"`
		Distances              []struct {
			Activity string  `json:"activity"`
			Distance float64 `json:"distance"`
		} `json:"distances"`
		Elevation           float64 `json:"elevation"`
		FairlyActiveMinutes int     `json:"fairlyActiveMinutes"`
		Floors              int     `json:"floors"`
		HeartRateZones      []struct {
			CaloriesOut float64 `json:"caloriesOut"`
			Max         int     `json:"max"`
			Min         int     `json:"min"`
			Minutes     int     `json:"minutes"`
			Name        string  `json:"name"`
		} `json:"heartRateZones"`
		LightlyActiveMinutes int  `json:"lightlyActiveMinutes"`
		MarginalCalories     int  `json:"marginalCalories"`
		RestingHeartRate     int  `json:"restingHeartRate"`
		SedentaryMinutes     int  `json:"sedentaryMinutes"`
		Steps                int  `json:"steps"`
		UseEstimation        bool `json:"useEstimation"`
		VeryActiveMinutes    int  `json:"veryActiveMinutes"`
	} `json:"summary"`
}

// ActivitiesDailyValue contains a single record of a daily activity value.
type ActivitiesDailyValue struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

// ActivitiesDay contains user daily activity logs.
// Only one dataset is used and the other ones are empty.
type ActivitiesDay struct {
	Steps                []ActivitiesDailyValue `json:"activities-steps,omitempty"`
	Calories             []ActivitiesDailyValue `json:"activities-calories,omitempty"`
	Distance             []ActivitiesDailyValue `json:"activities-distance,omitempty"`
	Floors               []ActivitiesDailyValue `json:"activities-floors,omitempty"`
	Elevation            []ActivitiesDailyValue `json:"activities-elevation,omitempty"`
	MinutesSedentary     []ActivitiesDailyValue `json:"activities-minutesSedentary,omitempty"`
	MinutesLightlyActive []ActivitiesDailyValue `json:"activities-minutesLightlyActive,omitempty"`
	MinutesFairlyActive  []ActivitiesDailyValue `json:"activities-minutesFairlyActive,omitempty"`
	MinutesVeryActive    []ActivitiesDailyValue `json:"activities-minutesVeryActive,omitempty"`
	ActivityCalories     []ActivitiesDailyValue `json:"activities-activityCalories,omitempty"`

	TrackerSteps                []ActivitiesDailyValue `json:"activities-tracker-steps,omitempty"`
	TrackerCalories             []ActivitiesDailyValue `json:"activities-tracker-calories,omitempty"`
	TrackerDistance             []ActivitiesDailyValue `json:"activities-tracker-distance,omitempty"`
	TrackerFloors               []ActivitiesDailyValue `json:"activities-tracker-floors,omitempty"`
	TrackerElevation            []ActivitiesDailyValue `json:"activities-tracker-elevation,omitempty"`
	TrackerMinutesSedentary     []ActivitiesDailyValue `json:"activities-tracker-minutesSedentary,omitempty"`
	TrackerMinutesLightlyActive []ActivitiesDailyValue `json:"activities-tracker-minutesLightlyActive,omitempty"`
	TrackerMinutesFairlyActive  []ActivitiesDailyValue `json:"activities-tracker-minutesFairlyActive,omitempty"`
	TrackerMinutesVeryActive    []ActivitiesDailyValue `json:"activities-tracker-minutesVeryActive,omitempty"`
	TrackerActivityCalories     []ActivitiesDailyValue `json:"activities-tracker-activityCalories,omitempty"`
}

// ActivitiesDailyValue contains a single record of an intra-day activity value.
type ActivitiesIntradayValue struct {
	Dataset []struct {
		Time  string  `json:"time"`
		Value float64 `json:"value"`
	} `json:"dataset,omitempty"`
	DatasetInterval int    `json:"datasetInterval,omitempty"`
	DatasetType     string `json:"datasetType,omitempty"`
}

// ActivitiesIntraday contains user intraday activity logs.
// Only one resource pair is used and the other ones are empty.
type ActivitiesIntraday struct {
	CaloriesDaily  []ActivitiesDailyValue  `json:"activities-calories,omitempty"`
	Calories       ActivitiesIntradayValue `json:"activities-calories-intraday,omitempty"`
	StepsDaily     []ActivitiesDailyValue  `json:"activities-steps,omitempty"`
	Steps          ActivitiesIntradayValue `json:"activities-steps-intraday,omitempty"`
	DistanceDaily  []ActivitiesDailyValue  `json:"activities-distance,omitempty"`
	Distance       ActivitiesIntradayValue `json:"activities-distance-intraday,omitempty"`
	FloorsDaily    []ActivitiesDailyValue  `json:"activities-floors,omitempty"`
	Floors         ActivitiesIntradayValue `json:"activities-floors-intraday,omitempty"`
	ElevationDaily []ActivitiesDailyValue  `json:"activities-elevation,omitempty"`
	Elevation      ActivitiesIntradayValue `json:"activities-elevation-intraday,omitempty"`
}

// ActivitiesDaySummaryJSON returns a summary of activities of a requested day as a raw JSON.
// Date must be in the format "yyyy-MM-dd".
func (f *Fitbit) ActivitiesDaySummaryJSON(day string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/date/" + day + ".json")
}

// ActivitiesDaySummary converts the raw JSON to the ActivitiesDay type.
func (f *Fitbit) ActivitiesDaySummary(jsn []byte) (ActivitiesDaySummary, error) {
	summary := ActivitiesDaySummary{}
	if err := json.Unmarshal(jsn, &summary); err != nil {
		return summary, err
	}

	return summary, nil
}

// ActivitiesDayJSON returns a time series of activities of a requested day as a raw JSON.
// Date must be in the format "yyyy-MM-dd".
func (f *Fitbit) ActivitiesDayJSON(day string, resourcePath activities.Daily) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/" + string(resourcePath) + "/date/" + day + "/1d.json")
}

// ActivitiesDay converts the raw JSON to the ActivitiesDay type.
func (f *Fitbit) ActivitiesDay(jsn []byte) (ActivitiesDay, error) {
	act := ActivitiesDay{}
	if err := json.Unmarshal(jsn, &act); err != nil {
		return act, err
	}

	return act, nil
}

// ActivitiesIntradayJSON returns a time series of activities of a requested day as a raw JSON.
// Date must be in the format "yyyy-MM-dd".
// Resolution is either "1min" or "15min".
func (f *Fitbit) ActivitiesIntradayJSON(day string, resourcePath activities.Intra, resolution string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/" + string(resourcePath) + "/date/" + day + "/1d/" + resolution + ".json")
}

// ActivitiesIntraday converts the raw JSON to the ActivitiesIntraday type.
func (f *Fitbit) ActivitiesIntraday(jsn []byte) (ActivitiesIntraday, error) {
	act := ActivitiesIntraday{}
	if err := json.Unmarshal(jsn, &act); err != nil {
		return act, err
	}

	return act, nil
}
