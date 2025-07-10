package intraday

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func get(
	uri string,
	timeout time.Duration,
	referer string,
	userAgent string,
) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", referer)
	req.Header.Set("Accept-Language", "en-us,en;q=0.5")
	req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept", "application/json, text/javascript, */*")

	httpClient := http.Client{Timeout: timeout}
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
) ([]byte, error) {
	var contents []byte
	var err error
	retriesMax := len(pauseBeforeRetry)
	retries := retriesMax
	for retries > 0 {
		contents, err = get(uri, timeout, referer, userAgent)
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

// GetIntradayData retrieves intraday data for a given instrument identified by its MIC and ISIN.
func GetIntradayData(
	isin string,
	mic string,
	mnemonic string,
	typ string,
	timeout time.Duration,
	pauseBeforeRetry []time.Duration,
	userAgent string,
) ([]byte, error) {
	uri := getURI(isin, mic)
	ref := getReferer(isin, mic, typ)
	label := fmt.Sprintf("%s-%s-%s-%s", mic, typ, mnemonic, isin)
	if bs, err := getWithRetries(uri, label, timeout, pauseBeforeRetry, ref, userAgent); err != nil {
		return nil, err
	} else {
		return bs, nil
	}
}
