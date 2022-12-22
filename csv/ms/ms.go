package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type MsDate struct {
	time.Time
}

func (t *MsDate) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}

	t.Time = date
	return nil
}

type MsTimeSeries struct {
	TimeSeries struct {
		Security []struct {
			HistoryDetail []struct {
				EndDate MsDate  `json:"EndDate"`
				Value   float64 `json:"Value"`
			} `json:"HistoryDetail"`
			Id string `json:"Id"`
		} `json:"Security"`
	} `json:"TimeSeries"`
}

func unmarshalMsTimeSeries(jsn []byte) (MsTimeSeries, error) {
	msts := MsTimeSeries{}
	if err := json.Unmarshal(jsn, &msts); err != nil {
		return msts, err
	}

	return msts, nil
}

// makeGETRequest makes a new GET request to a given URL using the given HTTP client.
func makeGETRequest(httpClient *http.Client, targetURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "fitbit")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
