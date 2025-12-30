package main

import (
	"encoding/json"
	"fmt"
	"ms/ms"
	"os"
	"strings"
	"time"

	mscsv "ms/csv"
)

const configFileName = "ms.json"

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
		if err = index.updateSeries(cfg.Repository); err != nil {
			fmt.Printf("%s: %s\n", index.Mnemonic, err)
		}
	}

	fmt.Println("done")
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

func (i *index) updateSeries(repository string) error {
	s1, err := mscsv.ReadSeries(repository, i.Mnemonic)
	if err != nil {
		return fmt.Errorf("cannot read csv file: %w", err)
	}

	date := time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})
	if len(s1) > 0 {
		date = s1[len(s1)-1].Date.Add(24 * time.Hour)
	}

	s2, err := i.downloadSeries(date)
	if err != nil {
		return err
	}

	if len(s2) > 0 {
		s1 = append(s1, s2...) // s2[1:len(s2)]...
		if err = mscsv.WriteSeries(repository, i.Mnemonic, s1); err != nil {
			return fmt.Errorf("cannot write csv file: %w", err)
		}
	}

	return nil
}

func (i *index) downloadSeries(startDate time.Time) ([]mscsv.Point, error) {
	if msts, err := ms.Download(i.MorningstarID, i.BaseCurrency, startDate); err != nil {
		return nil, fmt.Errorf("cannot donload: %w", err)
	} else {
		series := make([]mscsv.Point, 0)
		secs := msts.TimeSeries.Security
		if len(secs) == 0 {
			fmt.Println("no data downloaded in time series")
			return series, nil
		}
		for _, p := range msts.TimeSeries.Security[0].HistoryDetail {
			series = append(series, mscsv.Point{
				Date:  p.EndDate.Time,
				Value: float64(p.Value),
			})
		}

		return series, nil
	}
}
