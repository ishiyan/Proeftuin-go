package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	estr "estr/estr"
)

const configFileName = "estr.json"

type config struct {
	Repository string `json:"repository"`
	Actual     bool   `json:"actual"`
	Pre        bool   `json:"pre"`
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

func downloadSeries(what estr.What, startDate time.Time) ([]estr.Point, error) {
	if pts, err := estr.Fetch(what); err != nil {
		return nil, fmt.Errorf("cannot download: %w", err)
	} else {
		flt := make([]estr.Point, 0)
		for _, p := range pts {
			if p.Date.Before(startDate) {
				continue
			}
			flt = append(flt, p)
		}

		return flt, nil
	}
}

func updateSeries(repository string, what estr.What) error {
	s1, err := estr.ReadCSV(repository, what)
	if err != nil {
		return fmt.Errorf("cannot read csv file: %w", err)
	}

	date := time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})
	if len(s1) > 0 {
		date = s1[len(s1)-1].Date.Add(24 * time.Hour)
	}

	s2, err := downloadSeries(what, date)
	if err != nil {
		return err
	}

	if len(s2) > 0 {
		s1 = append(s1, s2...) // s2[1:len(s2)]...
		if err = estr.WriteCSV(repository, what, s1); err != nil {
			return fmt.Errorf("cannot write csv file: %w", err)
		}
	}

	return nil
}

func main() {
	const delimiter = "======================================="

	t := time.Now().Format("2006-01-02 15-04-05")
	fmt.Println(delimiter)
	fmt.Println(t)
	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("Cannot get configuration: %s", err))
	}

	if cfg.Pre {
		fmt.Println("Updating pre-series...")
		for _, w := range []estr.What{estr.EstrRatePre, estr.EstrVolumePre, estr.EstrTransactionsPre} {
			if err = updateSeries(cfg.Repository, w); err != nil {
				fmt.Printf("%s: %s\n", estr.WhatMnemonic(w), err)
			}
		}
	}

	if cfg.Actual {
		fmt.Println("Updating actual series...")
		for _, w := range []estr.What{estr.EstrRateAct, estr.EstrVolumeAct, estr.EstrTransactionsAct} {
			if err = updateSeries(cfg.Repository, w); err != nil {
				fmt.Printf("%s: %s\n", estr.WhatMnemonic(w), err)
			}
		}
	}

	fmt.Println("done")
	fmt.Println(delimiter)
}
