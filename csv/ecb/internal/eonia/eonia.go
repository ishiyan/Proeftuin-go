// Fetches the Internal Eonia Rate (EON) historical data from the ECB website.
package eonia

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
// https://data.ecb.europa.eu/
// https://data.ecb.europa.eu/data/datasets#alternative-datatable
/*
The EONIA rate was the closing rate for the overnight maturity calculated by collecting data on unsecured
overnight lendingin the euro area provided by banks belonging to the EONIA panel.

Following a recommendation made by the working group on euro risk-free rates on 14 March 2019, as of 2 October
for the trade date 1 October 2019 the European Money Market Institute (EMMI) changed the way it calculates the EONIA .
The EONIA methodology has been redefined as the euro short-term rate (€STR) plus a fixed spread, calculated
using the methodology adopted by the EMMI as the difference between the underlying interest rate of the EONIA
and the pre-€STR using daily data from 17 April 2018 to 16 April 2019. The ECB calculated this spread
as 0.085% (8.5 basis points). For this reason the volume information is not available anymore.
EMMI publishes EONIA for day T on T+1 at or shortly after 09.15 each TARGET2 business day.

The information is published in the ECB Statistical Data Warehouse on T+2 at, or shortly after, 09.15 each TARGET2 business day.
*/

// Actual data page.
// https://data.ecb.europa.eu/data/datasets/EON?dataset%5B0%5D=EONIA-%20Euro%20Interbank%20Offered%20Rate%20%28discontinued%29%20%28EON%29&advFilterDataset%5B0%5D=EONIA-%20Euro%20Interbank%20Offered%20Rate%20%28discontinued%29%20%28EON%29 
//
// Actual GET URL: Rate for the overnight maturity calculated as the euro short-term rate plus a spread of 8.5 basis points, Daily. Discontinued.
// 04 Jan 1999 to 31 Dec 2021
// Last updated: 4 Jan 2022 09:15 CET
// OBS is rate in percentage.
// https://data.ecb.europa.eu/data-detail-api/EON.D.EONIA_TO.RATE
// Referer.
// https://data.ecb.europa.eu/data/datasets/EON/EON.D.EONIA_TO.RATE
//
// Actual GET URL: Calculated by collecting data on unsecured overnight lending in the euro area - bank EONIA Total/aggregate,
// Volume for the overnight maturity [Discontinued], Daily.
// 04 Jan 1999 to 30 Sep 2019
// Last updated: 1 Oct 2019 19:01 CEST
// OBS is rate in total aggregated volume.
// https://data.ecb.europa.eu/data-detail-api/EON.D.EONIA_TO.VOLUME
// Referer.
// https://data.ecb.europa.eu/data/datasets/EON/EON.D.EONIA_TO.VOLUME

// Response JSON (note rate is null on weekends).
/*

[
// Workday, "OBS":"-0.505", "PERIOD":"2021-12-31"
"OBS":"-0.505","SERIES":"EON.D.EONIA_TO.RATE","OBS_VALUE_ENTITY":"-0.505","UNIT":"PC","PERIOD_ID":"113513",
"OBS_POINT":"lastDate","OBS_COM":null,"TREND_INDICATOR":"down","PERIOD_NAME":"31 Dec 2021",
"LEGEND":"Normal value","OBS_STATUS":"A","OBS_VALUE_AS_IS":"-0.505","PERIOD":"2021-12-31",
"FREQUENCY":"D","OBS_CONF":null,"PERIOD_DATA_COMP":"2021-12-31",
"VALID_FROM":"2022-01-04T08:15:00.000+00:00","OBS_PRE_BREAK":null
},{
// Workday, "OBS":"-0.495", "PERIOD":"2021-12-31"
"OBS":"-0.495","SERIES":"EON.D.EONIA_TO.RATE","OBS_VALUE_ENTITY":"-0.495","UNIT":"PC","PERIOD_ID":"113512",
"OBS_POINT":"lastDate","OBS_COM":null,"TREND_INDICATOR":"down","PERIOD_NAME":"30 Dec 2021",
"LEGEND":"Normal value","OBS_STATUS":"A","OBS_VALUE_AS_IS":"-0.495","PERIOD":"2021-12-30",
"FREQUENCY":"D","OBS_CONF":null,"PERIOD_DATA_COMP":"2021-12-30",
"VALID_FROM":"2022-01-01T08:15:00.000+00:00","OBS_PRE_BREAK":null
},{
. . .
},{
// Weekend, no data, "OBS":null, "PERIOD":"26 Dec 2021"
"PERIOD_NAME":"26 Dec 2021","OBS_STATUS":null,"OBS":null,"SERIES":"EON.D.EONIA_TO.RATE","UNIT":"PC",
"OBS_VALUE_ENTITY":"","PERIOD":"2021-12-26","FREQUENCY":"D","OBS_POINT":"lastDate","OBS_COM":null
},{
. . .
},{
"OBS":"3.200","SERIES":"EON.D.EONIA_TO.RATE","OBS_VALUE_ENTITY":"3.200","UNIT":"PC","PERIOD_ID":"105117",
"OBS_POINT":"lastDate","OBS_COM":null,"TREND_INDICATOR":"equal","PERIOD_NAME":"05 Jan 1999",
"LEGEND":"Normal value","OBS_STATUS":"A","OBS_VALUE_AS_IS":"3.2","PERIOD":"1999-01-05",
"FREQUENCY":"D","OBS_CONF":null,"PERIOD_DATA_COMP":"1999-01-05",
"VALID_FROM":"2009-01-07T05:35:00.000+00:00","OBS_PRE_BREAK":null
},{
"OBS":"3.200","SERIES":"EON.D.EONIA_TO.RATE","OBS_VALUE_ENTITY":"3.200","UNIT":"PC","PERIOD_ID":"105116",
"OBS_POINT":"lastDate","OBS_COM":null,"TREND_INDICATOR":null,"PERIOD_NAME":"04 Jan 1999",
"LEGEND":"Normal value","OBS_STATUS":"A","OBS_VALUE_AS_IS":"3.2","PERIOD":"1999-01-04",
"FREQUENCY":"D","OBS_CONF":null,"PERIOD_DATA_COMP":"1999-01-04",
"VALID_FROM":"2009-01-07T05:35:00.000+00:00","OBS_PRE_BREAK":null
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
