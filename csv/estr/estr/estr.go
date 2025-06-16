// Fetches the European Short-Term Rate (€STR) historical data from the ECB website.
package estr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Landing page.
// https://www.ecb.europa.eu/stats/financial_markets_and_interest_rates/euro_short-term_rate/html/index.en.html

// Actual data page.
// https://data.ecb.europa.eu/data/data-categories/financial-markets-and-interest-rates/euro-money-market/euro-short-term-rate
//
// Actual GET URL: Volume-weighted trimmed mean rate, percent.
// https://data.ecb.europa.eu/data-detail-api/EST.B.EU000A2X2A25.WT
// Referer.
// https://data.ecb.europa.eu/data/datasets/EST/EST.B.EU000A2X2A25.WT
//
// Actual GET URL: Total volume, millions of Euro.
// https://data.ecb.europa.eu/data-detail-api/EST.B.EU000A2X2A25.TT
// Referer.
// https://data.ecb.europa.eu/data/datasets/EST/EST.B.EU000A2X2A25.TT
//
// Actual GET URL: Number of transactions.
// https://data.ecb.europa.eu/data-detail-api/EST.B.EU000A2X2A25.NT
// Referer.
// https://data.ecb.europa.eu/data/datasets/EST/EST.B.EU000A2X2A25.NT

// Pre-data page.
// https://data.ecb.europa.eu/data/data-categories/financial-markets-and-interest-rates/euro-money-market/pre-euro-short-term-rate
//
// Pre-data GET URL: Volume-weighted trimmed mean rate, percent.
// https://data.ecb.europa.eu/data-detail-api/MMSR.B.U2._X._Z.S12._Z.U.BO.WT.D76.MA._Z._Z.EUR._Z
// Referer.
// https://data.ecb.europa.eu/data/datasets/MMSR/MMSR.B.U2._X._Z.S12._Z.U.BO.WT.D76.MA._Z._Z.EUR._Z
//
// Pre-data GET URL: Total volume, millions of Euro.
// https://data.ecb.europa.eu/data-detail-api/MMSR.B.U2._X._Z.S12._Z.U.BO.TT.D76.MA._Z._Z.EUR._Z
// Referer.
// https://data.ecb.europa.eu/data/datasets/MMSR/MMSR.B.U2._X._Z.S12._Z.U.BO.TT.D76.MA._Z._Z.EUR._Z
//
// Pre-data GET URL: Number of transactions.
// https://data.ecb.europa.eu/data/datasets/MMSR/MMSR.B.U2._X._Z.S12._Z.U.BO.NT.D76.MA._Z._Z.EUR._Z
// Referer.
// https://data.ecb.europa.eu/data-detail-api/MMSR.B.U2._X._Z.S12._Z.U.BO.NT.D76.MA._Z._Z.EUR._Z

// Response JSON (note rate is null on weekends).
/*
[{
// Workday, "OBS":"2.172", "PERIOD":"2025-06-02"
"OBS":"2.172","SERIES":"EST.B.EU000A2X2A25.WT","OBS_VALUE_ENTITY":"2.172","UNIT":"","PERIOD_ID":"445808",
"OBS_POINT":"lastDate","OBS_COM":null,"TREND_INDICATOR":"up","PERIOD_NAME":"02 Jun 2025",
"LEGEND":"Normal value","OBS_STATUS":"A","OBS_VALUE_AS_IS":"2.172","PERIOD":"2025-06-02",
"FREQUENCY":"B","OBS_CONF":"F","PERIOD_DATA_COMP":"2025-06-02",
"VALID_FROM":"2025-06-03T06:04:00.000+00:00","OBS_PRE_BREAK":null
},{
// Weekend, no data, "OBS":null, "PERIOD":"2025-06-01"
"PERIOD_NAME":"01 Jun 2025","OBS_STATUS":null,"OBS":null,"SERIES":"EST.B.EU000A2X2A25.WT",
"UNIT":"","OBS_VALUE_ENTITY":"","PERIOD":"2025-06-01","FREQUENCY":"B","OBS_POINT":"lastDate",
"OBS_COM":null
},{
// Weekend, no data, "OBS":null, "PERIOD":"2025-05-31"
"PERIOD_NAME":"31 May 2025","OBS_STATUS":null,"OBS":null,"SERIES":"EST.B.EU000A2X2A25.WT","UNIT":"",
"OBS_VALUE_ENTITY":"","PERIOD":"2025-05-31","FREQUENCY":"B","OBS_POINT":"lastDate","OBS_COM":null
},{
// Workday, "OBS":"2.161", "PERIOD":"2025-05-30"
"OBS":"2.161","SERIES":"EST.B.EU000A2X2A25.WT","OBS_VALUE_ENTITY":"2.161","UNIT":"","PERIOD_ID":"445805",
"OBS_POINT":"lastDate","OBS_COM":null,"TREND_INDICATOR":"equal","PERIOD_NAME":"30 May 2025",
"LEGEND":"Normal value","OBS_STATUS":"A","OBS_VALUE_AS_IS":"2.161","PERIOD":"2025-05-30","FREQUENCY":"B",
"OBS_CONF":"F","PERIOD_DATA_COMP":"2025-05-30","VALID_FROM":"2025-06-02T06:04:00.000+00:00",
"OBS_PRE_BREAK":null
}]
*/

// The response structure for the ECB €STR time series data.
/*
type seriesResponse []struct {
	Obs            string    `json:"OBS"`
	Series         string    `json:"SERIES"`
	ObsValueEntity string    `json:"OBS_VALUE_ENTITY"`
	Unit           string    `json:"UNIT"`
	PeriodID       string    `json:"PERIOD_ID"`
	ObsPoint       string    `json:"OBS_POINT"`
	ObsCom         any       `json:"OBS_COM"`
	TrendIndicator string    `json:"TREND_INDICATOR"`
	PeriodName     string    `json:"PERIOD_NAME"`
	Legend         string    `json:"LEGEND"`
	ObsStatus      string    `json:"OBS_STATUS"`
	ObsValueAsIs   string    `json:"OBS_VALUE_AS_IS"`
	Period         string    `json:"PERIOD"`
	Frequency      string    `json:"FREQUENCY"`
	ObsConf        string    `json:"OBS_CONF"`
	PeriodDataComp string    `json:"PERIOD_DATA_COMP"`
	ValidFrom      time.Time `json:"VALID_FROM"`
	ObsPreBreak    any       `json:"OBS_PRE_BREAK"`
}
*/

type What int

const (
	EstrRateAct What = iota
	EstrVolumeAct
	EstrTransactionsAct
	EstrRatePre
	EstrVolumePre
	EstrTransactionsPre
)

const (
	rateAct      = "https://data.ecb.europa.eu/data-detail-api/EST.B.EU000A2X2A25.WT"
	rateActRef   = "https://data.ecb.europa.eu/data/datasets/EST/EST.B.EU000A2X2A25.WT"
	volumeAct    = "https://data.ecb.europa.eu/data-detail-api/EST.B.EU000A2X2A25.TT"
	volumeActRef = "https://data.ecb.europa.eu/data/datasets/EST/EST.B.EU000A2X2A25.TT"
	transAct     = "https://data.ecb.europa.eu/data-detail-api/EST.B.EU000A2X2A25.NT"
	transActRef  = "https://data.ecb.europa.eu/data/datasets/EST/EST.B.EU000A2X2A25.NT"

	ratePre      = "https://data.ecb.europa.eu/data-detail-api/MMSR.B.U2._X._Z.S12._Z.U.BO.WT.D76.MA._Z._Z.EUR._Z"
	ratePreRef   = "https://data.ecb.europa.eu/data/datasets/MMSR/MMSR.B.U2._X._Z.S12._Z.U.BO.WT.D76.MA._Z._Z.EUR._Z"
	volumePre    = "https://data.ecb.europa.eu/data-detail-api/MMSR.B.U2._X._Z.S12._Z.U.BO.TT.D76.MA._Z._Z.EUR._Z"
	volumePreRef = "https://data.ecb.europa.eu/data/datasets/MMSR/MMSR.B.U2._X._Z.S12._Z.U.BO.TT.D76.MA._Z._Z.EUR._Z"
	transPre     = "https://data.ecb.europa.eu/data-detail-api/MMSR.B.U2._X._Z.S12._Z.U.BO.NT.D76.MA._Z._Z.EUR._Z"
	transPreRef  = "https://data.ecb.europa.eu/data/datasets/MMSR/MMSR.B.U2._X._Z.S12._Z.U.BO.NT.D76.MA._Z._Z.EUR._Z"

	userAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	timeoutSec = 60
	retriesMax = 5
	sleepSec   = 5
)

var gets = [6]string{rateAct, volumeAct, transAct, ratePre, volumePre, transPre}
var refs = [6]string{rateActRef, volumeActRef, transActRef, ratePreRef, volumePreRef, transPreRef}

type estrSeries []struct {
	Date  string `json:"PERIOD"`
	Value string `json:"OBS"`
}

func fetch(what What) (*estrSeries, error) {
	req, err := http.NewRequest("GET", gets[what], nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", refs[what])
	httpClient := http.Client{Timeout: time.Duration(timeoutSec) * time.Second}

	var resp *http.Response
	retries := retriesMax
	for retries > 0 {
		resp, err = httpClient.Do(req)
		if err != nil {
			retries -= 1
			if retries < 1 {
				err := fmt.Errorf("cannot do request after retries %d: %w", retriesMax, err)
				return nil, err
			} else {
				time.Sleep(sleepSec * time.Second)
			}
		} else {
			break
		}
	}
	defer resp.Body.Close()

	jsn, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	ct := estrSeries{}
	if err = json.Unmarshal(jsn, &ct); err != nil {
		return nil, fmt.Errorf("cannot unmarshal content: %w", err)
	}

	return &ct, nil
}

type Point struct {
	Date  time.Time
	Value float64
}

func Fetch(what What) ([]Point, error) {
	series, err := fetch(what)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch data: %w", err)
	}

	var result []Point

	// Iterate in reverse order since last date is the first in JSON response
	for i := len(*series) - 1; i >= 0; i-- {
		s := (*series)[i]
		if s.Value == "" {
			continue // Skip empty values
		}

		date, err := time.Parse("2006-01-02", s.Date)
		if err != nil {
			return nil, fmt.Errorf("cannot parse date '%s': %w", s.Date, err)
		}

		value, err := strconv.ParseFloat(s.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse value '%s': %w", s.Value, err)
		}

		result = append(result, Point{Date: date, Value: value})
	}

	return result, nil
}
