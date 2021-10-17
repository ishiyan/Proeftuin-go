package fitbit

import (
	"encoding/json"
	"time"
)

// SleepDay contains data of a sleep day.
type SleepDay struct {
	Sleep []struct {
		DateOfSleep string `json:"dateOfSleep"`
		Duration    int    `json:"duration"`
		Efficiency  int    `json:"efficiency"`
		EndTime     string `json:"endTime"`
		InfoCode    int    `json:"infoCode"`
		IsMainSleep bool   `json:"isMainSleep"`
		Levels      struct {
			Data []struct {
				DateTime string `json:"dateTime"`
				Level    string `json:"level"`
				Seconds  int    `json:"seconds"`
			} `json:"data"`
			ShortData []struct {
				DateTime string `json:"dateTime"`
				Level    string `json:"level"`
				Seconds  int    `json:"seconds"`
			} `json:"shortData"`
			Summary struct {
				Deep struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
				} `json:"deep"`
				Light struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
				} `json:"light"`
				Rem struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
				} `json:"rem"`
				Wake struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes"`
				} `json:"wake"`
				Asleep struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes,omitempty"`
				} `json:"asleep,omitempty"`
				Awake struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes,omitempty"`
				} `json:"awake,omitempty"`
				Restless struct {
					Count               int `json:"count"`
					Minutes             int `json:"minutes"`
					ThirtyDayAvgMinutes int `json:"thirtyDayAvgMinutes,omitempty"`
				} `json:"restless,omitempty"`
			} `json:"summary"`
		} `json:"levels,omitempty"`
		LogID               int64  `json:"logId"`
		MinutesAfterWakeup  int    `json:"minutesAfterWakeup"`
		MinutesAsleep       int    `json:"minutesAsleep"`
		MinutesAwake        int    `json:"minutesAwake"`
		MinutesToFallAsleep int    `json:"minutesToFallAsleep"`
		StartTime           string `json:"startTime"`
		TimeInBed           int    `json:"timeInBed"`
		Type                string `json:"type"`
	} `json:"sleep"`
	Summary struct {
		Stages struct {
			Deep  int `json:"deep"`
			Light int `json:"light"`
			Rem   int `json:"rem"`
			Wake  int `json:"wake"`
		} `json:"stages"`
		TotalMinutesAsleep int `json:"totalMinutesAsleep"`
		TotalSleepRecords  int `json:"totalSleepRecords"`
		TotalTimeInBed     int `json:"totalTimeInBed"`
	} `json:"summary,omitempty"`
	Meta struct {
		RetryDuration int    `json:"retryDuration"`
		State         string `json:"state"`
	} `json:"meta,omitempty"`
}

// SleepGoal describes a sleep goal.
type SleepGoal struct {
	Consistency struct {
		AwakeRestlessPercentage float64 `json:"awakeRestlessPercentage"`
		FlowID                  int     `json:"flowId"`
		RecommendedSleepGoal    int     `json:"recommendedSleepGoal"`
		TypicalDuration         int     `json:"typicalDuration"`
		TypicalWakeupTime       string  `json:"typicalWakeupTime"`
	} `json:"consistency"`
	Goal struct {
		Bedtime     string    `json:"bedtime"`
		MinDuration int       `json:"minDuration"`
		UpdatedOn   time.Time `json:"updatedOn"`
		WakeupTime  string    `json:"wakeupTime"`
	} `json:"goal"`
}

// SleepDayJSON returns the sleep data for a given date as a raw JSON.
// Date must be in the format "yyyy-MM-dd".
func (f *Fitbit) SleepDayJSON(day string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1.2/user/-/sleep/date/" + day + ".json")
}

// SleepDay converts the raw JSON to the HeartDay type.
func (f *Fitbit) SleepDay(jsn []byte) (SleepDay, error) {
	sleep := SleepDay{}
	if err := json.Unmarshal(jsn, &sleep); err != nil {
		return sleep, err
	}

	return sleep, nil
}

// SleepGoalJSON returns the sleep goal data as a raw JSON.
func (f *Fitbit) SleepGoalJSON() ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1.2/user/-/sleep/goal.json")
}

// SleepGoal converts the raw JSON to the SleepGoal type.
func (f *Fitbit) SleepGoal(jsn []byte) (SleepGoal, error) {
	goal := SleepGoal{}
	if err := json.Unmarshal(jsn, &goal); err != nil {
		return goal, err
	}

	return goal, nil
}
