// Fetches the European Short-Term Rate (€STR) historical data from the ECB website.
package silso

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Landing page.
// https://www.sidc.be/SILSO/home

// Actual data page.
// https://www.sidc.be/SILSO/datafiles
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
)

var gets = [6]string{rateAct, volumeAct, transAct, ratePre, volumePre, transPre}
var refs = [6]string{rateActRef, volumeActRef, transActRef, ratePreRef, volumePreRef, transPreRef}
var labs = [6]string{"rate act", "volume act", "trans act", "rate pre", "volume pre", "trans pre"}
var nams = [6]string{"EST.B.EU000A2X2A25.WT", "EST.B.EU000A2X2A25.TT", "EST.B.EU000A2X2A25.NT",
	"MMSR.B.U2._X._Z.S12._Z.U.BO.WT.D76.MA._Z._Z.EUR._Z",
	"MMSR.B.U2._X._Z.S12._Z.U.BO.TT.D76.MA._Z._Z.EUR._Z",
	"MMSR.B.U2._X._Z.S12._Z.U.BO.NT.D76.MA._Z._Z.EUR._Z"}

type estrSeries []struct {
	Date  string `json:"PERIOD"`
	Value string `json:"OBS"`
}

func get(
	uri string,
	timeout time.Duration,
	referer string,
	userAgent string,
	verbose bool,
) ([]byte, error) {
	if verbose {
		log.Println(uri)
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", referer)
	req.Header.Set("Accept", "application/json, text/javascript, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")

	// Create HTTP client with timeout and proxy settings
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment, // Uses system proxy settings
	}
	httpClient := http.Client{Timeout: timeout, Transport: transport}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("download failed %s: %w", uri, err)
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body %s: %w", uri, err)
	}

	return contents, nil
}

func getWithRetries(
	uri string,
	label string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	referer string,
	userAgent string,
	verbose bool,
) ([]byte, error) {
	var contents []byte
	var err error
	retriesMax := len(pauseBeforeRetry)
	retries := retriesMax
	for retries > 0 {
		contents, err = get(uri, timeout, referer, userAgent, verbose)
		if err != nil {
			if retries > 1 {
				log.Printf("%s: download failed, retrying (%d of %d left): %v\n", label, retries, retriesMax, err)
			} else {
				return nil, fmt.Errorf("%s: download failed, giving up (%d of %d left): %v", label, retries, retriesMax, err)
			}
			retries--
			continue
		}
		break
	}

	return contents, nil
}

func fetch(
	what What,
	writeToFile bool,
	downloadFolder string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	userAgent string,
	verbose bool,
) (*estrSeries, error) {
	if bs, err := getWithRetries(gets[what], labs[what], timeout, pauseBeforeRetry, refs[what],
		userAgent, verbose); err != nil {
		return nil, err
	} else {
		if writeToFile {
			if _, err := os.Stat(downloadFolder); os.IsNotExist(err) {
				if err = os.MkdirAll(downloadFolder, os.ModePerm); err != nil {
					log.Printf("cannot create directory '%s': %s\n", downloadFolder, err)
				}
			}
			file := filepath.Join(downloadFolder,
				nams[what]+".json")
			if err := os.WriteFile(file, bs, 0644); err != nil {
				log.Printf("cannot write to file %s: %s\n", file, err)
			}
		}

		var series estrSeries
		if err := json.Unmarshal(bs, &series); err != nil {
			return nil, fmt.Errorf("cannot unmarshal response: %w", err)
		}
		return &series, nil
	}
}

type Point struct {
	Date  time.Time
	Value float64
}

func Fetch(
	what What,
	writeToFile bool,
	downloadFolder string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	userAgent string,
	verbose bool,
) ([]Point, error) {
	series, err := fetch(what, writeToFile, downloadFolder, timeout, pauseBeforeRetry, userAgent, verbose)
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
