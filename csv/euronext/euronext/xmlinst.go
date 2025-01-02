package euronext

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type XmlIcb struct {
	Icb1 string `xml:"icb1,attr" json:"icb1"`
	Icb2 string `xml:"icb2,attr" json:"icb2"`
	Icb3 string `xml:"icb3,attr" json:"icb3"`
	Icb4 string `xml:"icb4,attr" json:"icb4"`
}

type XmlStock struct {
	Cfi         string  `xml:"cfi,attr" json:"cfi"`
	Compartment string  `xml:"compartment,attr" json:"compartment"`
	Currency    string  `xml:"currency,attr" json:"currency"`
	Shares      string  `xml:"shares,attr" json:"shares"`
	TradingMode string  `xml:"tradingMode,attr" json:"tradingMode"`
	Icb         *XmlIcb `xml:"icb" json:"icb"`
}

type XmlIndex struct {
	BaseCap         *string `xml:"baseCap,attr" json:"baseCap"`
	BaseCapCurrency *string `xml:"baseCapCurrency,attr" json:"baseCapCurrency"`
	BaseDate        string  `xml:"baseDate,attr" json:"baseDate"`
	BaseLevel       string  `xml:"baseLevel,attr" json:"baseLevel"`
	CalcFreq        string  `xml:"calcFreq,attr" json:"calcFreq"`
	CapFactor       string  `xml:"capFactor,attr" json:"capFactor"`
	Currency        string  `xml:"currency,attr" json:"currency"`
	Family          string  `xml:"family,attr" json:"family"`
	Kind            string  `xml:"kind,attr" json:"kind"`
	Weighting       string  `xml:"weighting,attr" json:"weighting"`
	Icb             *XmlIcb `xml:"icb" json:"icb"`
}

type XmlTarget struct {
	Isin   string `xml:"isin,attr" json:"isin"`
	Mep    string `xml:"mep,attr" json:"mep"`
	Mic    string `xml:"mic,attr" json:"mic"`
	Name   string `xml:"name,attr" json:"name"`
	Symbol string `xml:"symbol,attr" json:"symbol"`
	Vendor string `xml:"vendor,attr" json:"vendor"`
}

type XmlInav struct {
	Currency string      `xml:"currency,attr" json:"currency"`
	Target   []XmlTarget `xml:"target" json:"target"`
}

type XmlFund struct {
	Cfi         string `xml:"cfi,attr" json:"cfi"`
	Currency    string `xml:"currency,attr" json:"currency"`
	Issuer      string `xml:"issuer,attr" json:"issuer"`
	Shares      string `xml:"shares,attr" json:"shares"`
	TradingMode string `xml:"tradingMode,attr" json:"tradingMode"`
}

type XmlEtv struct {
	AllInFees         string `xml:"allInFees,attr" json:"allInFees"`
	Cfi               string `xml:"cfi,attr" json:"cfi"`
	Currency          string `xml:"currency,attr" json:"currency"`
	DividendFrequency string `xml:"dividendFrequency,attr" json:"dividendFrequency"`
	ExpenseRatio      string `xml:"expenseRatio,attr" json:"expenseRatio"`
	Issuer            string `xml:"issuer,attr" json:"issuer"`
	Shares            string `xml:"shares,attr" json:"shares"`
	TradingMode       string `xml:"tradingMode,attr" json:"tradingMode"`
}

type XmlUnderlying struct {
	Isin   string `xml:"isin,attr" json:"isin"`
	Mep    string `xml:"mep,attr" json:"mep"`
	Mic    string `xml:"mic,attr" json:"mic"`
	Name   string `xml:"name,attr" json:"name"`
	Symbol string `xml:"symbol,attr" json:"symbol"`
	Vendor string `xml:"vendor,attr" json:"vendor"`
}

type XmlEtf struct {
	Cfi               string        `xml:"cfi,attr" json:"cfi"`
	Currency          string        `xml:"currency,attr" json:"currency"`
	DividendFrequency string        `xml:"dividendFrequency,attr" json:"dividendFrequency"`
	ExpositionType    string        `xml:"expositionType,attr" json:"expositionType"`
	Fraction          string        `xml:"fraction,attr" json:"fraction"`
	IndexFamily       string        `xml:"indexFamily,attr" json:"indexFamily"`
	Issuer            string        `xml:"issuer,attr" json:"issuer"`
	LaunchDate        string        `xml:"launchDate,attr" json:"launchDate"`
	Mer               string        `xml:"mer,attr" json:"mer"`
	Ter               string        `xml:"ter,attr" json:"ter"`
	TradingMode       string        `xml:"tradingMode,attr" json:"tradingMode"`
	Inav              XmlInav       `xml:"inav" json:"inav"`
	Underlying        XmlUnderlying `xml:"underlying" json:"underlying"`
}

type XmlInstrument struct {
	Cfi           *string   `xml:"cfi,attr" json:"cfi,omitempty"`
	Description   *string   `xml:"description,attr" json:"description,omitempty"`
	File          string    `xml:"file,attr" json:"file"`
	FoundInSearch string    `xml:"foundInSearch,attr" json:"foundInSearch"`
	Isin          string    `xml:"isin,attr" json:"isin"`
	Mep           string    `xml:"mep,attr" json:"mep"`
	Mic           string    `xml:"mic,attr" json:"mic"`
	Name          string    `xml:"name,attr" json:"name"`
	Notes         *string   `xml:"notes,attr" json:"notes,omitempty"`
	Symbol        string    `xml:"symbol,attr" json:"symbol"`
	Tradingmode   *string   `xml:"tradingmode,attr" json:"tradingmode,omitempty"`
	Type          string    `xml:"type,attr" json:"type"`
	Vendor        *string   `xml:"vendor,attr" json:"vendor,omitempty"`
	Etf           *XmlEtf   `xml:"etf" json:"etf,omitempty"`
	Etv           *XmlEtv   `xml:"etv" json:"etv,omitempty"`
	Fund          *XmlFund  `xml:"fund" json:"fund,omitempty"`
	Inav          *XmlInav  `xml:"inav" json:"inav,omitempty"`
	Index         *XmlIndex `xml:"index" json:"index,omitempty"`
	Stock         *XmlStock `xml:"stock" json:"stock,omitempty"`
}

type XmlInstruments struct {
	Instrument []XmlInstrument `xml:"instrument" json:"instrument"`
}

func ReadXmlInstrumentsFile(fileName string) (*XmlInstruments, error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open xml file: %w", err)
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	xmlInstruments := XmlInstruments{}
	err = xml.Unmarshal(byteValue, &xmlInstruments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal xml file: %w", err)
	}

	return &xmlInstruments, nil
}

func ReadJsonInstrumentsFile(fileName string) (*XmlInstruments, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open json file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read json file: %w", err)
	}

	xmlInstruments := XmlInstruments{}
	err = json.Unmarshal(byteValue, &xmlInstruments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	return &xmlInstruments, nil
}

func WriteXmlInstrumentsFile(fileName string, xmlInstruments *XmlInstruments) error {
	xmlData, err := xml.MarshalIndent(xmlInstruments, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal xml: %w", err)
	}

	xmlFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer xmlFile.Close()

	_, err = xmlFile.Write(xmlData)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil
}

func WriteJsonInstrumentsFile(fileName string, xmlInstruments *XmlInstruments) error {
	jsonData, err := json.MarshalIndent(xmlInstruments, "", "  ")
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

/*
type XmlEuronextInstruments struct {
	XmlEuronextInstrument []struct {
		Cfi           *string `xml:"cfi,attr"`
		Description   *string `xml:"description,attr"`
		File          string  `xml:"file,attr"`
		FoundInSearch string  `xml:"foundInSearch,attr"`
		Isin          string  `xml:"isin,attr"`
		Mep           string  `xml:"mep,attr"`
		Mic           string  `xml:"mic,attr"`
		Name          string  `xml:"name,attr"`
		Notes         *string `xml:"notes,attr"`
		Symbol        string  `xml:"symbol,attr"`
		Tradingmode   *string `xml:"tradingmode,attr"`
		Type          string  `xml:"type,attr"`
		Vendor        *string `xml:"vendor,attr"`
		Etf           *struct {
			Cfi               string   `xml:"cfi,attr"`
			Currency          string   `xml:"currency,attr"`
			DividendFrequency string   `xml:"dividendFrequency,attr"`
			ExpositionType    string   `xml:"expositionType,attr"`
			Fraction          string   `xml:"fraction,attr"`
			IndexFamily       string   `xml:"indexFamily,attr"`
			Issuer            string   `xml:"issuer,attr"`
			LaunchDate        string   `xml:"launchDate,attr"`
			Mer               *float64 `xml:"mer,attr"`
			Ter               string   `xml:"ter,attr"`
			TradingMode       string   `xml:"tradingMode,attr"`
			Inav              struct {
				Isin   string `xml:"isin,attr"`
				Mep    string `xml:"mep,attr"`
				Mic    string `xml:"mic,attr"`
				Name   string `xml:"name,attr"`
				Symbol string `xml:"symbol,attr"`
				Vendor string `xml:"vendor,attr"`
			} `xml:"inav"`
			Underlying struct {
				Isin   string `xml:"isin,attr"`
				Mep    string `xml:"mep,attr"`
				Mic    string `xml:"mic,attr"`
				Name   string `xml:"name,attr"`
				Symbol string `xml:"symbol,attr"`
				Vendor string `xml:"vendor,attr"`
			} `xml:"underlying"`
		} `xml:"etf"`
		Etv *struct {
			AllInFees         string `xml:"allInFees,attr"`
			Cfi               string `xml:"cfi,attr"`
			Currency          string `xml:"currency,attr"`
			DividendFrequency string `xml:"dividendFrequency,attr"`
			ExpenseRatio      string `xml:"expenseRatio,attr"`
			Issuer            string `xml:"issuer,attr"`
			Shares            string `xml:"shares,attr"`
			TradingMode       string `xml:"tradingMode,attr"`
		} `xml:"etv"`
		Fund *struct {
			Cfi         string `xml:"cfi,attr"`
			Currency    string `xml:"currency,attr"`
			Issuer      string `xml:"issuer,attr"`
			Shares      string `xml:"shares,attr"`
			TradingMode string `xml:"tradingMode,attr"`
		} `xml:"fund"`
		Inav *struct {
			Currency string `xml:"currency,attr"`
			Target   []struct {
				Isin   string `xml:"isin,attr"`
				Mep    string `xml:"mep,attr"`
				Mic    string `xml:"mic,attr"`
				Name   string `xml:"name,attr"`
				Symbol string `xml:"symbol,attr"`
				Vendor string `xml:"vendor,attr"`
			} `xml:"target"`
		} `xml:"inav"`
		Index *struct {
			BaseCap         *string `xml:"baseCap,attr"`
			BaseCapCurrency *string `xml:"baseCapCurrency,attr"`
			BaseDate        string  `xml:"baseDate,attr"`
			BaseLevel       string  `xml:"baseLevel,attr"`
			CalcFreq        string  `xml:"calcFreq,attr"`
			CapFactor       string  `xml:"capFactor,attr"`
			Currency        string  `xml:"currency,attr"`
			Family          string  `xml:"family,attr"`
			Kind            string  `xml:"kind,attr"`
			Weighting       string  `xml:"weighting,attr"`
			Icb             *struct {
				Icb1 int  `xml:"icb1,attr"`
				Icb2 *int `xml:"icb2,attr"`
				Icb3 *int `xml:"icb3,attr"`
			} `xml:"icb"`
		} `xml:"index"`
		Stock *struct {
			Cfi         string `xml:"cfi,attr"`
			Compartment string `xml:"compartment,attr"`
			Currency    string `xml:"currency,attr"`
			Shares      string `xml:"shares,attr"`
			TradingMode string `xml:"tradingMode,attr"`
			Icb         struct {
				Icb1 string `xml:"icb1,attr"`
				Icb2 string `xml:"icb2,attr"`
				Icb3 string `xml:"icb3,attr"`
				Icb4 string `xml:"icb4,attr"`
			} `xml:"icb"`
		} `xml:"stock"`
	} `xml:"instrument"`
}

func ReadXmlInstrumentsFile(fileName string) (*XmlEuronextInstruments, error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open xml file: %w", err)
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	euronextInstruments := XmlEuronextInstruments{}
	err = xml.Unmarshal(byteValue, &euronextInstruments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal xml file: %w", err)
	}

	return &euronextInstruments, nil
}
*/
