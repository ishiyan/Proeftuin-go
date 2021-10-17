package fitbit

import (
	"encoding/json"
)

// HeartDay contains a summary of heartrates for a given date range.
type HeartDay struct {
	ActivitiesHeart []ActivitiesHeart `json:"activities-heart"`
}

// HeartIntraday contains the Heart Rate sensor measures in 'Beats per minute'.
type HeartIntraday struct {
	ActivitiesHeart         []ActivitiesHeart       `json:"activities-heart,omitempty"`
	ActivitiesHeartIntraday ActivitiesHeartIntraday `json:"activities-heart-intraday,omitempty"`
}

// HeartRateZone contains the heart rate zones of different types like cardio.
type HeartRateZone struct {
	CaloriesOut float64 `json:"caloriesOut"`
	Max         int     `json:"max"`
	Min         int     `json:"min"`
	Minutes     int     `json:"minutes"`
	Name        string  `json:"name"`
}

// ActivitiesHeart dayly data.
type ActivitiesHeart struct {
	DateTime string `json:"dateTime"`
	Value    struct {
		CustomHeartRateZones []interface{}   `json:"customHeartRateZones,omitempty"`
		HeartRateZones       []HeartRateZone `json:"heartRateZones"`
		RestingHeartRate     int             `json:"restingHeartRate"`
	} `json:"value"`
}

// ActivitiesHeartIntraday intraday data.
type ActivitiesHeartIntraday struct {
	Dataset []struct {
		Time  string `json:"time"`
		Value int    `json:"value"`
	} `json:"dataset"`
	DatasetInterval int    `json:"datasetInterval"`
	DatasetType     string `json:"datasetType"`
}

// HeartDayJSON returns the heart log by a given date as a raw JSON.
// The date must be in the format "yyyy-MM-dd".
func (f *Fitbit) HeartDayJSON(day string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + day + "/1d.json")
}

// HeartLogPeriodJSON returns the heart log of a given time range by date as a raw JSON.
// Dates must be in the format "yyyy-MM-dd".
func (f *Fitbit) HeartLogPeriodJSON(startDay, endDay string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + startDay + "/" + endDay + ".json")
}

// HeartDay converts the raw JSON to the HeartDay type.
func (f *Fitbit) HeartDay(jsn []byte) (HeartDay, error) {
	heart := HeartDay{}
	if err := json.Unmarshal(jsn, &heart); err != nil {
		return HeartDay{}, err
	}

	return heart, nil
}

// HeartIntradayJSON returns the heart log by a given date in the given resolution as a raw JSON.
// Day must be in the format "yyyy-MM-dd".
// Resolution (detail-level) must be "1min" or "1sec".
func (f *Fitbit) HeartIntradayJSON(day, resolution string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + day + "/1d/" + resolution + ".json")
}

// HeartIntradayPeriodJSON returns the heart log for a given date, resolution and the time range as a raw JSON.
// Day must be in the format "yyyy-MM-dd".
// Resolution (detail-level) can be "1min" or "1sec".
// TimeFrom and timeTo are inclusive and are in the format "00:00" for hour:minute.
func (f *Fitbit) HeartIntradayPeriodJSON(day, resolution, timeFrom, timeTo string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + day + "/1d/" + resolution + "/time/" + timeFrom + "/" + timeTo + ".json")
}

// HeartIntraday converts the raw JSON to the HeartIntraday type.
func (f *Fitbit) HeartIntraday(jsn []byte) (HeartIntraday, error) {
	heartIntra := HeartIntraday{}
	if err := json.Unmarshal(jsn, &heartIntra); err != nil {
		return HeartIntraday{}, err
	}

	return heartIntra, nil
}
