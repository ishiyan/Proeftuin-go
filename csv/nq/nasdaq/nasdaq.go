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

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
const cookieName = "ak_bmsc"
const empty = ""

var cookieValue = empty

func (nrls *NasdaqRealtimeLastSale) convertToNasdaqTrade(dateEst time.Time) (*NasdaqTrade, error) {

	s := strings.TrimSpace(nrls.Time)
	t, err := time.Parse("15:04:05", s)
	if err != nil {
		return nil, fmt.Errorf("cannot parse time '%s': %w", nrls.Time, err)
	}

	s = strings.ReplaceAll(nrls.Price, ",", empty)
	s = strings.TrimLeft(s, "$")
	s = strings.TrimSpace(s)
	p, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse price '%s': %w", nrls.Price, err)
	}

	s = strings.ReplaceAll(nrls.Volume, ",", empty)
	s = strings.TrimSpace(s)
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse volume '%s': %w", nrls.Volume, err)
	}

	return &NasdaqTrade{
		Time:   t.AddDate(dateEst.Year(), int(dateEst.Month())-1, dateEst.Day()-1),
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

func ResetCookie() {
	cookieValue = empty
}

func getCookie() error {
	//req, err := http.NewRequest("GET", "https://www.nasdaq.com/market-activity/stocks/adbe", nil)
	req, err := http.NewRequest("GET", "https://api.nasdaq.com/api/quote/aapl/realtime-trades", nil)
	if err != nil {
		cookieValue = empty
		return fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	httpClient := http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		cookieValue = empty
		return fmt.Errorf("cannot do request: %w", err)
	}
	defer resp.Body.Close()

	for _, q := range resp.Cookies() {
		fmt.Println("response cookie: " + q.Name + " value: " + q.Value)
		if q.Name == cookieName {
			cookieValue = q.Value
			fmt.Println("obtained access cookie: " + cookieValue)
			return nil
		}
	}

	cookieValue = empty
	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("no '%s' cookie found in %d response cookies, body is: %s", cookieName, len(resp.Cookies()), string(body))
}

func get(targetURL string) ([]byte, error) {
repeat:
	//if cookieValue == empty {
	//	_ = getCookie()
	//	//if err := getCookie(); err != nil {
	//	//	return nil, fmt.Errorf("cannot obtain access cookie: %w", err)
	//	//}
	//}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	//if cookieValue != empty {
	//	req.AddCookie(&http.Cookie{Name: cookieName, Value: cookieValue})
	//}
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
		cookieValue = empty
		goto repeat
	}

	return contents, nil
}

func sessionDate() (time.Time, error) {
	today := time.Now().Add(time.Hour * -6)
	/*loc, err := time.LoadLocation("EST")
	if err != nil {
		return today, fmt.Errorf("cannot load EST timezone: %w", err)
	}

	today = today.In(loc)*/
	dow := today.Weekday()
	/*if dow == time.Tuesday {
		return today.AddDate(0, 0, -1), nil
	} else if dow == time.Wednesday {
		return today.AddDate(0, 0, -2), nil
	} else if dow == time.Thursday {
		return today.AddDate(0, 0, -3), nil
	} else*/ if dow == time.Saturday {
		return today.AddDate(0, 0, -1), nil
	} else if dow == time.Sunday {
		return today.AddDate(0, 0, -2), nil
	} else {
		/*if today.Hour() < 9 || (today.Hour() == 9 && today.Minute() <= 25) {
			return today.AddDate(0, 0, -1), nil
		}

		if today.Hour() < 15 || (today.Hour() == 15 && today.Minute() <= 35) {
			return today, fmt.Errorf("time now should be outside 09:25 till 15:35 EST: %v", today)
		}*/

		return today, nil
	}
}

func convertToCSV(series []NasdaqTrade) []string {
	const timeFormat = "2006-01-02 15:04:05"

	csv := make([]string, 0)
	for _, p := range series {
		s := fmt.Sprintf("%s;%v;%v\n", p.Time.Format(timeFormat), p.Price, p.Volume)
		csv = append(csv, s)
	}

	return csv
}

func RetrieveSession(mnemonic string) (time.Time, []string, []NasdaqRealtimeJSON, error) {
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

	csv := make([]string, 0)
	series := make([]NasdaqTrade, 0)
	jsons := make([]NasdaqRealtimeJSON, 0)
	today, err := sessionDate()
	if err != nil {
		return today, csv, jsons, err
	}

	url := "https://api.nasdaq.com/api/quote/" + mnemonic + "/realtime-trades?limit=99999999&fromTime="

	for _, p := range periods {
		b, err := get(url + p)
		if err != nil {
			return today, csv, jsons, fmt.Errorf("cannot get url '%s': %w", url+p, err)
		}

		jsons = append(jsons, NasdaqRealtimeJSON{
			Period: strings.ReplaceAll(p, ":", "-"),
			JSON:   b,
		})

		rt, err := unmarshalRealtime(b)
		if err != nil {
			return today, csv, jsons, fmt.Errorf("cannot unmarshal url '%s': %w", url+p, err)
		}

		for i := len(rt.Data.Rows) - 1; i >= 0; i-- {
			el := rt.Data.Rows[i]
			cv, err := el.convertToNasdaqTrade(today)
			if err != nil {
				return today, csv, jsons, fmt.Errorf("cannot convert url '%s': %w", url+p, err)
			}

			series = append(series, *cv)
		}
	}

	csv = convertToCSV(series)
	return today, csv, jsons, nil
}

type NasdaqSymbol struct {
	Mnemonic     string `json:"mnemonic"`     // stocks, etf
	Name         string `json:"name"`         // stocks, etf
	Company      string `json:"company"`      // stocks
	CompanyUrl   string `json:"companyUrl"`   // stocks
	Description  string `json:"description"`  // stocks
	Currency     string `json:"currency"`     // stocks, etf
	Mic          string `json:"mic"`          // stocks, etf
	Exchange     string `json:"exchange"`     // stocks, etf
	IsListed     bool   `json:"isListed"`     // stocks, etf
	IsNasdaq100  bool   `json:"isNasdaq100"`  // stocks, etf
	StockType    string `json:"stockType"`    // stocks
	AssetClass   string `json:"assetClass"`   // stosks, etf
	Industry     string `json:"industry"`     // stocks
	Sector       string `json:"sector"`       // stocks
	Region       string `json:"region"`       // stocks
	ExpenseRatio string `json:"expenseRatio"` // etf
}

/*
{
    "data":{
        "filters":null,
        "table":{
            "headers":{
                "symbol":"Symbol","name":"Name","lastsale":"Last Sale","netchange":"Net Change","pctchange":"% Change","marketCap":"Market Cap"
			},
            "rows":[
                {"symbol":"AAPL","name":"Apple Inc. Common Stock","lastsale":"$146.57","netchange":"-2.83","pctchange":"-1.894%","marketCap":"2,541,133,923,800","url":"/market-activity/stocks/aapl"},
                {"symbol":"ZTAQW","name":"Zimmer Energy Transition Acquisition Corp. Warrants","lastsale":"$0.1804","netchange":"-0.02","pctchange":"-9.98%","marketCap":"NA","url":"/market-activity/stocks/ztaqw"}
            ]
        },
        "totalrecords":7786,
        "asof":"Last price as of Feb 24, 2023"
    },
    "message":null,
    "status":{"rCode":200,"bCodeMessage":null,"developerMessage":null}
}
{
	"data":{
		"dataAsOf":"2/24/2023 8:00:00 PM",
		"data":{
			"headers":{
				"symbol":"SYMBOL","companyName":"NAME","lastSalePrice":"LAST PRICE","netChange":"NET CHANGE","percentageChange":"% CHANGE","deltaIndicator":"DELTA","oneYearPercentage":"1 yr % CHANGE"
			},
			"rows":[
				{"oneYearPercentage":"-1.22%","symbol":"AAA","companyName":"Investment Managers Series Trust II AXS First Priority CLO Bon","lastSalePrice":"$24.6900","netChange":"-0.0150","percentageChange":"-0.06072874493927126%","deltaIndicator":"down"},
				{"oneYearPercentage":"-23.69%","symbol":"WIZ","companyName":"Merlyn.AI Bull-Rider Bear-Fighter ETF","lastSalePrice":"$26.4700","netChange":"0.4165","percentageChange":"1.5466023022651318%","deltaIndicator":"up"}
			]
		}
	},
	"message":null,
	"status":{"rCode":200,"bCodeMessage":null,"developerMessage":null}
}
*/

type NasdaqScreenerSymbol struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type NasdaqScreenerTable struct {
	Rows []NasdaqScreenerSymbol `json:"rows"`
}

type NasdaqScreenerData struct {
	Table        NasdaqScreenerTable `json:"table"`
	TotalRecords int                 `json:"totalRecords"`
}

type NasdaqScreener struct {
	Data NasdaqScreenerData `json:"data"`
}

type NasdaqScreenerEtfSymbol struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
}

type NasdaqScreenerEtfData struct {
	Rows []NasdaqScreenerEtfSymbol `json:"rows"`
}

type NasdaqScreenerEtfDataData struct {
	Data NasdaqScreenerEtfData `json:"data"`
}

type NasdaqScreenerEtf struct {
	Data NasdaqScreenerEtfDataData `json:"data"`
}

/*
{
    "data":{
        "symbol":"AAPL",
        "companyName":"Apple Inc. Common Stock",
        "stockType":"Common Stock",
        "exchange":"NASDAQ-GS",
        "isNasdaqListed":true,
        "isNasdaq100":true,
        "isHeld":false,
        "primaryData":{
            "lastSalePrice":"$146.43",
            "netChange":"-2.97",
            "percentageChange":"-1.99%",
            "deltaIndicator":"down",
            "lastTradeTimestamp":"Feb 24, 2023 10:19 AM ET",
            "isRealTime":true,
            "bidPrice":"$146.42",
            "askPrice":"$146.43",
            "bidSize":"300",
            "askSize":"156",
            "volume":"11,763,041"
        },
        "secondaryData":null,
        "marketStatus":"Open",
        "assetClass":"STOCKS",
        "keyStats":null,
        "notifications":[]},
        "message":null,
        "status":{
            "rCode":200,
            "bCodeMessage":null,
            "developerMessage":null
        }
    }
}
{
	"data":{
		"symbol":"SGOL",
		"companyName":"abrdn Physical Gold Shares ETF",
		"stockType":null,
		"exchange":"PSE",
		"isNasdaqListed":false,
		"isNasdaq100":false,
		"isHeld":false,
		"primaryData":{
			"lastSalePrice":"$17.34",
			"netChange":"-0.13",
			"percentageChange":"-0.74%",
			"deltaIndicator":"down",
			"lastTradeTimestamp":"Feb 24, 2023",
			"isRealTime":false,
			"bidPrice":"N/A",
			"askPrice":"N/A",
			"bidSize":"N/A",
			"askSize":"N/A",
			"volume":"1,776,935"},
		"secondaryData":null,
		"marketStatus":"Closed",
		"assetClass":"ETF",
		"keyStats":null,
		"notifications":[]
	},
	"message":null,
	"status":{
		"rCode":200,
		"bCodeMessage":null,
		"developerMessage":null
	}
}
*/

type NasdaqInfoData struct {
	Symbol         string `json:"symbol"`
	CompanyName    string `json:"companyName"`
	StockType      string `json:"stockType"`
	AssetClass     string `json:"assetClass"`
	Exchange       string `json:"exchange"`
	IsNasdaqListed bool   `json:"isNasdaqListed"`
	IsNasdaq100    bool   `json:"isNasdaq100"`
}

type NasdaqInfo struct {
	Data NasdaqInfoData `json:"data"`
}

/*
{
    "data":{
        "ModuleTitle":{"label":"Module Title","value":"Company Description"},
        "CompanyName":{"label":"Company Name","value":"Apple Inc."},
        "Symbol":{"label":"Symbol","value":"AAPL"},
        "Address":{"label":"Address","value":"ONE APPLE PARK WAY, CUPERTINO, California, 95014, United States"},
        "Phone":{"label":"Phone","value":"+1 408 996-1010"},
        "Industry":{"label":"Industry","value":"Computer Manufacturing"},
        "Sector":{"label":"Sector","value":"Technology"},
        "Region":{"label":"Region","value":"North America"},
        "CompanyDescription":{"label":"Company Description","value":"Apple designs a wide variety of ..."},
        "CompanyUrl":{"label":"Company Url","value":"https://www.apple.com"},
        "KeyExecutives":{"label":"Key Executives","value":[{"name":"Jeffrey E. Williams","title":"Chief Operating Officer"}]}
    },
    "message":null,
    "status":{
        "rCode":200,
        "bCodeMessage":null,
        "developerMessage":null
    }
}
*/

type NasdaqLabelValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type NasdaqProfileData struct {
	Symbol             NasdaqLabelValue `json:"Symbol"`
	CompanyName        NasdaqLabelValue `json:"CompanyName"`
	CompanyDescription NasdaqLabelValue `json:"CompanyDescription"`
	CompanyUrl         NasdaqLabelValue `json:"CompanyUrl"`
	Industry           NasdaqLabelValue `json:"Industry"`
	Sector             NasdaqLabelValue `json:"Sector"`
	Region             NasdaqLabelValue `json:"Region"`
}

type NasdaqProfile struct {
	Data NasdaqProfileData `json:"data"`
}

/*
{
    "data":{
        "symbol":"ADBE",
        "summaryData":{
            "Exchange":{"label":"Exchange","value":"NASDAQ-GS"},
            "Sector":{"label":"Sector","value":"Technology"},
            "Industry":{"label":"Industry","value":"Computer Software: Prepackaged Software"},
            "OneYrTarget":{"label":"1 Year Target","value":"$387.50"},
            "TodayHighLow":{"label":"Today's High/Low","value":"N/A"},
            "ShareVolume":{"label":"Share Volume","value":"2,918"},
            "AverageVolume":{"label":"Average Volume","value":"2,772,771"},
            "PreviousClose":{"label":"Previous Close","value":"$347.02"},
            "FiftTwoWeekHighLow":{"label":"52 Week High/Low","value":"$479.21/$274.73"},
            "MarketCap":{"label":"Market Cap","value":"153,848,268,000"},
            "PERatio":{"label":"P/E Ratio","value":34.36},
            "ForwardPE1Yr":{"label":"Forward P/E 1 Yr.","value":"28.56"},
            "EarningsPerShare":{"label":"Earnings Per Share(EPS)","value":"$10.10"},
            "AnnualizedDividend":{"label":"Annualized Dividend","value":"N/A"},
            "ExDividendDate":{"label":"Ex Dividend Date","value":"N/A"},
            "DividendPaymentDate":{"label":"Dividend Pay Date","value":"N/A"},
            "Yield":{"label":"Current Yield","value":"N/A"},
            "Beta":{"label":"Beta","value":1.0}
		},
        "assetClass":"STOCKS",
        "additionalData":null,
        "bidAsk":{
            "Bid * Size":{"label":"Bid * Size","value":"$335.80 * 1"},
            "Ask * Size":{"label":"Ask * Size","value":"$337.00 * 2"}
		}
    },
    "message":null,
    "status":{"rCode":200,"bCodeMessage":null,"developerMessage":null}
}
{
	"data":{
		"symbol":"SGOL",
		"summaryData":{
			"TodayHighLow":{"label":"Today's High/Low","value":"$17.39/$17.324"},
			"ShareVolume":{"label":"Share Volume","value":"1,776,935"},
			"FiftyDayAvgDailyVol":{"label":"50 Day Avg. Daily Volume","value":"1,325,509"},
			"PreviousClose":{"label":"Previous Close","value":"$17.47"},
			"FiftTwoWeekHighLow":{"label":"52 Week High/Low","value":"$19.86/$15.5"},
			"MarketCap":{"label":"Market Cap","value":"150,858,000"},
			"AnnualizedDividend":{"label":"Annualized Dividend","value":"N/A"},
			"ExDividendDate":{"label":"Ex Dividend Date","value":"N/A"},
			"DividendPaymentDate":{"label":"Dividend Pay Date","value":"N/A"},
			"Yield":{"label":"Current Yield","value":"N/A"},
			"Alpha":{"label":"Alpha","value":"2.76"},
			"WeightedAlpha":{"label":"Weighted Alpha","value":"-1.70"},
			"Beta":{"label":"Beta","value":0.14},
			"StandardDeviation":{"label":"Standard Deviation","value":"unch"},
			"AvgDailyVol20Days":{"label":"Average Daily Volume 20 Days","value":"2,381,855"},
			"AvgDailyVol65Days":{"label":"Average Daily Volume 65 Days","value":"2,195,469"},
			"AUM":{"label":"Assets Under Management (,000)","value":"2,427,251"},
			"ExpenseRatio":{"label":"Expense Ratio","value":"0.17%"}
		},
		"assetClass":"ETF",
		"additionalData":null,
		"bidAsk":{
			"Bid * Size":{"label":"Bid * Size","value":"N/A"},
			"Ask * Size":{"label":"Ask * Size","value":"N/A"}
		}
	},
	"message":null,
	"status":{
		"rCode":200,
		"bCodeMessage":null,
		"developerMessage":null
	}
}
*/

type NasdaqSummaryDataSummaryData struct {
	Exchange     NasdaqLabelValue `json:"Exchange"`
	Industry     NasdaqLabelValue `json:"Industry"`
	Sector       NasdaqLabelValue `json:"Sector"`
	ExpenseRatio NasdaqLabelValue `json:"ExpenseRatio"`
}

type NasdaqSummaryData struct {
	Symbol      string                       `json:"symbol"`
	SummaryData NasdaqSummaryDataSummaryData `json:"summaryData"`
	AssetClass  string                       `json:"assetClass"`
}

type NasdaqSummary struct {
	Data NasdaqSummaryData `json:"data"`
}

const urlScreenerEtf = "https://api.nasdaq.com/api/screener/etf?tableonly=true&limit=999999&offset=0&download=true"
const urlScreenerStocks = "https://api.nasdaq.com/api/screener/stocks?tableonly=true&limit=999999&offset=0&download=false"
const infoUrlPrefix = "https://api.nasdaq.com/api/quote/"
const infoUrlSuffixStocks = "/info?assetclass=stocks"
const infoUrlSuffixEtf = "/info?assetclass=etf"
const summaryUrlPrefix = "https://api.nasdaq.com/api/quote/"
const summaryUrlSuffixStocks = "/summary?assetclass=stocks"
const summaryUrlSuffixEtf = "/summary?assetclass=etf"
const profileUrlPrefix = "https://api.nasdaq.com/api/company/"
const profileUrlSuffix = "/company-profile"
const ExchangeAll = ""
const ExchangeNasdaq = "&exchange=NASDAQ"
const ExchangeNyse = "&exchange=NYSE"
const ExchangeAmex = "&exchange=AMEX"

func RetrieveSymbolsStock(exchange string) ([]NasdaqSymbol, error) {
	syms := make([]NasdaqSymbol, 0)

	url := urlScreenerStocks + exchange
	b, err := get(url)
	if err != nil {
		return syms, fmt.Errorf("cannot get screener url '%s': %w", url, err)
	}

	scr := NasdaqScreener{}
	err = json.Unmarshal(b, &scr)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal NasdaqScreener: %w", err)
	}

	fmt.Printf("retrieved %d symbols\n", scr.Data.TotalRecords)
	for _, s := range scr.Data.Table.Rows {
		// "BRK/A" => "BRK-A", "BRK/B" => "BRK-B" ...
		s.Symbol = strings.ReplaceAll(s.Symbol, "/", "-")
		s.Symbol = strings.TrimSpace(s.Symbol)

		url = infoUrlPrefix + s.Symbol + infoUrlSuffixStocks
		b, err = get(url)
		if err != nil {
			e := fmt.Errorf("'%s' skipped: cannot get info url '%s': %w", s.Symbol, url, err)
			fmt.Println(e)
			continue
		}

		inf := NasdaqInfo{}
		if err := json.Unmarshal(b, &inf); err != nil {
			e := fmt.Errorf("'%s' skipped: cannot unmarshal NasdaqInfoData: %w", s.Symbol, err)
			fmt.Println(e, "bytes:", string(b))
			continue
		}

		url = profileUrlPrefix + s.Symbol + profileUrlSuffix
		b, err = get(url)
		if err != nil {
			e := fmt.Errorf("'%s' skipped: cannot get company profile url '%s': %w", s.Symbol, url, err)
			fmt.Println(e)
			continue
		}

		prof := NasdaqProfile{}
		if err := json.Unmarshal(b, &prof); err != nil {
			e := fmt.Errorf("'%s' skipped: cannot unmarshal NasdaqProfileData: %w", s.Symbol, err)
			fmt.Println(e, "bytes:", string(b))
			continue
		}

		exch := strings.TrimSpace(inf.Data.Exchange)
		ind := strings.TrimSpace(prof.Data.Industry.Value)
		sec := strings.TrimSpace(prof.Data.Sector.Value)
		asc := strings.TrimSpace(inf.Data.AssetClass)
		if len(exch) < 1 || len(ind) < 1 || len(sec) < 1 || len(asc) < 1 {
			url = summaryUrlPrefix + s.Symbol + summaryUrlSuffixStocks
			b, err = get(url)
			if err != nil {
				e := fmt.Errorf("'%s': cannot get summary url '%s': %w", s.Symbol, url, err)
				fmt.Println(e)
				goto skipped
			}

			summ := NasdaqSummary{}
			if err := json.Unmarshal(b, &summ); err != nil {
				e := fmt.Errorf("'%s': cannot unmarshal NasdaqSummary: %w", s.Symbol, err)
				fmt.Println(e, "bytes:", string(b))
				goto skipped
			}

			if len(exch) < 1 {
				exch = strings.TrimSpace(summ.Data.SummaryData.Exchange.Value)
			}

			if len(ind) < 1 {
				ind = strings.TrimSpace(summ.Data.SummaryData.Industry.Value)
			}

			if len(sec) < 1 {
				sec = strings.TrimSpace(summ.Data.SummaryData.Sector.Value)
			}

			if len(asc) < 1 {
				asc = strings.TrimSpace(summ.Data.AssetClass)
			}
		}

	skipped:
		sym := NasdaqSymbol{
			Mnemonic:    s.Symbol,
			Name:        strings.TrimSpace(inf.Data.CompanyName),
			Company:     strings.TrimSpace(prof.Data.CompanyName.Value),
			CompanyUrl:  strings.TrimSpace(prof.Data.CompanyUrl.Value),
			Description: strings.TrimSpace(prof.Data.CompanyDescription.Value),
			Currency:    "USD",
			Mic:         exchangeToMic(s.Symbol, exch),
			Exchange:    exch,
			IsListed:    inf.Data.IsNasdaqListed,
			IsNasdaq100: inf.Data.IsNasdaq100,
			StockType:   strings.TrimSpace(inf.Data.StockType),
			AssetClass:  asc,
			Industry:    ind,
			Sector:      sec,
			Region:      strings.TrimSpace(prof.Data.Region.Value),
		}

		syms = append(syms, sym)

		missing := empty
		if len(sym.Name) < 1 {
			missing = missing + " name,"
		}
		if len(sym.Company) < 1 {
			missing = missing + " company,"
		}
		/*
			if len(sym.CompanyUrl) < 1 {
				missing = missing + " company url,"
			}
		*/
		if len(sym.Description) < 1 {
			missing = missing + " description,"
		}
		if len(sym.Exchange) < 1 {
			missing = missing + " exchange,"
		}
		if len(sym.StockType) < 1 {
			sym.StockType = "STOCKS"
		}
		if len(sym.AssetClass) < 1 {
			missing = missing + " asset class,"
		}
		if len(sym.Industry) < 1 {
			missing = missing + " industry,"
		}
		if len(sym.Sector) < 1 {
			missing = missing + " sector,"
		}
		if len(sym.Region) < 1 {
			missing = missing + " region,"
		}

		if len(missing) > 0 {
			fmt.Println(sym.Mnemonic + ": empty" + missing)
		}
	}

	return syms, nil
}

func RetrieveSymbolsEtf() ([]NasdaqSymbol, error) {
	syms := make([]NasdaqSymbol, 0)

	url := urlScreenerEtf
	b, err := get(url)
	if err != nil {
		return syms, fmt.Errorf("cannot get screener url '%s': %w", url, err)
	}

	scr := NasdaqScreenerEtf{}
	err = json.Unmarshal(b, &scr)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal NasdaqScreenerEtf: %w", err)
	}

	fmt.Printf("retrieved %d symbols\n", len(scr.Data.Data.Rows))
	for _, s := range scr.Data.Data.Rows {
		// "BRK/A" => "BRK-A", "BRK/B" => "BRK-B" ...
		s.Symbol = strings.ReplaceAll(s.Symbol, "/", "-")
		s.Symbol = strings.TrimSpace(s.Symbol)

		url = infoUrlPrefix + s.Symbol + infoUrlSuffixEtf
		b, err = get(url)
		if err != nil {
			e := fmt.Errorf("'%s' skipped: cannot get info url '%s': %w", s.Symbol, url, err)
			fmt.Println(e)
			continue
		}

		inf := NasdaqInfo{}
		if err := json.Unmarshal(b, &inf); err != nil {
			e := fmt.Errorf("'%s' skipped: cannot unmarshal NasdaqInfoData: %w", s.Symbol, err)
			fmt.Println(e, "bytes:", string(b))
			continue
		}

		url = summaryUrlPrefix + s.Symbol + summaryUrlSuffixEtf
		b, err = get(url)
		if err != nil {
			e := fmt.Errorf("'%s' skipped: cannot get summary url '%s': %w", s.Symbol, url, err)
			fmt.Println(e)
			continue
		}

		summ := NasdaqSummary{}
		if err := json.Unmarshal(b, &summ); err != nil {
			e := fmt.Errorf("'%s' skipped: cannot unmarshal NasdaqSummary: %w", s.Symbol, err)
			fmt.Println(e, "bytes:", string(b))
			continue
		}

		exch := strings.TrimSpace(inf.Data.Exchange)
		if len(exch) < 1 {
			exch = strings.TrimSpace(summ.Data.SummaryData.Exchange.Value)
		}

		asc := strings.TrimSpace(inf.Data.AssetClass)
		if len(asc) < 1 {
			asc = strings.TrimSpace(summ.Data.AssetClass)
		}

		nam := strings.TrimSpace(inf.Data.CompanyName)
		if len(nam) < 1 {
			nam = s.CompanyName
		}

		sym := NasdaqSymbol{
			Mnemonic:     s.Symbol,
			Name:         nam,
			Currency:     "USD",
			Mic:          exchangeToMic(s.Symbol, exch),
			Exchange:     exch,
			IsListed:     inf.Data.IsNasdaqListed,
			IsNasdaq100:  inf.Data.IsNasdaq100,
			AssetClass:   asc,
			ExpenseRatio: summ.Data.SummaryData.ExpenseRatio.Value,
		}

		syms = append(syms, sym)

		missing := empty
		if len(sym.Name) < 1 {
			missing = missing + " name,"
		}
		if len(sym.Exchange) < 1 {
			missing = missing + " exchange,"
		}
		if len(sym.AssetClass) < 1 {
			sym.AssetClass = "ETF"
		}
		if len(sym.ExpenseRatio) < 1 {
			missing = missing + " expense ratio,"
		}

		if len(missing) > 0 {
			fmt.Println(sym.Mnemonic + ": empty" + missing)
		}
	}

	return syms, nil
}

func exchangeToMic(symbol, exchange string) string {
	switch exchange {
	case "NASDAQ-GS":
		return "XNGS" // segment of XNAS
	case "NASDAQ-GM":
		return "XNGM" // segment of XNAS
	case "NASDAQ-CM":
		return "XNCM" // segment of XNAS
	case "NYSE":
		return "XNYS"
	case "AMEX":
		return "AMXO" // segment of XNYS
	case "PSE":
		return "PSE" // PSE or XPHS
	case "BATS":
		return "BATS" // segment of XCBO - CBOE GLOBAL MARKETS
	default:
		fmt.Println(symbol + ": unknown mic for exchange: '" + exchange + "', setting to XXXX")
		return "XXXX"
	}
}
