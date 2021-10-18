package fitbit

import (
	"encoding/json"
	"time"
)

// https://dev.fitbit.com/build/reference/web-api/activity/

const (
	ActivitiesCalories             = "activities/calories"
	ActivitiesCaloriesBMR          = "activities/caloriesBMR"
	ActivitiesSteps                = "activities/steps"
	ActivitiesDistance             = "activities/distance"
	ActivitiesFloors               = "activities/floors"
	ActivitiesElevation            = "activities/elevation"
	ActivitiesMinutesSedentary     = "activities/minutesSedentary"
	ActivitiesMinutesLightlyActive = "activities/minutesLightlyActive"
	ActivitiesMinutesFairlyActive  = "activities/minutesFairlyActive"
	ActivitiesMinutesVeryActive    = "activities/minutesVeryActive"
	ActivitiesMctivityCalories     = "activities/activityCalories"

	ActivitiesTrackerCalories             = "activities/tracker/calories"
	ActivitiesTrackerSteps                = "activities/tracker/steps"
	ActivitiesTrackerDistance             = "activities/tracker/distance"
	ActivitiesTrackerFloors               = "activities/tracker/floors"
	ActivitiesTrackerElevation            = "activities/tracker/elevation"
	ActivitiesTrackerMinutesSedentary     = "activities/tracker/minutesSedentary"
	ActivitiesTrackerMinutesLightlyActive = "activities/tracker/minutesLightlyActive"
	ActivitiesTrackerMinutesFairlyActive  = "activities/tracker/minutesFairlyActive"
	ActivitiesTrackerMinutesVeryActive    = "activities/tracker/minutesVeryActive"
	ActivitiesTrackerActivityCalories     = "activities/tracker/activityCalories"
)

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

// ActivitiesValue contains a single record of an activity value.
type ActivitiesValue struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

// ActivitiesDay contains user activity logs, only one dataset is used and the other ones are empty.
type ActivitiesDay struct {
	ActivitiesTrackerSteps                []ActivitiesValue `json:"activities-tracker-steps,omitempty"`
	ActivitiesTrackerCalories             []ActivitiesValue `json:"activities-tracker-calories,omitempty"`
	ActivitiesTrackerDistance             []ActivitiesValue `json:"activities-tracker-distance,omitempty"`
	ActivitiesTrackerFloors               []ActivitiesValue `json:"activities-tracker-floors,omitempty"`
	ActivitiesTrackerElevation            []ActivitiesValue `json:"activities-tracker-elevation,omitempty"`
	ActivitiesTrackerMinutesSedentary     []ActivitiesValue `json:"activities-tracker-minutesSedentary,omitempty"`
	ActivitiesTrackerMinutesLightlyActive []ActivitiesValue `json:"activities-tracker-minutesLightlyActive,omitempty"`
	ActivitiesTrackerMinutesFairlyActive  []ActivitiesValue `json:"activities-tracker-minutesFairlyActive,omitempty"`
	ActivitiesTrackerMinutesVeryActive    []ActivitiesValue `json:"activities-tracker-minutesVeryActive,omitempty"`
	ActivitiesTrackerActivityCalories     []ActivitiesValue `json:"activities-tracker-activityCalories,omitempty"`
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
func (f *Fitbit) ActivitiesDayJSON(day string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/date/" + day + ".json")
}
