package euronext

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type JsonInd struct {
	Rows []struct {
		TradeID int    `json:"tradeId"`
		Time    string `json:"time"`
		Price   string `json:"price"`
		Volume  string `json:"volume"`
		Type    string `json:"type"`
	} `json:"rows"`
	Count         int    `json:"count"`
	Date          string `json:"date"`
	CountFiltered int    `json:"countFiltered"`
	Startdate     string `json:"startdate"`
	TradeTypeList []struct {
		Code  string `json:"code"`
		IDNXT string `json:"idNXT"`
		Label string `json:"label"`
	} `json:"tradeTypeList"`
	TimeZone         string `json:"timeZone"`
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
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"Shares"`
	} `json:"sliderFilters"`
}

func ReadJsonIndFile(fileName string) (*JsonInd, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open json file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read json file: %w", err)
	}

	jsonInd := JsonInd{}
	err = json.Unmarshal(byteValue, &jsonInd)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	return &jsonInd, nil
}

func WriteJsonIndFile(fileName string, jsonEod *JsonInd) error {
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
