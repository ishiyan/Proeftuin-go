package main

import (
	"archive/zip"
	"encoding/json"
	"fitbit/fitbit"
	"fitbit/fitbit/scopes"
	"fmt"
	"os"
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
		layout      = "2006-01-02"
		minFileSize = int64(8)
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
		}

		date = date.AddDate(0, 0, 1)
	}

	return nil
}

func archive(fb *fitbit.Fitbit, date, path string) error {
	const failed = "failed"

	fmt.Printf("Retrieving '%s' ...", date)

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
		nam := date + "/" + d.name

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
	name string
	body []byte
}

func collect(fb *fitbit.Fitbit, date string) ([]entry, error) {
	const cannotGet = "cannot get %s %s: %w"

	entries := []entry{
		{"profile.json", nil},
		{"badges.json", nil},
		{"hr-1sec.json", nil},    // heartrate/1sec.json
		{"hr-1min.json", nil},    // heartrate/1min.json
		{"hr-1day.json", nil},    // heartrate/1day.json
		{"sleep-goal.json", nil}, // sleep/goal.json
		{"sleep-day.json", nil},  // sleep/day.json
		{"devices.json", nil},
		{"activities-summary.json", nil}, // activities/summary.json
	}

	b, err := fb.ProfileJSON(0)
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "profile", date, err)
	}

	entries[0].body = b

	b, err = fb.BadgesJSON(0)
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "badges", date, err)
	}

	entries[1].body = b

	b, err = fb.HeartIntradayJSON(date, "1sec")
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "1-sec heart rate", date, err)
	}

	entries[2].body = b

	b, err = fb.HeartIntradayJSON(date, "1min")
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "1-min heart rate", date, err)
	}

	entries[3].body = b

	b, err = fb.HeartDayJSON(date)
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "1-day heart rate", date, err)
	}

	entries[4].body = b

	b, err = fb.SleepGoalJSON()
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "sleep goal", date, err)
	}

	entries[5].body = b

	b, err = fb.SleepDayJSON(date)
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "sleep day", date, err)
	}

	entries[6].body = b

	b, err = fb.DevicesJSON(0)
	if err != nil {
		return entries, fmt.Errorf(cannotGet, "devices", date, err)
	}

	entries[7].body = b

	return entries, nil
}
