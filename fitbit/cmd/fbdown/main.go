package main

import (
	"archive/zip"
	"encoding/json"
	"fitbit/fitbit"
	"fitbit/fitbit/activities"
	"fitbit/fitbit/scopes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const configFileName = "fbdown.json"

func main() {
	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("Cannot get configuration: %s", err))
	}

	fb, err := fitbit.New(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURL, []string{
		scopes.Activity,
		scopes.Settings,
		scopes.Location,
		scopes.Social,
		scopes.Heartrate,
		scopes.Profile,
		scopes.Sleep,
		scopes.Nutrition,
		scopes.Weight,
	})
	if err != nil {
		panic(fmt.Sprintf("Cannot create fitbit session: %s", err))
	}

	if err = lookback(cfg, fb, archive); err != nil {
		panic(fmt.Sprintf("Cannot lookback: %s", err))
	}
}

type config struct {
	// ClientID is the Fitbit API application id from the settings on dev.fitbit.com.
	ClientID string `json:"clientID"`

	// ClientSecret is the Fitbit API application secret the your settings on dev.fitbit.com.
	ClientSecret string `json:"clientSecret"`

	// RedirectURL is the Fitbit API redirect URI from the settings on dev.fitbit.com.
	// When a user grants or denies access from the authorization screen, they are redirected
	// back to your application with information necessary to complete the authorization.
	RedirectURL string `json:"redirectURL"`

	// ArchiveDir is a peth to a directory containing archived daily zip files.
	// Should have a training directory delimiter.
	ArchiveDir string `json:"archiveDir"`

	// LookbackDate is an inclusive lookback date to get the data.
	// The format is '2021-09-27'.
	LookbackDate string `json:"lookbackDate"`
}

func readConfig(fileName string) (*config, error) {
	var conf config

	f, err := os.Open(fileName)
	if err != nil {
		return &conf, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&conf)
	if err != nil {
		return &conf, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	if !strings.HasSuffix(conf.ArchiveDir, "/") {
		conf.ArchiveDir += "/"
	}

	return &conf, nil
}

func lookback(conf *config, fb *fitbit.Fitbit, f func(*fitbit.Fitbit, string, string) error) error {
	const (
		layout       = "2006-01-02"
		minFileSize  = int64(8)
		rateLimitFmt = "rate limit remaining: %d of %d, reset time '%v'\n"
	)

	if _, err := os.Stat(conf.ArchiveDir); os.IsNotExist(err) {
		if err = os.MkdirAll(conf.ArchiveDir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create archive directory '%s': %w", conf.ArchiveDir, err)
		}
	}

	date, err := time.Parse(layout, conf.LookbackDate)
	if err != nil {
		return fmt.Errorf("cannot parse lookback date '%s': %w", conf.LookbackDate, err)
	}

	today := time.Now().AddDate(0, 0, -1)
	for date.Before(today) {
		d := date.Format(layout)
		path := conf.ArchiveDir + d + ".zip"

		if fi, err := os.Stat(path); os.IsNotExist(err) || fi.Size() < minFileSize {
			if err = f(fb, d, path); err != nil {
				return fmt.Errorf("cannot archive date '%s': %w", d, err)
			}

			rl := fb.Ratelimit()
			fmt.Printf(rateLimitFmt, rl.Remaining, rl.Limit, rl.Reset)
		}

		date = date.AddDate(0, 0, 1)
	}

	return nil
}

func archive(fb *fitbit.Fitbit, date, path string) error {
	const failed = "failed"

	fmt.Printf("Retrieving '%s' ... ", date)

	data, err := collect(fb, date)
	if err != nil {
		fmt.Println(failed)
		return err
	}

	fmt.Print("done\nZipping ... ")

	z, err := os.Create(path)
	if err != nil {
		fmt.Println(failed)
		return fmt.Errorf("cannot create '%s': %w", path, err)
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()

	for _, d := range data {
		nam := date + "/" + d.file

		f, err := w.Create(nam)
		if err != nil {
			fmt.Println(failed)
			return fmt.Errorf("cannot create zip entry '%s': %w", nam, err)
		}

		_, err = f.Write(d.body)
		if err != nil {
			fmt.Println(failed)
			return fmt.Errorf("cannot write zip entry '%s': %w", nam, err)
		}
	}

	fmt.Println("done")

	return nil
}

type entry struct {
	file string
	body []byte
	msg  string
}

//nolint:gocyclo // reviewed
func collect(fb *fitbit.Fitbit, date string) ([]*entry, error) {
	const (
		cannotGet                                = "cannot get %s %s: %w"
		msgProfile                               = "profile"
		msgBadges                                = "badges"
		msg1secHeartRate                         = "1-sec heart rate"
		msg1minHeartRate                         = "1-min heart rate"
		msg1dayHeartRate                         = "1-day heart rate"
		msgSleepGoal                             = "sleep goal"
		msgSleepDay                              = "sleep day"
		msgDevices                               = "devices"
		msgActivitiesSummary                     = "activities summary"
		msgActivitiesCalories                    = "activities calories"
		msgActivitiesCaloriesBMR                 = "activities calories bmr"
		msgActivitiesSteps                       = "activities steps"
		msgActivitiesDistance                    = "activities distance"
		msgActivitiesFloors                      = "activities floors"
		msgActivitiesElevation                   = "activities elevation"
		msgActivitiesMinutesSedentary            = "activities minutes sedentary"
		msgActivitiesMinutesLightlyActive        = "activities minutes lightly active"
		msgActivitiesMinutesFairlyActive         = "activities minutes fairly active"
		msgActivitiesMinutesVeryActive           = "activities minutes very active"
		msgActivitiesActivityCalories            = "activities activity calories"
		msgActivitiesTrackerCalories             = "activities tracker calories"
		msgActivitiesTrackerSteps                = "activities tracker steps"
		msgActivitiesTrackerDistance             = "activities tracker distance"
		msgActivitiesTrackerFloors               = "activities tracker floors"
		msgActivitiesTrackerElevation            = "activities tracker elevation"
		msgActivitiesTrackerMinutesSedentary     = "activities tracker minutes sedentary"
		msgActivitiesTrackerMinutesLightlyActive = "activities tracker minutes lightly active"
		msgActivitiesTrackerMinutesFairlyActive  = "activities tracker minutes fairly active"
		msgActivitiesTrackerMinutesVeryActive    = "activities tracker minutes very active"
		msgActivitiesTrackerActivityCalories     = "activities tracker activity calories"
		msgActivitiesIntraCalories1Min           = "1min activities intra calories"
		msgActivitiesIntraSteps1Min              = "1min activities intra steps"
		msgActivitiesIntraDistance1Min           = "1min activities intra distance"
		msgActivitiesIntraFloors1Min             = "1min activities intra floors"
		msgActivitiesIntraElevation1Min          = "1min activities intra elevation"
		msgActivitiesIntraCalories15Min          = "15min activities intra calories"
		msgActivitiesIntraSteps15Min             = "15min activities intra steps"
		msgActivitiesIntraDistance15Min          = "15min activities intra distance"
		msgActivitiesIntraFloors15Min            = "15min activities intra floors"
		msgActivitiesIntraElevation15Min         = "15min activities intra elevation"
		msgActivitiesTypes                       = "activities types"
		msgActivitiesLifetimeStats               = "activities lifetime stats"
		msgActivitiesGoalsDaily                  = "activities goals daily"
		msgActivitiesGoalsWeekly                 = "activities goals weekly"
		msgActivitiesList                        = "activities list"
	)

	var b []byte

	var err error

	entries := []*entry{
		{"profile.json", nil, msgProfile},
		{"badges.json", nil, msgBadges},
		{"heartrate/1sec.json", nil, msg1secHeartRate},
		{"heartrate/1min.json", nil, msg1minHeartRate},
		{"heartrate/1day.json", nil, msg1dayHeartRate},
		{"sleep/goal.json", nil, msgSleepGoal},
		{"sleep/day.json", nil, msgSleepDay},
		{"devices.json", nil, msgDevices},
		{"activities/summary.json", nil, msgActivitiesSummary},
		{"activities/calories.json", nil, msgActivitiesCalories},
		{"activities/caloriesBMR.json", nil, msgActivitiesCaloriesBMR},
		{"activities/steps.json", nil, msgActivitiesSteps},
		{"activities/distance.json", nil, msgActivitiesDistance},
		{"activities/floors.json", nil, msgActivitiesFloors},
		{"activities/elevation.json", nil, msgActivitiesElevation},
		{"activities/minutesSedentary.json", nil, msgActivitiesMinutesSedentary},
		{"activities/minutesLightlyActive.json", nil, msgActivitiesMinutesLightlyActive},
		{"activities/minutesFairlyActive.json", nil, msgActivitiesMinutesFairlyActive},
		{"activities/minutesVeryActive.json", nil, msgActivitiesMinutesVeryActive},
		{"activities/activityCalories.json", nil, msgActivitiesActivityCalories},
		{"activities/tracker/calories.json", nil, msgActivitiesTrackerCalories},
		{"activities/tracker/steps.json", nil, msgActivitiesTrackerSteps},
		{"activities/tracker/distance.json", nil, msgActivitiesTrackerDistance},
		{"activities/tracker/floors.json", nil, msgActivitiesTrackerFloors},
		{"activities/tracker/elevation.json", nil, msgActivitiesTrackerElevation},
		{"activities/tracker/minutesSedentary.json", nil, msgActivitiesTrackerMinutesSedentary},
		{"activities/tracker/minutesLightlyActive.json", nil, msgActivitiesTrackerMinutesLightlyActive},
		{"activities/tracker/minutesFairlyActive.json", nil, msgActivitiesTrackerMinutesFairlyActive},
		{"activities/tracker/minutesVeryActive.json", nil, msgActivitiesTrackerMinutesVeryActive},
		{"activities/tracker/activityCalories.json", nil, msgActivitiesTrackerActivityCalories},
		{"activities/1min/calories.json", nil, msgActivitiesIntraCalories1Min},
		{"activities/1min/steps.json", nil, msgActivitiesIntraSteps1Min},
		{"activities/1min/distance.json", nil, msgActivitiesIntraDistance1Min},
		{"activities/1min/floors.json", nil, msgActivitiesIntraFloors1Min},
		{"activities/1min/elevation.json", nil, msgActivitiesIntraElevation1Min},
		{"activities/15min/calories.json", nil, msgActivitiesIntraCalories15Min},
		{"activities/15min/steps.json", nil, msgActivitiesIntraSteps15Min},
		{"activities/15min/distance.json", nil, msgActivitiesIntraDistance15Min},
		{"activities/15min/floors.json", nil, msgActivitiesIntraFloors15Min},
		{"activities/15min/elevation.json", nil, msgActivitiesIntraElevation15Min},
		{"activities/types.json", nil, msgActivitiesTypes},
		{"activities/lifetimeStats.json", nil, msgActivitiesLifetimeStats},
		{"activities/goals-daily.json", nil, msgActivitiesGoalsDaily},
		{"activities/goals-weekly.json", nil, msgActivitiesGoalsWeekly},
		{"activities/list.json", nil, msgActivitiesList},
	}

	for _, e := range entries {
		switch e.msg {
		case msgProfile:
			b, err = fb.ProfileJSON(0)
		case msgBadges:
			b, err = fb.BadgesJSON(0)
		case msg1secHeartRate:
			b, err = fb.HeartIntradayJSON(date, "1sec")
		case msg1minHeartRate:
			b, err = fb.HeartIntradayJSON(date, "1min")
		case msg1dayHeartRate:
			b, err = fb.HeartDayJSON(date)
		case msgSleepGoal:
			b, err = fb.SleepGoalJSON()
		case msgSleepDay:
			b, err = fb.SleepDayJSON(date)
		case msgDevices:
			b, err = fb.DevicesJSON(0)
		case msgActivitiesSummary:
			b, err = fb.ActivitiesDaySummaryJSON(date)
		case msgActivitiesCalories:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyCalories)
		case msgActivitiesCaloriesBMR:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyCaloriesBMR)
		case msgActivitiesSteps:
			b, err = fb.ActivitiesDayJSON(date, activities.DailySteps)
		case msgActivitiesDistance:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyDistance)
		case msgActivitiesFloors:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyFloors)
		case msgActivitiesElevation:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyElevation)
		case msgActivitiesMinutesSedentary:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyMinutesSedentary)
		case msgActivitiesMinutesLightlyActive:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyMinutesLightlyActive)
		case msgActivitiesMinutesFairlyActive:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyMinutesFairlyActive)
		case msgActivitiesMinutesVeryActive:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyMinutesVeryActive)
		case msgActivitiesActivityCalories:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyActivityCalories)
		case msgActivitiesTrackerCalories:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerCalories)
		case msgActivitiesTrackerSteps:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerSteps)
		case msgActivitiesTrackerDistance:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerDistance)
		case msgActivitiesTrackerFloors:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerFloors)
		case msgActivitiesTrackerElevation:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerElevation)
		case msgActivitiesTrackerMinutesSedentary:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerMinutesSedentary)
		case msgActivitiesTrackerMinutesLightlyActive:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerMinutesLightlyActive)
		case msgActivitiesTrackerMinutesFairlyActive:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerMinutesFairlyActive)
		case msgActivitiesTrackerMinutesVeryActive:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerMinutesVeryActive)
		case msgActivitiesTrackerActivityCalories:
			b, err = fb.ActivitiesDayJSON(date, activities.DailyTrackerActivityCalories)
		case msgActivitiesIntraCalories1Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraCalories, "1min")
		case msgActivitiesIntraSteps1Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraSteps, "1min")
		case msgActivitiesIntraDistance1Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraDistance, "1min")
		case msgActivitiesIntraFloors1Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraFloors, "1min")
		case msgActivitiesIntraElevation1Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraElevation, "1min")
		case msgActivitiesIntraCalories15Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraCalories, "15min")
		case msgActivitiesIntraSteps15Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraSteps, "15min")
		case msgActivitiesIntraDistance15Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraDistance, "15min")
		case msgActivitiesIntraFloors15Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraFloors, "15min")
		case msgActivitiesIntraElevation15Min:
			b, err = fb.ActivitiesIntradayJSON(date, activities.IntraElevation, "15min")
		case msgActivitiesTypes:
			b, err = fb.ActivitiesTypesJSON()
		case msgActivitiesLifetimeStats:
			b, err = fb.ActivitiesLifetimeStatsJSON()
		case msgActivitiesGoalsDaily:
			b, err = fb.ActivitiesGoalJSON("daily")
		case msgActivitiesGoalsWeekly:
			b, err = fb.ActivitiesGoalJSON("weekly")
		case msgActivitiesList:
			b, err = fb.ActivitiesListJSON(date)
		}

		e.body = b
		if err != nil {
			return entries, fmt.Errorf(cannotGet, e.msg, date, err)
		}
	}

	list, err := fb.ActivitiesList(b)
	if err != nil {
		return entries, fmt.Errorf("cannot convert activities list '%s': %w", date, err)
	}

	const (
		maxLinkLen = 24
		base10     = 10
		layout     = "2006-01-02"
	)

	for i := range list.Activities {
		d := list.Activities[i].OriginalStartTime.Format(layout)
		if d != date {
			continue
		}

		id := strconv.FormatInt(list.Activities[i].LogID, base10)

		link := list.Activities[i].TcxLink
		if len(link) > maxLinkLen {
			b, err = fb.ActivitiesBytes(link)
			if err != nil {
				return entries, fmt.Errorf("cannot get activities tsx '%s': %w", link, err)
			}

			ent := &entry{"activities/list-" + id + ".tcx", b, ""}
			entries = append(entries, ent)
		}

		link = list.Activities[i].HeartRateLink
		if len(link) > maxLinkLen {
			b, err = fb.ActivitiesBytes(link)
			if err != nil {
				return entries, fmt.Errorf("cannot get activities heart rate '%s': %w", link, err)
			}

			ent := &entry{"activities/list-" + id + "-heartrate.json", b, ""}
			entries = append(entries, ent)
		}

		link = list.Activities[i].CaloriesLink
		if len(link) > maxLinkLen {
			b, err = fb.ActivitiesBytes(link)
			if err != nil {
				return entries, fmt.Errorf("cannot get activities calories '%s': %w", link, err)
			}

			ent := &entry{"activities/list-" + id + "-calories.json", b, ""}
			entries = append(entries, ent)
		}
	}

	return entries, nil
}
