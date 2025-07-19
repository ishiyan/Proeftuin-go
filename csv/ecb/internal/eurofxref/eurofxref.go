// Fetches the Euro foreign exchange reference rates historical data from the ECB website.
package eurofxref

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Landing page, also actual data page, serves as referrer.
// https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html
// Information page.
// https://data.ecb.europa.eu/data/data-categories/2771720/data-information

// Daily latest reference rates GET URL.
// https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml
// Referer: landing page
//
// Last 90 day reference rates history GET URL.
// http://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml
// Referer: landing page
//
// Full day reference rates history GET URL.
// http://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist.xml
// Referer: landing page

// Response XML.
/*
<?xml version="1.0" encoding="UTF-8"?>
<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
<gesmes:subject>Reference rates</gesmes:subject>
<gesmes:Sender>
<gesmes:name>European Central Bank</gesmes:name>
</gesmes:Sender>
<Cube>
  <Cube time="2025-07-18">
    <Cube currency="USD" rate="1.165"/>
    <Cube currency="JPY" rate="172.94"/>
    . . .
    <Cube currency="THB" rate="37.688"/>
    <Cube currency="ZAR" rate="20.6351"/>
  </Cube>
  <Cube time="2025-07-17">
    <Cube currency="USD" rate="1.1579"/>
	<Cube currency="JPY" rate="172.28"/>
	. . .
	<Cube currency="THB" rate="37.678"/>
	<Cube currency="ZAR" rate="20.7253"/>
  </Cube>
  . . .
</Cube>
</gesmes:Envelope>
*/

// The response structure for the reference rates history GET URL.
/*
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Gesmes  string   `xml:"gesmes,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Subject string   `xml:"subject"`
	Sender  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube struct {
		Text string `xml:",chardata"`
		Cube []struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
			Cube []struct {
				Text     string `xml:",chardata"`
				Currency string `xml:"currency,attr"`
				Rate     string `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}
*/

type envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Gesmes  string   `xml:"gesmes,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Subject string   `xml:"subject"`
	Sender  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube struct {
		Text string `xml:",chardata"`
		Cube []struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
			Cube []struct {
				Text     string `xml:",chardata"`
				Currency string `xml:"currency,attr"`
				Rate     string `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}

type What int

const (
	EurFxRefLast What = iota
	EurFxRef90
	EurFxRefFull
)

const (
	eurFxRefLast    = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	eurFxRef90      = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
	eurFxRefFull    = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist.xml"
	eurFxRefReferer = "https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html"
)

var gets = [3]string{eurFxRefLast, eurFxRef90, eurFxRefFull}
var labs = [3]string{eurFxRefReferer, eurFxRefReferer, eurFxRefReferer}
var nams = [3]string{"eurofxref-daily", "eurofxref-hist-90d", "eurofxref-hist"}

type Point struct {
	Date  time.Time
	Value float64
}

type PointSeriesMap map[string][]Point

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
	req.Header.Set("Accept", "application/xml, text/javascript, */*")
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
) (*envelope, error) {
	if bs, err := getWithRetries(gets[what], labs[what], timeout, pauseBeforeRetry, eurFxRefReferer,
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
				nams[what]+".xml")
			if err := os.WriteFile(file, bs, 0644); err != nil {
				log.Printf("cannot write to file %s: %s\n", file, err)
			}
		}

		var envel envelope
		if err := xml.Unmarshal(bs, &envel); err != nil {
			return nil, fmt.Errorf("cannot unmarshal response: %w", err)
		}
		return &envel, nil
	}
}

func Fetch(
	what What,
	writeToFile bool,
	downloadFolder string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	userAgent string,
	verbose bool,
) (PointSeriesMap, error) {
	envel, err := fetch(what, writeToFile, downloadFolder, timeout, pauseBeforeRetry, userAgent, verbose)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch data: %w", err)
	}

	result := make(PointSeriesMap)

	// Iterate in reverse order since last date is the first in XML response
	cs := envel.Cube.Cube
	for i := len(cs) - 1; i >= 0; i-- {
		s := cs[i]
		if len(s.Cube) == 0 {
			continue // Skip empty cubes
		}

		date, err := time.Parse("2006-01-02", s.Time)
		if err != nil {
			return nil, fmt.Errorf("cannot parse date '%s': %w", s.Time, err)
		}

		for _, c := range s.Cube {
			if c.Currency == "" || c.Rate == "" {
				continue // Skip entries with empty currency or rate
			}

			value, err := strconv.ParseFloat(c.Rate, 64)
			if err != nil {
				return nil, fmt.Errorf("cannot parse rate '%s' for currency '%s': %w", c.Rate, c.Currency, err)
			}

			currency := strings.ToUpper(c.Currency)
			if _, exists := result[currency]; !exists {
				result[currency] = make([]Point, 0)
			}
			result[currency] = append(result[currency], Point{Date: date, Value: value})
		}
	}

	return result, nil
}
