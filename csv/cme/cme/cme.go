package cme

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
{
  "props":{
    "pageNumber":1,
    "pageTotal":1,
    "pageSize":12},
  "entries":[
    {"size":"1","indicator":"-","price":"4214.5","date":"18 Apr 2023","time":"17:59:58","strike":"0","type":" ","key":5916398547969},
    {"size":"1","indicator":"-","price":"4214","date":"18 Apr 2023","time":"17:57:44","strike":"0","type":" ","key":5916398547968},
    {"size":"1","indicator":"-","price":"4214","date":"18 Apr 2023","time":"17:16:46","strike":"0","type":" ","key":5916398547967},
    {"size":"0","indicator":"Open","price":"4214.5","date":"18 Apr 2023","time":"17:12:27","strike":"0","type":" ","key":5916398547966},
    {"size":"1","indicator":"-","price":"4214.5","date":"18 Apr 2023","time":"17:12:27","strike":"0","type":" ","key":5916398547965}
  ]
  "tradeDate":"18 Apr 2023",
  "productDescription":"E-mini S&P 500 Futures Sep 2023 Globex"
}*/

type RealtimeEntry struct {
	Key       int64  `json:"key"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Indicator string `json:"indicator"`
	Type      string `json:"type"`
}

type RealtimeProps struct {
	PageNumber int `json:"pageNumber"`
	PageTotal  int `json:"pageTotal"`
	PageSize   int `json:"pageSize"`
}

type Realtime struct {
	Props              RealtimeProps   `json:"props"`
	Entries            []RealtimeEntry `json:"entries"`
	TradeDate          string          `json:"tradeDate"`
	ProductDescription string          `json:"productDescription"`
}

type RealtimeEntryParsed struct {
	Key       int64
	Time      time.Time
	Price     float64
	Size      int
	Indicator string
	Type      string
}

type RealtimePageJSON struct {
	Page int
	JSON []byte
}

type RealtimeTimeslotJSON struct {
	Timeslot string
	JSON     []RealtimePageJSON
}

type RealtimeJSON struct {
	Extension string
	Timeslots []RealtimeTimeslotJSON
}

type FutureSymbol struct {
	Future        string `json:"future"`        // ES
	Mnemonic      string `json:"mnemonic"`      // ESM23 (Globex code)
	Name          string `json:"name"`          // E-mini S&P 500 Futures
	Code          string `json:"code"`          // 133
	ContractMonth string `json:"contractMonth"` // JUN 2023
	ContractCode  string `json:"contractCode"`  // M3
	Currency      string `json:"currency"`      // USD
	Mic           string `json:"mic"`           // CMX
}

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
const referer = "https://www.cmegroup.com/markets/equities/sp/e-mini-sandp500.timeAndSales.html"
const empty = ""

func str2datetime(d, t string) (time.Time, error) {
	// Mon Jan 2 15:04:05 -0700 MST 2006, "18 Apr 2023", "17:12:27"
	const layout = "2 Jan 2006 15:04:05"
	return time.Parse(layout, d+" "+t)
}

func (re *RealtimeEntry) parse(daydelta int) (*RealtimeEntryParsed, error) {
	s := strings.TrimSpace(re.Date)
	t, err := str2datetime(s, strings.TrimSpace(re.Time))
	if err != nil {
		return nil, fmt.Errorf("cannot parse date '%s' and time '%s': %w", re.Date, re.Time, err)
	}

	t = t.AddDate(0, 0, daydelta)

	s = strings.ReplaceAll(re.Price, ",", empty)
	s = strings.TrimLeft(s, "$")
	s = strings.TrimSpace(s)
	p, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse price '%s': %w", re.Price, err)
	}

	s = strings.ReplaceAll(re.Size, ",", empty)
	s = strings.TrimSpace(s)
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse size '%s': %w", re.Size, err)
	}

	return &RealtimeEntryParsed{
		Key:       re.Key,
		Time:      t,
		Price:     p,
		Size:      int(v),
		Indicator: re.Indicator,
		Type:      re.Type,
	}, nil
}

func unmarshalRealtime(jsn []byte) (*Realtime, error) {
	rt := Realtime{}
	if err := json.Unmarshal(jsn, &rt); err != nil {
		return nil, fmt.Errorf("cannot unmarshal Realtime: %w", err)
	}

	return &rt, nil
}

func get(targetURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", referer)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encoding", "identity")
	httpClient := http.Client{Timeout: time.Duration(120) * time.Second}

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

	return contents, nil
}

func convertToCSV(series []RealtimeEntryParsed) []string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006-01-02 15:04:05"

	csv := make([]string, 0)
	for _, p := range series {
		s := fmt.Sprintf("%s;%v;%v;%v\n", p.Time.Format(layout), p.Price, p.Size, p.Key)
		csv = append(csv, s)
	}

	return csv
}

// RetrieveTimeslot downloads all time and sales pages for the given timeslot.
func RetrieveTimeslot(sym FutureSymbol, entryDate time.Time,
	timeslot, daydelta int) ([]RealtimePageJSON, []RealtimeEntryParsed, error) {
	parsed := make([]RealtimeEntryParsed, 0)
	jsons := make([]RealtimePageJSON, 0)

	if daydelta != 0 {
		entryDate = entryDate.AddDate(0, 0, daydelta)
	}

	// sym.Code=133 and sym.ContractCode=U3 for ESU23
	url := "https://www.cmegroup.com/CmeWS/mvc/TimeandSales/" + sym.Code +
		"/G/" + sym.ContractCode +
		"?timeSlot=" + strconv.Itoa(timeslot) +
		"&entryDate=" + entryDate.Format("20060102") +
		"&pageSize=12&pageNumber="
	pg := 1
	pg100 := 100

nextPage:
	page := url + strconv.Itoa(pg)
	if pg == 1 {
		fmt.Printf("getting %s\n", page)
	}

	b, err := get(page)
	if err != nil {
		return jsons, parsed, fmt.Errorf("cannot get url '%s': %w", page, err)
	}

	jsons = append(jsons, RealtimePageJSON{
		Page: pg,
		JSON: b,
	})

	rt, err := unmarshalRealtime(b)
	if err != nil {
		return jsons, parsed, fmt.Errorf("cannot unmarshal url '%s': %w\n%s", page, err, string(b))
	}

	for i := len(rt.Entries) - 1; i >= 0; i-- {
		el := rt.Entries[i]
		cv, err := el.parse(daydelta)
		if err != nil {
			return jsons, parsed, fmt.Errorf("cannot convert url '%s': %w", page, err)
		}

		parsed = append(parsed, *cv)
	}

	if pg < rt.Props.PageTotal {
		if pg == 1 {
			fmt.Printf("total pages %d\n", rt.Props.PageTotal)
		}
		pg++
		if pg == pg100 {
			fmt.Printf("..%d", pg100)
			pg100 += 100
		}
		goto nextPage
	}

	if pg100 > 100 {
		fmt.Printf("\n")
	}

	return jsons, parsed, nil
}

// RetrieveCode retrievs all timeslots for the given code.
func RetrieveCode(sym FutureSymbol, entryDate time.Time) ([]string, []RealtimeEntryParsed, RealtimeJSON, error) {
	timeSlots := []int{
		17, // 17:00 - 17:59:59
		18, // 18:00 - 18:59:59
		19, // 19:00 - 19:59:59
		20, // 20:00 - 20:59:59
		21, // 21:00 - 21:59:59
		22, // 22:00 - 22:59:59
		23, // 23:00 - 23:59:59
		0,  // 00:00 - 00:59:59
		1,  // 01:00 - 01:59:59
		2,  // 02:00 - 02:59:59
		3,  // 03:00 - 03:59:59
		4,  // 04:00 - 04:59:59
		5,  // 05:00 - 05:59:59
		6,  // 06:00 - 06:59:59
		7,  // 07:00 - 07:59:59
		8,  // 08:00 - 08:59:59
		9,  // 00:00 - 09:59:59
		10, // 10:00 - 10:59:59
		11, // 11:00 - 11:59:59
		12, // 12:00 - 12:59:59
		13, // 13:00 - 13:59:59
		14, // 14:00 - 14:59:59
		15} // 15:00 - 15:59:59

	csv := make([]string, 0)
	series := make([]RealtimeEntryParsed, 0)
	jsons := RealtimeJSON{
		Extension: sym.ContractCode,
		Timeslots: make([]RealtimeTimeslotJSON, 0),
	}

	for _, p := range timeSlots {
		daydelta := 0
		if p > 16 {
			daydelta = -1
		}

		pages, entries, err := RetrieveTimeslot(sym, entryDate, p, daydelta)
		if err != nil {
			return csv, series, jsons, fmt.Errorf("cannot retrieve timeslot '%d': %w", p, err)
		}

		ts := empty
		if p < 10 {
			ts = "0"
		}

		jsons.Timeslots = append(jsons.Timeslots, RealtimeTimeslotJSON{
			Timeslot: ts + strconv.Itoa(p),
			JSON:     pages,
		})

		series = append(series, entries...)
	}

	sort.Slice(series, func(i, j int) bool {
		return series[i].Key < series[j].Key
	})

	return convertToCSV(series), series, jsons, nil
}
