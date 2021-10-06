package fitbit

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"
)

type API struct {
	ClientID, ClientSecret, RedirectURI, AuthorizeURI, Scopes string
	Auth                                                      FitbitAuth
}

type FitbitAuth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type ActivitySteps struct {
	Steps []DataPoint `json:"activities-steps"`
}

type DataPoint struct {
	Time  string `json:"dateTime"`
	Value string `json:"value"`
}

type HeartDataPoint struct {
	Date  string         `json:"dateTime"`
	Value HeartDataValue `json:"value"`
}

type HeartDataValue struct {
	RestingHeartRate int `json:"restingHeartRate"`
}

type ActivityHeart struct {
	HeartData []HeartDataPoint `json:"activities-heart"`
}

type HeartateIntradayPoint struct {
	Time  string `json:"time"`
	Value int    `json:"value"`
}

type HeartIntraday struct {
	Dataset         []HeartateIntradayPoint `json:"dataset"`
	DatasetInterval int                     `json:"DatasetInterval"`
	DatasetType     string                  `json:"datasetType"`
}
type ActivityHeartSeries struct {
	HeartData     []HeartDataPoint `json:"activities-heart"`
	HeartIntraday HeartIntraday    `json:"activities-heart-intraday"`
}

type NormalisedHeartPoint struct {
	Timestamp time.Time
	Value     int
}

func New(clientID, clientSecret, redirectURI string) API {
	authorizeTmpl := "response_type=code&client_id={{.ClientID}}&redirect_uri={{.RedirectURI}}&scope={{.Scopes}}&expires_in=604800"

	tmpl, err := template.New("test").Parse(authorizeTmpl)
	if err != nil {
		panic(err)
	}

	api := API{}
	api.ClientID = clientID
	api.ClientSecret = clientSecret
	api.RedirectURI = redirectURI
	api.Scopes = "activity heartrate location nutrition profile settings sleep social weight"

	var authorizeBuf bytes.Buffer

	err = tmpl.Execute(&authorizeBuf, api)
	if err != nil {
		panic(err)
	}

	api.AuthorizeURI = "https://www.fitbit.com/oauth2/authorize?" + url.PathEscape(authorizeBuf.String())

	return api
}

func (api *API) EncodeBasicAuth() string {
	authstring := api.ClientID + ":" + api.ClientSecret
	return base64.StdEncoding.EncodeToString([]byte(authstring))
}

func (api *API) GetProfile() string {
	req, _ := http.NewRequest("GET", "https://api.fitbit.com/1/user/-/profile.json", nil)
	req.Header.Set("Authorization", "Bearer "+api.Auth.AccessToken)
	res, _ := http.DefaultClient.Do(req)
	profiledata, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return string(profiledata)
}

func (api *API) GetActivitySteps() ActivitySteps {
	req, _ := http.NewRequest("GET", "https://api.fitbit.com/1/user/-/activities/steps/date/today/1y.json", nil)
	req.Header.Set("Authorization", "Bearer "+api.Auth.AccessToken)
	res, _ := http.DefaultClient.Do(req)

	var activitySteps ActivitySteps

	decoder := json.NewDecoder(res.Body)

	decerr := decoder.Decode(&activitySteps)
	if decerr != nil {
		panic(decerr)
	}

	res.Body.Close()

	return activitySteps
}

func (api *API) GetRestingHeartrate() ActivityHeart {
	req, _ := http.NewRequest("GET", "https://api.fitbit.com/1/user/-/activities/heart/date/today/1y.json", nil)
	req.Header.Set("Authorization", "Bearer "+api.Auth.AccessToken)
	res, _ := http.DefaultClient.Do(req)

	var activityHeart ActivityHeart

	decoder := json.NewDecoder(res.Body)

	decerr := decoder.Decode(&activityHeart)
	if decerr != nil {
		panic(decerr)
	}

	res.Body.Close()

	return activityHeart
}

func (series *ActivityHeartSeries) GetNormalisedSeries(timezone string) []NormalisedHeartPoint {
	loc, _ := time.LoadLocation(timezone)

	var points []NormalisedHeartPoint

	format := "2006-01-02 15:04:05"
	for _, datapoint := range series.HeartIntraday.Dataset {
		timestamp, e := time.ParseInLocation(format, series.HeartData[0].Date+" "+datapoint.Time, loc)
		if e != nil {
			panic(e)
		}

		points = append(points, NormalisedHeartPoint{Timestamp: timestamp, Value: datapoint.Value})
	}

	return points
}

func (api *API) GetHeartrateTimeSeries(date string) ActivityHeartSeries {
	req, _ := http.NewRequest("GET", "https://api.fitbit.com/1/user/-/activities/heart/date/"+date+"/1d/1sec.json", nil)
	req.Header.Set("Authorization", "Bearer "+api.Auth.AccessToken)
	res, _ := http.DefaultClient.Do(req)

	var series ActivityHeartSeries

	dec := json.NewDecoder(res.Body)

	decerr := dec.Decode(&series)
	if decerr != nil {
		panic(decerr)
	}

	return series
}

func (api *API) LoadAccessToken(code string) {
	form := url.Values{}
	form.Add("clientId", api.ClientID)
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", api.RedirectURI)
	form.Add("code", code)

	req, err := http.NewRequest("POST", "https://api.fitbit.com/oauth2/token", strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+api.EncodeBasicAuth())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(res.Body)

	var auth FitbitAuth

	decerr := decoder.Decode(&auth)
	if decerr != nil {
		panic(decerr)
	}

	res.Body.Close()

	api.Auth = auth
}
