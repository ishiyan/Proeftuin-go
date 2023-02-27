package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"nq/nasdaq"
)

const configFileName = "nqrt.json"

type symbol struct {
	Mnemonic   string `json:"mnemonic"`
	Mic        string `json:"mic"`
	AssetClass string `json:"assetClass"`
}

type symbols struct {
	Symbols []symbol `json:"symbols"`
}

type config struct {
	Repository        string `json:"repository"`
	RetryDelayMinutes []int  `json:"retryDelayMinutes"`
}

func main() {
	t := time.Now().Format("2006-01-02 15-04-05")
	fmt.Println("=======================================")
	fmt.Println(t)
	fmt.Println("=======================================")

	symbolsPtr := flag.String("symbols", "nasdaq.json", "symbols json file name")
	flag.Parse()

	sym, err := readSymbols(*symbolsPtr)
	if err != nil {
		panic(fmt.Sprintf("cannot read symbols: %s", err))
	}

	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("cannot read configuration: %s", err))
	}

	err = ensureDirectoryExists(cfg.Repository)
	if err != nil {
		panic(fmt.Sprintf("cannot create repository directory: %s", err))
	}

	l := len(sym.Symbols)
	for i, s := range sym.Symbols {
		s.Mnemonic = strings.ToLower(s.Mnemonic)
		s.AssetClass = strings.ToLower(s.AssetClass)
		s.Mic = strings.ToLower(s.Mic)

		if len(s.Mic) < 1 {
			fmt.Printf("%s: empty mic, setting to 'other'\n", s.Mnemonic)
			s.Mic = "other"
		}

		p := fmt.Sprintf("(%d of %d)", i+1, l)
		if err = s.archive(cfg.Repository, p, cfg.RetryDelayMinutes); err != nil {
			fmt.Printf("%s: %s\n", s.Mnemonic, err)
		}
	}

	fmt.Println("finished " + time.Now().Format("2006-01-02 15-04-05"))
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

func readSymbols(fileName string) (*symbols, error) {
	var s symbols

	f, err := os.Open(fileName)
	if err != nil {
		return &s, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&s)
	if err != nil {
		return &s, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	return &s, nil
}

func ensureDirectoryExists(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = os.MkdirAll(directory, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory '%s': %w", directory, err)
		}
	}

	return nil
}

func (s *symbol) archive(repository, prefix string, retryDelayMins []int) error {
	path := repository + s.Mic + "/" + s.AssetClass + "/" + s.Mnemonic
	if s.Mnemonic == "prn" {
		path += "_"
	}
	if s.Mnemonic == "com" {
		path += "_"
	}

	fmt.Printf("%s '%s' to '%s' ... ", prefix, s.Mnemonic, path)

	err := ensureDirectoryExists(path)
	if err != nil {
		return fmt.Errorf("cannot create symbol repository directory '%s': %s", path, err)
	}

	var t time.Time
	var csv []string
	var json []nasdaq.NasdaqRealtimeJSON

	path += "/"
	retriesMax := len(retryDelayMins)
	retries := 0
	for retries < retriesMax {
		t, csv, json, err = nasdaq.RetrieveSession(s.Mnemonic)
		if err != nil {
			retries += 1
			err := fmt.Errorf("failed to retrieve trades, retries (%d of %d): %w", retries, retriesMax, err)
			fmt.Println(err)
			nasdaq.ResetCookie()
			if retries >= retriesMax {
				fmt.Printf("giving up after %d retries\n", retriesMax)
				return err
			} else {
				mins := retryDelayMins[retries]
				fmt.Printf("waiting %d minutes before %d retry ...\n", mins, retries+1)
				time.Sleep(time.Duration(mins) * time.Minute)
			}
		} else {
			break
		}
	}

	file := path + t.Format("2006-01-02") // _15-04-05
	fz := file + "_trade.zip"

	a := 0
again:
	_, err = os.Stat(fz)
	if err == nil {
		a++
		fz = file + fmt.Sprintf("_trade(%d).zip", a)
		goto again
	}

	z, err := os.Create(fz)
	if err != nil {
		return fmt.Errorf("cannot create '%s': %w", fz, err)
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()

	td := t.Format("2006-01-02") + "_"
	nam := td + "trade.csv"
	f, err := w.Create(nam)
	if err != nil {
		return fmt.Errorf("cannot create zip entry '%s': %w", nam, err)
	}

	joined := strings.Join(csv, "")
	_, err = f.Write([]byte(joined))
	if err != nil {
		return fmt.Errorf("cannot write zip entry '%s': %w", nam, err)
	}

	for _, j := range json {
		nam = td + j.Period + "_trade.json"
		f, err = w.Create(nam)
		if err != nil {
			return fmt.Errorf("cannot create zip entry '%s': %w", nam, err)
		}

		_, err = f.Write(j.JSON)
		if err != nil {
			return fmt.Errorf("cannot write zip entry '%s': %w", nam, err)
		}
	}

	fmt.Println("done")
	return nil
}
