package ms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type MsClient struct{}

type MsDate struct {
	time.Time
}

func (t *MsDate) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse("2006-01-02", string(b[1:len(b)-1]))
	if err != nil {
		return fmt.Errorf("cannot unmarshal date '%v': %w", string(b), err)
	}

	t.Time = date
	return nil
}

type MsValue float64

func (t *MsValue) UnmarshalJSON(b []byte) (err error) {
	v, err := strconv.ParseFloat(string(b[1:len(b)-1]), 64)
	if err != nil {
		return fmt.Errorf("cannot unmarshal date '%v': %w", string(b), err)
	}

	*t = MsValue(v)
	return nil
}

type MsTimeSeries struct {
	TimeSeries struct {
		Security []struct {
			HistoryDetail []struct {
				EndDate MsDate  `json:"EndDate"`
				Value   MsValue `json:"Value"`
			} `json:"HistoryDetail"`
			Id string `json:"Id"`
		} `json:"Security"`
	} `json:"TimeSeries"`
}

func unmarshal(jsn []byte) (*MsTimeSeries, error) {
	msts := MsTimeSeries{}
	if err := json.Unmarshal(jsn, &msts); err != nil {
		return nil, fmt.Errorf("cannot unmarshal time series: %w", err)
	}

	return &msts, nil
}

func get(targetURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", "ms")

	httpClient := http.Client{Timeout: time.Duration(300) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot do request: %w", err)
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	return contents, nil
}

func Download(msID, currency string, startDate time.Time) (*MsTimeSeries, error) {

	date := startDate.Format("2006-01-02")

	url := "http://lt.morningstar.com/api/rest.svc/timeseries_price/hvqzxf7smz?id=" + msID +
		"&idtype=MSID&startDate=" + date +
		"&Currencyid=" + currency + "&outputtype=json"
	fmt.Println(url)

	if bs, err := get(url); err != nil {
		return nil, fmt.Errorf("cannot retrieve: %w", err)
	} else {
		if msts, err := unmarshal(bs); err != nil {
			return nil, fmt.Errorf("cannot retrieve: %w", err)
		} else {
			return msts, nil
		}
	}
}
