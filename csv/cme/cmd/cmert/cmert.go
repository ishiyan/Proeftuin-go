package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"nq/cme"
	"os"
	"strconv"
	"strings"
	"time"
)

const configFileName = "cmert.json"

type symbols struct {
	Symbols []cme.FutureSymbol `json:"symbols"`
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

	symbolsPtr := flag.String("symbols", "cme.json", "symbols json file name")
	daysbackPtr := flag.Int("daysback", 0, "days back before entry date")
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
		s.Future = strings.ToLower(s.Future)
		s.Mic = strings.ToLower(s.Mic)

		p := fmt.Sprintf("(%d of %d)", i+1, l)
		if err = archive(s, cfg.Repository, p, cfg.RetryDelayMinutes, *daysbackPtr); err != nil {
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

func sessionDate() (time.Time, error) {
	today := time.Now().Add(time.Hour * -7)
	/*loc, err := time.LoadLocation("EST")
	if err != nil {
		return today, fmt.Errorf("cannot load EST timezone: %w", err)
	}

	today = today.In(loc)*/
	dow := today.Weekday()
	if dow == time.Sunday {
		return today.AddDate(0, 0, -1), nil
	} else if dow == time.Monday {
		return today.AddDate(0, 0, -2), nil
	} else {
		return today/*.AddDate(0, 0, -1)*/, nil
	}
}

func ensureDirectoryExists(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = os.MkdirAll(directory, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory '%s': %w", directory, err)
		}
	}

	return nil
}

func archive(sym cme.FutureSymbol, repository, prefix string, retryDelayMins []int, daysback int) error {
	path := repository + sym.Mic + "/" + sym.Future + "/" + sym.Mnemonic // xcme/es/esm23
	fmt.Printf("%s '%s' to '%s' ...\n", prefix, sym.Mnemonic, path)

	err := ensureDirectoryExists(path)
	if err != nil {
		return fmt.Errorf("cannot create symbol repository directory '%s': %s", path, err)
	}

	t, err := sessionDate()
	if err != nil {
		return err
	}

	fmt.Printf("entry date: '%s'\n", t.Format("2006-01-02"))
	if daysback != 0 {
		t = t.AddDate(0, 0, daysback)
		fmt.Printf("entry date: '%s' with %d days back\n", t.Format("2006-01-02"), daysback)
	}

	// var parsed []cme.RealtimeEntryParsed
	var json cme.RealtimeJSON
	var csv []string

	path += "/"
	retriesMax := len(retryDelayMins)
	retries := 0
	for retries < retriesMax {
		csv, _, json, err = cme.RetrieveCode(sym, t)
		if err != nil {
			retries += 1
			err := fmt.Errorf("failed to retrieve trades, retries (%d of %d): %w", retries, retriesMax, err)
			fmt.Println(err)
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

	td := t.Format("2006-01-02") + "_trade"
	nam := td + ".csv"
	f, err := w.Create(nam)
	if err != nil {
		return fmt.Errorf("cannot create zip entry '%s': %w", nam, err)
	}

	joined := strings.Join(csv, "")
	_, err = f.Write([]byte(joined))
	if err != nil {
		return fmt.Errorf("cannot write zip entry '%s': %w", nam, err)
	}

	for _, j := range json.Timeslots {
		nam = td + "_" + j.Timeslot + "_"
		for _, i := range j.JSON {
			n := nam + strconv.Itoa(i.Page) + ".json"
			f, err = w.Create(n)
			if err != nil {
				return fmt.Errorf("cannot create zip entry '%s': %w", n, err)
			}

			_, err = f.Write(i.JSON)
			if err != nil {
				return fmt.Errorf("cannot write zip entry '%s': %w", n, err)
			}
		}
	}

	fmt.Println("done")
	return nil
}
