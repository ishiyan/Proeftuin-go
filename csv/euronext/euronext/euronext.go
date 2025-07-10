package euronext

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

func get(targetURL string) ([]byte, error) {
repeat:
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	httpClient := http.Client{Timeout: time.Duration(60) * time.Second}

	var resp *http.Response
	const retriesMax = 5
	retries := retriesMax
	for retries > 0 {
		resp, err = httpClient.Do(req)
		if err != nil {
			err := fmt.Errorf("cannot do request, retries (%d of %d): %w", retries, retriesMax, err)
			fmt.Println(err)
			retries -= 1
			if retries < 1 {
				return nil, err
			} else {
				time.Sleep(2 * time.Second)
			}
		} else {
			break
		}
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	if contents[0] == '<' {
		goto repeat
	}

	return contents, nil
}

func getEodHistoryURL(isin string, mic string, isAdjusted bool) string {
	var adjusted string
	if isAdjusted {
		adjusted = "Y"
	} else {
		adjusted = "N"
	}

	return fmt.Sprintf(
		"https://live.euronext.com/en/ajax/AwlHistoricalPrice/getFullDownloadAjax/%s-%s"+
			"?format=csv&decimal_separator=.&date_form=d%%2Fm%%2FY&op=&&adjusted=%s"+
			"&base100=&startdate=2000-01-01&enddate=2034-12-31",
		strings.ToUpper(isin), strings.ToUpper(mic), adjusted)
}

func getEodHistory(isin string, mic string, isAdjusted bool) ([]byte, error) {
	url := getEodHistoryURL(isin, mic, isAdjusted)
	if bs, err := get(url); err != nil {
		return nil, fmt.Errorf("cannot retrieve %s: %w", url, err)
	} else {
		return bs, nil
	}
}

func DownloadEodHistory(isin string, mic string) ([]byte, []byte, error) {
	bs_adj, err := getEodHistory(isin, mic, true)
	if err != nil {
		return bs_adj, nil, fmt.Errorf("cannot get EOD adjusted history: %w", err)
	}

	bs_raw, err := getEodHistory(isin, mic, false)
	if err != nil {
		return bs_adj, bs_raw, fmt.Errorf("cannot get EOD not-adjusted history: %w", err)
	}

	return bs_adj, bs_raw, nil
}

func EnsureDirectoryExists(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = os.MkdirAll(directory, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory '%s': %w", directory, err)
		}
	}

	return nil
}

func SessionDate() (time.Time, error) {
	today := time.Now().UTC().Add(time.Hour * 1)
	dow := today.Weekday()
	switch dow {
	case time.Saturday:
		return today.AddDate(0, 0, -1), nil
	case time.Sunday:
		return today.AddDate(0, 0, -2), nil
	default:
		if today.Hour() < 19 {
			return today.AddDate(0, 0, -1), nil
		}
		return today, nil
	}
}

func BackupFile(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", nil
	}

	backupFilename := fmt.Sprintf("%s.bak", filename)

	if _, err := os.Stat(backupFilename); err == nil {
		if err := os.Remove(backupFilename); err != nil {
			es := fmt.Sprintf("cannot delete existing backup file '%s': ", backupFilename)
			return es, fmt.Errorf("%s%w", es, err)
		}
	}

	if err := os.Rename(filename, backupFilename); err != nil {
		es := fmt.Sprintf("cannot rename file '%s' to '%s': ", filename, backupFilename)
		return es, fmt.Errorf("%s%w", es, err)
	}

	return "", nil
}
