package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"nq/nasdaq"
)

const configFileName = "nasdaq.json"

type symbol struct {
	Mnemonic    string `json:"mnemonic"`
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	StockType   string `json:"stockType"`
	Exchange    string `json:"exchange"`
	IsListed    bool   `json:"isListed"`
	IsNasdaq100 bool   `json:"isNasdaq100"`
}

type config struct {
	Repository string   `json:"repository"`
	Symbols    []symbol `json:"symbols"`
}

func main() {
	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("Cannot get configuration: %s", err))
	}

	for _, s := range cfg.Symbols {
		if err = s.update(cfg.Repository); err != nil {
			fmt.Printf("%s: %s\n", s.Mnemonic, err)
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

func (s *symbol) update(repository string) error {
	a, b, err := nasdaq.RetrieveSession(s.Mnemonic)
	if err != nil {
		return err
	}

	fmt.Print(a, b)
	return nil
}
