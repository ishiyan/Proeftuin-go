package fitbit

import (
	"encoding/json"
	"fitbit/fitbit/activities"
	"net/url"
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

// ActivitiesTypes is a tree of all valid Fitbit public activities from the activities catalog
// as well as private custom activities the user created in the format requested.
// If the activity has levels, also gets a list of activity level details.
type ActivitiesTypes struct {
	Categories []struct {
		Activities []struct {
			AccessLevel    string `json:"accessLevel"`
			ActivityLevels []struct {
				ID          int64   `json:"id"`
				MaxSpeedMPH float64 `json:"maxSpeedMPH"`
				Mets        float64 `json:"mets"`
				MinSpeedMPH float64 `json:"minSpeedMPH"`
				Name        string  `json:"name"`
			} `json:"activityLevels,omitempty"`
			HasSpeed bool    `json:"hasSpeed"`
			ID       int     `json:"id"`
			Name     string  `json:"name"`
			Mets     float64 `json:"mets,omitempty"`
		} `json:"activities"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SubCategories []struct {
			Activities []struct {
				AccessLevel string  `json:"accessLevel"`
				HasSpeed    bool    `json:"hasSpeed"`
				ID          int64   `json:"id"`
				Mets        float64 `json:"mets"`
				Name        string  `json:"name"`
			} `json:"activities"`
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"subCategories,omitempty"`
	} `json:"categories"`
}

// ActivitiesLifetimeStats contains the account lifetime statistics.
type ActivitiesLifetimeStats struct {
	Best struct {
		Total struct {
			Distance struct {
				Date  string  `json:"date"`
				Value float64 `json:"value"`
			} `json:"distance"`
			Floors struct {
				Date  string  `json:"date"`
				Value float64 `json:"value"`
			} `json:"floors"`
			Steps struct {
				Date  string `json:"date"`
				Value int64  `json:"value"`
			} `json:"steps"`
		} `json:"total"`
		Tracker struct {
			Distance struct {
				Date  string  `json:"date"`
				Value float64 `json:"value"`
			} `json:"distance"`
			Floors struct {
				Date  string  `json:"date"`
				Value float64 `json:"value"`
			} `json:"floors"`
			Steps struct {
				Date  string `json:"date"`
				Value int64  `json:"value"`
			} `json:"steps"`
		} `json:"tracker"`
	} `json:"best"`
	Lifetime struct {
		Total struct {
			ActiveScore float64 `json:"activeScore"`
			CaloriesOut float64 `json:"caloriesOut"`
			Distance    float64 `json:"distance"`
			Floors      int64   `json:"floors"`
			Steps       int64   `json:"steps"`
		} `json:"total"`
		Tracker struct {
			ActiveScore float64 `json:"activeScore"`
			CaloriesOut float64 `json:"caloriesOut"`
			Distance    float64 `json:"distance"`
			Floors      int64   `json:"floors"`
			Steps       int64   `json:"steps"`
		} `json:"tracker"`
	} `json:"lifetime"`
}

// ActivitiesGoal contains the activities goal of an user.
type ActivitiesGoal struct {
	Goals struct {
		ActiveMinutes int     `json:"activeMinutes,omitempty"`
		CaloriesOut   int     `json:"caloriesOut,omitempty"`
		Distance      float64 `json:"distance"`
		Floors        int     `json:"floors"`
		Steps         int     `json:"steps"`
	} `json:"goals"`
}

type ActivitiesList struct {
	Activities []struct {
		ActiveDuration    int `json:"activeDuration"`
		ActiveZoneMinutes struct {
			MinutesInHeartRateZones []struct {
				MinuteMultiplier int    `json:"minuteMultiplier"`
				Minutes          int    `json:"minutes"`
				Order            int    `json:"order"`
				Type             string `json:"type"`
				ZoneName         string `json:"zoneName"`
			} `json:"minutesInHeartRateZones"`
			TotalMinutes int `json:"totalMinutes"`
		} `json:"activeZoneMinutes,omitempty"`
		ActivityLevel []struct {
			Minutes int    `json:"minutes"`
			Name    string `json:"name"`
		} `json:"activityLevel"`
		ActivityName          string    `json:"activityName"`
		ActivityTypeID        int       `json:"activityTypeId"`
		Calories              int       `json:"calories"`
		CaloriesLink          string    `json:"caloriesLink"`
		Distance              float64   `json:"distance"`
		DistanceUnit          string    `json:"distanceUnit"`
		Duration              int       `json:"duration"`
		ElevationGain         float64   `json:"elevationGain"`
		HasActiveZoneMinutes  bool      `json:"hasActiveZoneMinutes"`
		LastModified          time.Time `json:"lastModified"`
		LogID                 int64     `json:"logId"`
		LogType               string    `json:"logType"`
		ManualValuesSpecified struct {
			Calories bool `json:"calories"`
			Distance bool `json:"distance"`
			Steps    bool `json:"steps"`
		} `json:"manualValuesSpecified"`
		OriginalDuration  int       `json:"originalDuration"`
		OriginalStartTime time.Time `json:"originalStartTime"`
		PoolLength        int       `json:"poolLength,omitempty"`
		PoolLengthUnit    string    `json:"poolLengthUnit,omitempty"`
		Source            struct {
			ID              string   `json:"id"`
			Name            string   `json:"name"`
			TrackerFeatures []string `json:"trackerFeatures"`
			Type            string   `json:"type"`
			URL             string   `json:"url"`
		} `json:"source"`
		Speed            float64   `json:"speed"`
		StartTime        time.Time `json:"startTime"`
		AverageHeartRate int       `json:"averageHeartRate,omitempty"`
		DetailsLink      string    `json:"detailsLink,omitempty"`
		HeartRateLink    string    `json:"heartRateLink,omitempty"`
		HeartRateZones   []struct {
			CaloriesOut float64 `json:"caloriesOut"`
			Max         int     `json:"max"`
			Min         int     `json:"min"`
			Minutes     int     `json:"minutes"`
			Name        string  `json:"name"`
		} `json:"heartRateZones,omitempty"`
		Pace    float64 `json:"pace,omitempty"`
		Steps   int     `json:"steps,omitempty"`
		TcxLink string  `json:"tcxLink,omitempty"`
	} `json:"activities"`
	Pagination struct {
		BeforeDate string `json:"beforeDate"`
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		Sort       string `json:"sort"`
	} `json:"pagination"`
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

// ActivitiesTypesJSON returns a summary of activity types as a raw JSON.
func (f *Fitbit) ActivitiesTypesJSON() ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/activities.json")
}

// ActivitiesTypes converts the raw JSON to the ActivitiesTypes type.
func (f *Fitbit) ActivitiesTypes(jsn []byte) (ActivitiesTypes, error) {
	types := ActivitiesTypes{}
	if err := json.Unmarshal(jsn, &types); err != nil {
		return types, err
	}

	return types, nil
}

// ActivitiesLifetimeStatsJSON returns a lifetime stats as a raw JSON.
func (f *Fitbit) ActivitiesLifetimeStatsJSON() ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities.json")
}

// ActivitiesLifetimeStats converts the raw JSON to the ActivitiesLifetimeStats type.
func (f *Fitbit) ActivitiesLifetimeStats(jsn []byte) (ActivitiesLifetimeStats, error) {
	stats := ActivitiesLifetimeStats{}
	if err := json.Unmarshal(jsn, &stats); err != nil {
		return stats, err
	}

	return stats, nil
}

// ActivitiesGoalJSON returns a lifetime stats as a raw JSON.
// Period is "daily" or "weekly".
func (f *Fitbit) ActivitiesGoalJSON(period string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/goals/" + period + ".json")
}

// ActivitiesGoal converts the raw JSON to the ActivitiesGoal type.
func (f *Fitbit) ActivitiesGoal(jsn []byte) (ActivitiesGoal, error) {
	goal := ActivitiesGoal{}
	if err := json.Unmarshal(jsn, &goal); err != nil {
		return goal, err
	}

	return goal, nil
}

// ActivitiesListJSON returns a 20 activities after the specified date as a raw JSON.
// Date must be in the format "yyyy-MM-dd".
func (f *Fitbit) ActivitiesListJSON(date string) ([]byte, error) {
	parameterList := url.Values{}
	parameterList.Add("afterDate", date)
	parameterList.Add("sort", "asc")
	parameterList.Add("limit", "20")
	parameterList.Add("offset", "0")

	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/list.json?" + parameterList.Encode())
}

// ActivitiesList converts the raw JSON to the ActivitiesList type.
func (f *Fitbit) ActivitiesList(jsn []byte) (ActivitiesList, error) {
	list := ActivitiesList{}
	if err := json.Unmarshal(jsn, &list); err != nil {
		return list, err
	}

	return list, nil
}

// ActivitiesBytes returns a reply to a link as a raw bytes.
func (f *Fitbit) ActivitiesBytes(link string) ([]byte, error) {
	return f.makeGETRequest(link)
}
