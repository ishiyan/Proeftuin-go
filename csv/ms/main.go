package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const configFileName = "ms.json"

type point struct {
	date  time.Time
	value float64
}

type index struct {
	Mnemonic                string `json:"mnemonic"`
	MorningstarID           string `json:"msid"`
	Name                    string `json:"name"`
	NameSeries              string `json:"series"`
	BaseCurrency            string `json:"currency"`
	IndexAssetClass         string `json:"assetClass"`
	ReturnType              string `json:"returnType"`
	WeightingScheme         string `json:"weightingScheme"`
	DateInception           string `json:"dateInception"`
	DateStartPerformance    string `json:"dateStartPerformance"`
	FrequencyRebalance      string `json:"frequencyRebalance"`
	FrequencyReconstruction string `json:"frequencyReconstruction"`
	NumberConstituents      int    `json:"numberConstituents"`
	DocumentationURL        string `json:"doc"`
	Description             string `json:"description"`
}

type config struct {
	URL        string  `json:"url"`
	Repository string  `json:"repository"`
	Indices    []index `json:"indices"`
}

func main() {
	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("Cannot get configuration: %s", err))
	}

	for _, index := range cfg.Indices {
		fmt.Println(index)
	}
}

func readConfig(fileName string) (*config, error) {
	var conf config

	f, err := os.Open(fileName)
	if err != nil {
		return &conf, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&conf)
	if err != nil {
		return &conf, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	if !strings.HasSuffix(conf.Repository, "/") {
		conf.Repository += "/"
	}

	return &conf, nil
}

func ensureRepositoryExists(repository string) {
	if _, err := os.Stat(repository); os.IsNotExist(err) {
		if err = os.MkdirAll(repository, os.ModePerm); err != nil {
			panic(fmt.Sprintf("cannot create repository directory '%s': %s", repository, err))
		}
	}
}

func (i *index) readSeries(repository string) (string, string) {
	var f *os.File
	var err error

	filePath := repository + strings.ToLower(i.Mnemonic) + ".csv"

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		if f, err = os.Create(filePath); err != nil {
			panic(fmt.Sprintf("cannot create file '%s': %s", filePath, err))
		} else {
			f.Close()
			return "", ""
		}
	}

	if f, err = os.Open(filePath); err != nil {
		panic(fmt.Sprintf("cannot open file '%s': %s", filePath, err))
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comment = '#'
	csvReader.Comma = ';'

	//if f, err := os.ReadFile() {}
	return "", ""
}

/*
// DevicesJSON returns a list of user devices as a raw JSON.
func (f *Index) Download(code string) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/" + convertToRequestID(userID) + "/devices.json")
}

// Devices converts the raw JSON to the Device slice type.
func (f *Fitbit) Devices(jsn []byte) ([]Device, error) {
	device := []Device{}
	if err := json.Unmarshal(jsn, &device); err != nil {
		return device, err
	}

	return device, nil
}
*/
