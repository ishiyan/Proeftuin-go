package nasdaq

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
{
	data: {
		symbol:"tsla",
		totalRecords:16854,
		offset:0,
		limit:99999999,
		headers: {
			nlsTime:"NLS Time (ET)",
			nlsPrice:"NLS Price",
			nlsShareVolume:"NLS Share Volume"
		},
		rows:[
			{ nlsTime:"11:59:59", nlsPrice:"$ 200.461", nlsShareVolume:"200" },
			{ nlsTime:"11:30:59", nlsPrice:"$ 198.24", nlsShareVolume:"100"}
		],
		topTable:{
			headers:{
				nlsVolume:"Nasdaq Last Sale (NLS) Plus Volume",
				previousClose:"Previous Close",
				todayHighLow:"Today's High / Low*",
				fiftyTwoWeekHighLow:"52 Week High / Low"
			},
			rows:[
				{
					nlsVolume:"0",
					previousClose:"$0.00",
					todayHighLow:"$0.00/$0.00",
					fiftyTwoWeekHighLow:"$384.29/$101.81"
				}
			]
		},
		description:{
			message:"*Today’s High/Low is only updated during regular trading hours; and does not include trades occurring in pre-market or after-hours.",
			url:null
		},
		message:[
			"Data last updated Feb 17, 2023",
			"This page will resume updating on Feb 21, 2023 04:00 AM ET"
		]
	},
	message:null,
	status:{
		rCode:200,
		bCodeMessage:null,
		developerMessage:null
	}
}
*/

type NasdaqClient struct{}

type NasdaqRealtimeLastSale struct {
	Time   string `json:"nlsTime"`
	Price  string `json:"nlsPrice"`
	Volume string `json:"nlsShareVolume"`
}

type NasdaqRealtimeData struct {
	Symbol       string                   `json:"symbol"`
	TotalRecords int                      `json:"totalRecords"`
	Offset       int                      `json:"offset"`
	Limit        int                      `json:"limit"`
	Rows         []NasdaqRealtimeLastSale `json:"rows"`
	Message      []string                 `json:"message"`
}

type NasdaqRealtimeStatus struct {
	Code int `json:"rCode"`
}

type NasdaqRealtime struct {
	Data    NasdaqRealtimeData   `json:"data"`
	Message string               `json:"message"`
	Status  NasdaqRealtimeStatus `json:"status"`
}

type NasdaqRealtimeJSON struct {
	Period string
	JSON   []byte
}

type NasdaqTrade struct {
	Time   time.Time
	Price  float64
	Volume float64
}

func (nrls *NasdaqRealtimeLastSale) convertToNasdaqTrade(date time.Time) (*NasdaqTrade, error) {
	t, err := time.Parse("15:04:05", nrls.Time)
	if err != nil {
		return nil, fmt.Errorf("cannot parse time '%s': %w", nrls.Time, err)
	}

	s := strings.TrimLeft(nrls.Price, "$")
	p, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse price '%s': %w", nrls.Price, err)
	}

	v, err := strconv.ParseFloat(nrls.Volume, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse volume '%s': %w", nrls.Volume, err)
	}

	return &NasdaqTrade{
		Time:   t.AddDate(date.Year(), int(date.Month()), date.Day()),
		Price:  p,
		Volume: v,
	}, nil
}

func unmarshalRealtime(jsn []byte) (*NasdaqRealtime, error) {
	rt := NasdaqRealtime{}
	if err := json.Unmarshal(jsn, &rt); err != nil {
		return nil, fmt.Errorf("cannot unmarshal NasdaqRealtime: %w", err)
	}

	return &rt, nil
}

func get(targetURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", "nq")

	httpClient := http.Client{Timeout: time.Duration(300) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot do request: %w", err)
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	return contents, nil
}

func sessionDate() (time.Time, error) {
	today := time.Now()
	loc, err := time.LoadLocation("EST")
	if err != nil {
		return today, fmt.Errorf("cannot load EST timezone: %w", err)
	}

	today = today.In(loc)
	dow := today.Weekday()
	if dow == time.Saturday {
		return today.AddDate(0, 0, -1), nil
	} else if dow == time.Sunday {
		return today.AddDate(0, 0, -2), nil
	} else {
		if today.Hour() < 9 {
			if today.Minute() < 25 {
				return today.AddDate(0, 0, -1), nil
			}
		}

		if today.Hour() < 15 && today.Minute() < 35 {
			if today.Minute() < 20 {
				return today, fmt.Errorf("time now should be inside 09:25 till 15:35: %v", today)
			}
		}

		return today, nil
	}
}

func RetrieveSession(mnemonic string) ([]NasdaqTrade, []NasdaqRealtimeJSON, error) {
	periods := []string{
		"09:30",
		"10:00",
		"10:30",
		"11:00",
		"11:30",
		"12:00",
		"12:30",
		"13:00",
		"13:30",
		"14:00",
		"14:30",
		"15:00",
		"15:30"}

	series := make([]NasdaqTrade, 0)
	jsons := make([]NasdaqRealtimeJSON, 0)
	today, err := sessionDate()
	if err != nil {
		return series, jsons, err
	}

	url := "https://api.nasdaq.com/api/quote/" + mnemonic + "/realtime-trades?&limit=99999999&fromTime="

	for _, p := range periods {
		b, err := get(url + p)
		if err != nil {
			return series, jsons, fmt.Errorf("cannot get url '%s': %w", url+p, err)
		}

		jsons = append(jsons, NasdaqRealtimeJSON{
			Period: strings.ReplaceAll(p, ":", "-"),
			JSON:   b,
		})

		rt, err := unmarshalRealtime(b)
		if err != nil {
			return series, jsons, fmt.Errorf("cannot unmarshal url '%s': %w", url+p, err)
		}

		for i := len(rt.Data.Rows) - 1; i >= 0; i-- {
			el := rt.Data.Rows[i]
			cv, err := el.convertToNasdaqTrade(today)
			if err != nil {
				return series, jsons, fmt.Errorf("cannot convert url '%s': %w", url+p, err)
			}

			series = append(series, *cv)
		}
	}

	return series, jsons, nil
}
