package fitbit

import (
	"encoding/json"
)

// HeartDay contains a summary of heartrates for a given date range.
type HeartDay struct {
	ActivitiesHeart []struct {
		DateTime string `json:"dateTime"`
		Value    struct {
			CustomHeartRateZones []interface{}   `json:"customHeartRateZones,omitempty"`
			HeartRateZones       []HeartRateZone `json:"heartRateZones"`
			RestingHeartRate     int             `json:"restingHeartRate"`
		} `json:"value"`
	} `json:"activities-heart"`
	ActivitiesHeartIntraday ActivitiesHeartIntraday `json:"activities-heart-intraday,omitempty"`
}

// HeartIntraday with slightly different structure to HeartDay.
type HeartIntraday struct {
	ActivitiesHeart []struct {
		CustomHeartRateZones []interface{}   `json:"customHeartRateZones"`
		DateTime             string          `json:"dateTime"`
		HeartRateZones       []HeartRateZone `json:"heartRateZones"`
		Value                string          `json:"value"`
	} `json:"activities-heart"`
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

// ActivitiesHeartIntraday intraday data.
type ActivitiesHeartIntraday struct {
	Dataset []struct {
		Time  string `json:"time"`
		Value int    `json:"value"`
	} `json:"dataset"`
	DatasetInterval int    `json:"datasetInterval"`
	DatasetType     string `json:"datasetType"`
}

// HeartLogByDay returns the heart log by a given date.
// The date must be in the format "yyyy-MM-dd" or "today".
func (f *FitBit) HeartLogByDay(day string) (HeartDay, []byte, error) {
	// If not day is given assume today.
	if day == "" {
		day = "today"
	}

	contents, err := f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + day + "/1d.json")
	if err != nil {
		return HeartDay{}, contents, err
	}

	heart := HeartDay{}
	if err := json.Unmarshal(contents, &heart); err != nil {
		return HeartDay{}, contents, err
	}

	return heart, contents, nil
}

// HeartIntradayPeriod returns the heart log by a given date in the given resolution.
// Day must be in the format "yyyy-MM-dd" or "today".
// Resolution (detail-level) can be "1min" or "1sec".
func (f *FitBit) HeartIntraday(day, resolution string) (HeartIntraday, []byte, error) {
	// If not day is given assume today.
	if day == "" {
		day = "today"
	}

	// default to 1sec if resolution dos not match to 1min.
	if resolution != "1min" {
		resolution = "1sec"
	}

	contents, err := f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + day + "/1d/" + resolution + ".json")
	if err != nil {
		return HeartIntraday{}, contents, err
	}

	heartintra := HeartIntraday{}
	if err := json.Unmarshal(contents, &heartintra); err != nil {
		return HeartIntraday{}, contents, err
	}

	return heartintra, contents, nil
}

// HeartIntradayPeriod returns the heart log by a given date in the given resolution.
// Day must be in the format "yyyy-MM-dd" or "today".
// Resolution (detail-level) can be "1min" or "1sec".
// TimeFrom and timeTo are in the format "00:00" for hour:minute.
func (f *FitBit) HeartIntradayPeriod(day, resolution, timeFrom, timeTo string) (HeartIntraday, []byte, error) {
	// If not day is given assume today.
	if day == "" {
		day = "today"
	}

	if timeFrom == "" {
		timeFrom = "00:00"
	}

	if timeTo == "" {
		timeTo = "23:59"
	}

	// default to 1sec if resolution dos not match to 1min.
	if resolution != "1min" {
		resolution = "1sec"
	}

	contents, err := f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + day + "/1d/" + resolution + "/time/" + timeFrom + "/" + timeTo + ".json")
	if err != nil {
		return HeartIntraday{}, contents, err
	}

	heartintra := HeartIntraday{}
	if err := json.Unmarshal(contents, &heartintra); err != nil {
		return HeartIntraday{}, contents, err
	}

	return heartintra, contents, nil
}

// HeartLogByDateRange returns the calories log of a given time range by date
// date must be in the format yyyy-MM-dd.
func (f *FitBit) HeartLogByDateRange(startDay, endDay string) (HeartDay, []byte, error) {
	contents, err := f.makeGETRequest("https://api.fitbit.com/1/user/-/activities/heart/date/" + startDay + "/" + endDay + ".json")
	if err != nil {
		return HeartDay{}, contents, err
	}

	heart := HeartDay{}
	if err := json.Unmarshal(contents, &heart); err != nil {
		return HeartDay{}, contents, err
	}

	return heart, contents, nil
}
