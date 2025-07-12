package intraday

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func getLastWorkingDay(startDateDaysBack int) string {
	date := time.Now()

	// Go back to the most recent weekday (Monday-Friday)
	for date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		date = date.AddDate(0, 0, -1)
	}

	// Go back additional days as specified
	date = date.AddDate(0, 0, -startDateDaysBack)

	return date.Format("2006-01-02")
}

func post(
	uri string,
	timeout time.Duration,
	referer string,
	userAgent string,
	startDateDaysBack int,
	verbose bool,
) ([]byte, error) {
	bodyMap := map[string]string{
		// "startTime": "08:00",
		// "endTime":   "20:00",
		"nbitems":  "900000",
		"timezone": "CET",
		"date":     getLastWorkingDay(startDateDaysBack),
	}

	// Prepare POST data
	postData := url.Values{}
	for key, value := range bodyMap {
		postData.Set(key, value)
	}

	pd := postData.Encode()
	if verbose {
		log.Printf("%s POST body: %s\n", uri, pd)
	}
	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(pd))
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", referer)
	req.Header.Set("Accept-Language", "en-us,en;q=0.5")
	req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept", "application/json, text/javascript, */*")

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

func postWithRetries(
	uri string,
	label string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	referer string,
	userAgent string,
	startDateDaysBack int,
	verbose bool,
) ([]byte, error) {
	var contents []byte
	var err error
	retriesMax := len(pauseBeforeRetry)
	retries := retriesMax
	for retries > 0 {
		contents, err = post(uri, timeout, referer, userAgent, startDateDaysBack, verbose)
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

func getURI(isin string, mic string) string {
	return fmt.Sprintf(
		"https://live.euronext.com/en/ajax/getIntradayPriceFilteredData/%s-%s",
		strings.ToUpper(isin), strings.ToUpper(mic))
}

func getReferer(isin string, mic string, typ string) string {
	switch typ {
	case "index":
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/indices/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	case "stock":
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/equities/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	case "etv":
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/etvs/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	case "etf":
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/etfs/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	case "inav":
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/indices/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	case "fund":
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/funds/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	default:
		return fmt.Sprintf(
			"https://live.euronext.com/en/product/equities/%s-%s/quotes",
			strings.ToUpper(isin), strings.ToUpper(mic))
	}
}

// FetchIntradayData retrieves intraday data for a given instrument identified by its MIC and ISIN.
func FetchIntradayData(
	isin string,
	mic string,
	mnemonic string,
	typ string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	userAgent string,
	startDateDaysBack int,
	verbose bool,
) ([]byte, error) {
	uri := getURI(isin, mic)
	ref := getReferer(isin, mic, typ)
	label := fmt.Sprintf("%s-%s-%s-%s", mic, typ, mnemonic, isin)
	if bs, err := postWithRetries(uri, label, timeout, pauseBeforeRetry, ref, userAgent,
		startDateDaysBack, verbose); err != nil {
		return nil, err
	} else {
		return bs, nil
	}
}
