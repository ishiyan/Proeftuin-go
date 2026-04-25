// enxdecrypt decrypts CryptoJS AES encrypted responses from the Euronext API.
//
// Usage:
//
//	enxdecrypt [-p passphrase] < encrypted.json
//	enxdecrypt [-p passphrase] encrypted.json
//	enxdecrypt [-p passphrase] -url https://live.euronext.com/en/ajax/getIntradayPriceFilteredData/FR0003500008-XPAR
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"euronext/euronext/intraday"
)

func main() {
	passphrase := flag.String("p", intraday.DefaultPassphrase, "decryption passphrase")
	urlFlag := flag.String("url", "", "URL to fetch and decrypt")
	checkKey := flag.Bool("checkkey", false, "fetch the latest passphrase from Euronext before decrypting")
	flag.Parse()

	if *checkKey {
		fetched, err := intraday.FetchPassphrase()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot fetch passphrase: %v (continuing with current passphrase)\n", err)
		} else if fetched != *passphrase {
			fmt.Fprintf(os.Stderr, "passphrase changed: %s\n", fetched)
			*passphrase = fetched
		}
	}

	var data []byte
	var err error

	switch {
	case *urlFlag != "":
		data, err = fetchURL(*urlFlag)
	case flag.NArg() > 0:
		data, err = os.ReadFile(flag.Arg(0))
	default:
		data, err = io.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	if !intraday.IsEncryptedResponse(data) {
		// Not encrypted — print as-is.
		fmt.Print(string(data))
		return
	}

	plaintext, err := intraday.DecryptResponse(data, *passphrase)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println(string(plaintext))
}

func fetchURL(rawURL string) ([]byte, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
