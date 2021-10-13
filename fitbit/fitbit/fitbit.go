package fitbit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
)

const (
	base10 = 10
)

// FitBit represents a FitBit OAuth2 session.
type FitBit struct {
	mu         sync.RWMutex
	oauthConfg *oauth2.Config
	httpClient *http.Client
	ratelimit  Ratelimit
}

// Ratelimit includes the rate limit information provided on every request.
type Ratelimit struct {
	RateLimitAvailable int
	RateLimitUsed      int
	RateLimitReset     time.Time
}

// New creates a new FitBit OAuth2 session.
func New(clientID, clientSecret, redirectURL string, scopes []string) (*FitBit, error) {
	f := &FitBit{
		oauthConfg: &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			Endpoint:     fitbit.Endpoint,
		},
	}

	client, err := getClient(f.oauthConfg)
	if err != nil {
		return nil, fmt.Errorf("cannot construct fitbit session: %w", err)
	}

	f.httpClient = client

	return f, nil
}

// GetRatelimit returns the current rate limit information.
// Only available after a request to the API endpoint.
func (f *FitBit) Ratelimit() Ratelimit {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.ratelimit
}

// setCustomHeader sets custom request headers.
func setCustomHeader(req *http.Request) {
	req.Header.Set("User-Agent", "fitbit")
	req.Header.Set("Accept-Language", "de_DE")
	req.Header.Set("Accept-Locale", "de_DE")
}

// getRateLimit extracts rate limit data from the header of the response.
func (f *FitBit) getRateLimit(resp *http.Response) {
	f.mu.Lock()
	defer f.mu.Unlock()

	rateLimitData := resp.Header.Get("fitbit-rate-limit-remaining")
	if rateLimitData != "" {
		f.ratelimit.RateLimitUsed, _ = strconv.Atoi(rateLimitData)
	}

	rateLimitData = resp.Header.Get("fitbit-rate-limit-limit")
	if rateLimitData != "" {
		f.ratelimit.RateLimitAvailable, _ = strconv.Atoi(rateLimitData)
	}

	rateLimitData = resp.Header.Get("fitbit-rate-limit-reset")
	if rateLimitData != "" {
		remSec, _ := strconv.Atoi(rateLimitData)
		f.ratelimit.RateLimitReset = time.Now().Add(time.Second * time.Duration(remSec))
	}
}

// makeGETRequest makes a new GET request to a given URL using the OAuth2-enabled HTTP client.
func (f *FitBit) makeGETRequest(targetURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}

	setCustomHeader(req)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f.getRateLimit(resp)

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// makePOSTRequest makes a new POST request to a given URL using the OAuth2-enabled HTTP client.
func (f *FitBit) makePOSTRequest(targetURL string, param map[string]string) ([]byte, error) {
	form := url.Values{}
	for name, value := range param {
		form.Add(name, value)
	}

	req, err := http.NewRequest("POST", targetURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	setCustomHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f.getRateLimit(resp)

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// makeDELETERequest makes a new DELETE request to a given URL using the OAuth2-enabled HTTP client.
func (f *FitBit) makeDELETERequest(targetURL string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", targetURL, nil)
	if err != nil {
		return nil, err
	}

	setCustomHeader(req)

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f.getRateLimit(resp)

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

// convertToRequestID converts user ID to a request ID.
// The userID 0 is a current logged in user.
func convertToRequestID(userID uint64) string {
	// Default "-" is the current logged in user.
	requestID := "-"
	if userID > 0 {
		requestID = strconv.FormatUint(userID, base10)
	}

	return requestID
}
