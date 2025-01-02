package euronext

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// JSON format starts on 2012-09-07
// Filas are prefixed with MEP {AMS, BRU, LIS, PAR, OTH}
// From 2013-01-07 files are prefixed with MIC

type JsonEod struct {
	Data []struct {
		Isin           string `json:"ISIN"`
		Market         string `json:"MIC"`
		Date           string `json:"date"`
		Open           string `json:"open"`
		High           string `json:"high"`
		Low            string `json:"low"`
		Close          string `json:"close"`
		Nymberofshares string `json:"nymberofshares"`
		Numoftrades    string `json:"numoftrades"`
		Turnover       string `json:"turnover"`
		Currency       string `json:"currency"`
	} `json:"data"`
}

// 20190706:
// "ALXA" also "Euronext Amsterdam" ????
// "ALXP" also "null" ????
// "MLXB" also "null" ????
// "TNLA" also "null" ????
// "TNLB" also "null" ????
// "XBRU" also "null" ????

var jsonEodMarketToMic = map[string]string{
	"Euronext Amsterdam":          "XAMS",
	"Euronext Brussels":           "XBRU",
	"Euronext Lisbon":             "XLIS",
	"Euronext Paris":              "XPAR",
	"NYSE Euronext Amsterdam":     "XAMS",
	"NYSE Euronext Brussels":      "XBRU",
	"NYSE Euronext Lisbon":        "XLIS",
	"NYSE Euronext Paris":         "XPAR",
	"Euronext Growth Amsterdam":   "ALXA",
	"Euronext Growth Brussels":    "ALXB",
	"Euronext Growth Lisbon":      "ALXL",
	"Euronext Growth Paris":       "ALXP",
	"Euronext Access Lisbon":      "ENXL",
	"Euronext Access Brussels":    "MLXB",
	"Euronext Access Paris":       "XMLI",
	"Traded not listed Amsterdam": "TNLA",
	"Traded not listed Brussels":  "TNLB",
	"???":                         "XHFT",
	"XPAR":                        "XPAR",
}

func (je *JsonEod) MarketToMic() (string, error) {
	if len(je.Data) == 0 {
		return "", fmt.Errorf("no data found")
	}

	market := je.Data[0].Market
	mic, exists := jsonEodMarketToMic[market]
	if !exists {
		return "", fmt.Errorf("unknown market '%s'", market)
	}

	return mic, nil
}

func ReadJsonEodFile(fileName string) (*JsonEod, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open json file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read json file: %w", err)
	}

	jsonEod := JsonEod{}
	err = json.Unmarshal(byteValue, &jsonEod)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	return &jsonEod, nil
}

func WriteJsonEodFile(fileName string, jsonEod *JsonEod) error {
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
