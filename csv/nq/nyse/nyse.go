package nyse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

/*
https://data2-widgets.dataservices.theice.com/fsml?requestType=content&username=nysecomwebsite&key=oHhwWp17SzK9d77UJcnVMG6YGEAxxpjGr7K6x5VF48gmm8VMhYItfTYw%2FjtC1pWsKWOhDAZdafL%2FVTfPQ5yx5kitxtECJdIbgNnI2LkNYns%3D&cbid=7010&dataset=MQ_Fundamentals&fsmlParams=key%3DNVDA&json=true&callback=__gwt_jsonp__.P2.onSuccess
callback(
{
	"success":"true",
	"content":{
		"retrieved":"2023-02-28T15:22:14Z",
		"MQ_Fundamentals":{
			"REVQ_3":8288,
			"REVQ_4":7643,
			"EPSQDate_10":"",
			"REVQ_1":5931,
			"REVQ_2":6704,
			"REVQ_7":5661,
			"REVQ_8":5003,
			"REVQ_5":7103,
			"REVQ_6":6507,
			"REVQ":6051,
			"TotalDebt2EquityTTM":54.436451,
			"REVQ_9":4726,
			"REVQ_15":"",
			"EPSQDate_15":"",
			"ShortVolDate":"2023-02-27 00:00:00",
			"REVQ_14":"",
			"REVQ_13":"",
			"EPSQ_6":".952300",
			"EPSQDate_13":"",
			"REVQ_12":"",
			"EPSQ_7":".769700",
			"EPSQDate_14":"",
			"REVQ_11":"",
			"EPSQ_8":".588400",
			"EPSQDate_9":20201118,
			"EPSQDate_11":"",
			"REVQ_10":"",
			"EPSQ_9":".540500",
			"EPSQDate_8":20210224,
			"EPSQDate_12":"",
			"CurrentRatioTTM":3.515618,
			"EPSQ_14":"",
			"EPSQDate_7":20210526,
			"EPSQ_13":"",
			"EPSQDate_6":20210818,
			"Price2SalesTTM":21.6423,
			"EPSQDate_5":20211117,
			"LASTUPDATE":"2023-02-28 00:42:29",
			"EPSQ_15":"",
			"EPSQDate_4":20220216,
			"TotalRevenueFY":26974,
			"EPSQDate_3":20220525,
			"Price2BookTTM":25.9822,
			"EPSQDate_2":20220824,
			"EPSQDate_1":20221116,
			"expiry":3600,
			"EPSQ_10":"",
			"EPSQ_12":"",
			"EPSQ_11":"",
			"ShortVol":2717259,
			"ReturnOnAssetsTTM":10.233223,
			"PayoutRatioTTM":9.18801,
			"SEDOL":2379504,
			"status":"ok",
			"ReturnOnEquityTTM":17.933611,
			"Symbol":"NVDA",
			"EPSQDate":20230222,
			"QR1EPSEstimate":7.034131,
			"QR1ReportDate":"2023-07-31 00:00:00",
			"key":"NVDA",
			"fromcache":1,
			"EPSQ":".573900",
			"ISIN":"US67066G1040",
			"EPSQ_2":".262900",
			"EPSQ_3":".645600",
			"QuickRatioTTM":2.729544,
			"EPSQ_4":1.1993,
			"EBITDATTM":7794,
			"EPSQ_5":".986000",
			"EPSQ_1":".273900"
		}
	}
}
)
*/

type NyseSymbol struct {
	ISIN  string `json:"ISIN"`
	SEDOL string `json:"SEDOL"`
}

type NyseContentFundamentals struct {
	Fundamentals NyseSymbol `json:"MQ_Fundamentals"`
}

type NyseContent struct {
	Content NyseContentFundamentals `json:"content"`
}

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57"
const cookieName1 = "JSESSIONID"
const cookieName2 = "BIGipServerdata-widgets.c4z.dataservices.theice.com-8080"
const empty = ""

var cookieValue1 = empty
var cookieValue2 = empty

func unmarshalContent(jsn []byte) (*NyseContent, error) {
	// callback(JSON)
	jsn = bytes.TrimLeft(jsn, "callback(")
	jsn = bytes.TrimRight(jsn, ")")

	ct := NyseContent{}
	if err := json.Unmarshal(jsn, &ct); err != nil {
		return nil, fmt.Errorf("cannot unmarshal NyseContent: %w", err)
	}

	return &ct, nil
}

/*
	func GetCookie() error {
		//req, err := http.NewRequest("GET", "https://data1-widgets.dataservices.theice.com/fsml?requestType=content&username=nysecomwebsite&key=oHhwWp17SzK9d77UJcnVMG6YGEAxxpjGr7K6x5VF48gmm8VMhYItfTYw%2FjtC1pWsKWOhDAZdafL%2FVTfPQ5yx5kC6vSC%2Bs3slyVQAI9Snlko%3D&cbid=7010&dataset=MQ_Fundamentals&fsmlParams=key%3DNVDA&json=true&callback=__gwt_jsonp__.P2.onSuccess", nil)
		req, err := http.NewRequest("GET", "https://data2-widgets.dataservices.theice.com/fsml?requestType=content&username=nysecomwebsite&key=oHhwWp17SzK9d77UJcnVMG6YGEAxxpjGr7K6x5VF48gmm8VMhYItfTYw%2FjtC1pWsKWOhDAZdafL%2FVTfPQ5yx5g1uQwfEeM98wA7yxgCO5WI%3D&cbid=7010&dataset=MQ_Fundamentals&fsmlParams=key%3DNVDA&json=true", nil)
		if err != nil {
			cookieValue1 = empty
			cookieValue2 = empty
			return fmt.Errorf("cannot create request: %w", err)
		}

		req.Header.Set("Referer", "https://www.nyse.com/")
		req.Header.Set("User-Agent", userAgent)
		if cookieValue1 != empty {
			req.AddCookie(&http.Cookie{Name: cookieName1, Value: cookieValue1})
		}
		if cookieValue2 != empty {
			req.AddCookie(&http.Cookie{Name: cookieName2, Value: cookieValue2})
		}

		httpClient := http.Client{Timeout: time.Duration(30) * time.Second}
		resp, err := httpClient.Do(req)
		if err != nil {
			cookieValue1 = empty
			cookieValue2 = empty
			return fmt.Errorf("cannot do request: %w", err)
		}
		defer resp.Body.Close()

		for _, q := range resp.Cookies() {
			if q.Name == cookieName1 {
				cookieValue1 = q.Value
				fmt.Println("obtained cookie '" + cookieName1 + "': " + cookieValue1)
			}
			if q.Name == cookieName2 {
				cookieValue2 = q.Value
				fmt.Println("obtained cookie '" + cookieName2 + "': " + cookieValue2)
			}
		}

		body, err := io.ReadAll(resp.Body)
		return fmt.Errorf("err is: %w, body is: %s", err, string(body))
	}
*/
func GetSymbol(symbol string) (*NyseSymbol, error) /*(*NyseContent, error)*/ {
	const urlPrefix = "https://data2-widgets.dataservices.theice.com/fsml?requestType=content&username=nysecomwebsite&key=oHhwWp17SzK9d77UJcnVMG6YGEAxxpjGr7K6x5VF48gmm8VMhYItfTYw%2FjtC1pWsKWOhDAZdafL%2FVTfPQ5yx5hgr%2FsUG8MiD%2BfRBLLZR2ms%3D&cbid=7010&dataset=MQ_Fundamentals&fsmlParams=key%3D"
	const urlSuffix = "&json=true"
	url := urlPrefix + symbol + urlSuffix
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("Referer", "https://www.nyse.com/")
	req.Header.Set("User-Agent", userAgent)
	if cookieValue1 != empty {
		req.AddCookie(&http.Cookie{Name: cookieName1, Value: cookieValue1})
	}
	if cookieValue2 != empty {
		req.AddCookie(&http.Cookie{Name: cookieName2, Value: cookieValue2})
	}

	httpClient := http.Client{Timeout: time.Duration(30) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot do request: %w", err)
	}
	defer resp.Body.Close()

	for _, q := range resp.Cookies() {
		if q.Name == cookieName1 {
			cookieValue1 = q.Value
			fmt.Println("obtained cookie '" + cookieName1 + "': " + cookieValue1)
		}
		if q.Name == cookieName2 {
			cookieValue2 = q.Value
			fmt.Println("obtained cookie '" + cookieName2 + "': " + cookieValue2)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	str := string(body)
	sp1 := strings.Split(str, "\"ISIN\":")
	if len(sp1) < 2 {
		return nil, fmt.Errorf("cannot find ISIN in body %s : %w", str, err)
	}

	sp2 := strings.Split(sp1[1], ",")
	isin := strings.Trim(sp2[0], "\"")

	sp1 = strings.Split(str, "\"SEDOL\":")
	if len(sp1) < 2 {
		return nil, fmt.Errorf("cannot find SEDOL in body %s : %w", str, err)
	}

	sp2 = strings.Split(sp1[1], ",")
	sedol := strings.Trim(sp2[0], "\"")

	nc := &NyseSymbol{
		ISIN:  isin,
		SEDOL: sedol,
	}

	return nc, nil
}

/*
https://nyse.widgets.dataservices.theice.com/Login?auth=Q%2FCRa%2B4NFi2BfASmdlA%2BRF2yCe4GRjLRi%2Fn%2FFu2k8R%2B6UWZXeOd4A6LiEnqoSeff9I%2F2H66RCe3CEFVMBUiNvLFep22K4K%2FSaJoXk7FIUspAxMz5GwIuU5QgSoPm0eHZRe2g1QxkqBtyiRbETtLlrEHGqTc6F%2FqslO53SJvHkWE%3D&browser=false&client=mobile&callback=__gwt_jsonp__.P0.onSuccess
https://nyse.widgets.dataservices.theice.com/Login?auth=Q%2FCRa%2B4NFi2BfASmdlA%2BRF2yCe4GRjLRi%2Fn%2FFu2k8R%2B6UWZXeOd4A6LiEnqoSeff9I%2F2H66RCe3CEFVMBUiNvLFep22K4K%2FSaJoXk7FIUspAxMz5GwIuU5QgSoPm0eHZoAAN8N22gal3IsoKueO88xIsdOIdplK9TYrrAI%2B6pu4%3D&browser=false&client=mobile
__gwt_jsonp__.P0.onSuccess({
	"success":"true",
	"backupwebserver":"https://data1-widgets.dataservices.theice.com/",
	"cbid":"7010",
	"rememberid":"506c8bc431d3d4500c3ca1edbf3477bd",
	"dataserver":{
		"conn.dflt.cm":"none",
		"conn.http.port":"80",
		"conn.alt.url":"none",
		"conn.attempts":"2",
		"conn.type":"-10",
		"conn.dflt.port":"443",
		"conn.ssl.port":"443",
		"conn.dflt.url":"wss://stream2-300.dataservices.theice.com",
		"auth.keyp":"false",
		"conn.alt.cm":"none",
		"conn.ssl.url":"none",
		"snap.url":"snapshot?symbol={symbol}&type={type}&username={username}&key={key}&cbid={cbid}",
		"auth.key":"oHhwWp17SzK9d77UJcnVMG6YGEAxxpjGr7K6x5VF48gmm8VMhYItfTYw/jtC1pWsKWOhDAZdafL/VTfPQ5yx5qJogskvK/kBz3BM4j1zMG8=",
		"conn.http.cm":"none",
		"conn.ssl.cm":"none",
		"conn.http.url":"none",
		"conn.alt.port":"80",
		"auth.username":"nysecomwebsite"
	},
	"webserver":"https://data2-widgets.dataservices.theice.com/",
	"authentication":"J/ni1g+xJ/72Up41H5J8sWGM82QfitXM93gBIRQamTEz3L5oVa6BxS4k0PoyYxAeX9UvUfhfQnwuJjKYJfTdgQ==",
	"webserversession":"nysecomwebsite,oHhwWp17SzK9d77UJcnVMG6YGEAxxpjGr7K6x5VF48gmm8VMhYItfTYw/jtC1pWsKWOhDAZdafL/VTfPQ5yx5qJogskvK/kBz3BM4j1zMG8=,549cc549132d41006c9a1c76576e5a7b"
})
*/

/*
func get(targetURL string) ([]byte, error) {
repeat:

		if cookieValue == empty {
			if err := getCookie(); err != nil {
				return nil, fmt.Errorf("cannot obtain access cookie: %w", err)
			}
		}

		req, err := http.NewRequest("GET", targetURL, nil)
		if err != nil {
			return nil, fmt.Errorf("cannot create request: %w", err)
		}

		req.Header.Set("User-Agent", userAgent)
		req.AddCookie(&http.Cookie{Name: cookieName, Value: cookieValue})

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
*/
func RetrieveSymbolInfo(mnemonic string) (NyseSymbol, error) {
	//url := "https://api.nasdaq.com/api/quote/" + mnemonic + "/realtime-trades?limit=99999999&fromTime="
	return NyseSymbol{}, nil
}
