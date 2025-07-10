package intraday

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type StringOrInt string

func (t *StringOrInt) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*t = StringOrInt(s)
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*t = StringOrInt(fmt.Sprintf("%d", i))
		return nil
	}

	*t = StringOrInt("?")
	return nil
}

type TradeType struct {
	Code  string `json:"code"`
	IDNXT string `json:"idNXT"`
	Label string `json:"label"`
}

type Trade struct {
	TradeID StringOrInt `json:"tradeId"`
	Time    string      `json:"time"`
	Price   string      `json:"price"`
	Volume  string      `json:"volume"`
	Type    string      `json:"type"`
}

type JsonIntraday struct {
	Rows             []Trade     `json:"rows"`
	Count            int         `json:"count"`
	Date             string      `json:"date"`
	CountFiltered    int         `json:"countFiltered"`
	Startdate        string      `json:"startdate"`
	TradeTypeList    []TradeType `json:"tradeTypeList"`
	TimeZone         string      `json:"timeZone"`
	SliderTimeFilter struct {
		Time struct {
			Start        string `json:"start"`
			End          string `json:"end"`
			StartMinutes int    `json:"startMinutes"`
			EndMinutes   int    `json:"endMinutes"`
		} `json:"Time"`
	} `json:"sliderTimeFilter"`
	SliderFilters struct {
		Price struct {
			Min      string  `json:"min"`
			MinLimit string  `json:"minLimit"`
			Max      string  `json:"max"`
			MaxLimit float64 `json:"maxLimit"`
		} `json:"Price"`
		Shares struct {
			Min StringOrInt `json:"min"`
			Max StringOrInt `json:"max"`
		} `json:"Shares"`
	} `json:"sliderFilters"`
}

func ReadJsonIntradayFile(fileName string) (*JsonIntraday, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open json file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read json file: %w", err)
	}

	jsonInd := JsonIntraday{}
	err = json.Unmarshal(byteValue, &jsonInd)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	return &jsonInd, nil
}

func WriteJsonIntradayFile(fileName string, jsonEod *JsonIntraday) error {
	jsonData, err := json.MarshalIndent(jsonEod, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal json: %w", err)
	}

	jsonFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil
}

func NewTradeLabelMap() map[string]string {
	return map[string]string{
		strings.ToLower("Official opening index"):      "OOI", // code:"IA", idNXT:"IA"
		strings.ToLower("Real-time index"):             "RTI", // code:"I2", idNXT:"I2"
		strings.ToLower("Automatic indicative index"):  "AII", // code:"I3", idNXT:"I3"
		strings.ToLower("Options liquidation index"):   "OLI", // code:"I7", idNXT:"I7"
		strings.ToLower("Closing Reference index"):     "CRI", // code:"I5", idNXT:"I5"
		strings.ToLower("Preliminary Reference index"): "PRI", // code:"IB", idNXT:"IB"
		strings.ToLower("Confirmed Reference index"):   "FRI", // code:"IC", idNXT:"IC"
		strings.ToLower("Exchange Continuous"):         "ECO", // code:"00", idNXT:"240:0"
		strings.ToLower("Auction"):                     "AUC", // code:"00H", idNXT:"00H"
		strings.ToLower("Valuation Trade"):             "VAL", // code:"VT", idNXT:"VT"
		strings.ToLower("Retail Matching Facility"):    "RMF", // code:"0R", idNXT:"240:R"
		strings.ToLower("OffBook Out of market"):       "OBM", // code:"2H", idNXT:"242:H"
		strings.ToLower("OffBook Investment funds"):    "OBF", // code:"2I", idNXT:"242:I"
		strings.ToLower("OffBook On Exchange"):         "OBE", // code:"525", idNXT:"525"
		strings.ToLower("Trading at last"):             "TAL", // code:"00L", idNXT:"00L"
		strings.ToLower("Trade Cancellation"):          "TCA", // code:"24", idNXT:"24"
		strings.ToLower("Dark Trade"):                  "DKT", // code:"33", idNXT:"33"
		strings.ToLower("Request for Quote"):           "RFQ", // code:"104", idNXT:"104"
		strings.ToLower("Opening"):                     "OPN", // code:"0OP", idNXT:"240:O"
		strings.ToLower("Exchange Cross"):              "ECR", // code:"01", idNXT:"40:1"
		"":                                             "UNK", // unknown trade type
		// Add more mappings as needed
	}
}
