package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"euronext/euronext"
)

const configFileName = "eodimp20131228.json"

type config struct {
	Repository    string `json:"repository"`
	Downloads     string `json:"downloads"`
	XmlInstrumnts string `json:"xmlInstruments"`
}

type instrument struct {
	Mnemonic string `json:"mnemonic"`
	Mep      string `json:"mep"`
	Mic      string `json:"mic"`
	Isin     string `json:"isin"`
	Type     string `json:"type"`
}

type statis struct {
	markedInput int
	mergedNew   int
	mergedOld   int
	mergedSame  int
	mergedDiff  int
}

func main() {
	t := time.Now().Format("2006-01-02 15-04-05")
	fmt.Println("=======================================")
	fmt.Println(t)
	fmt.Println("=======================================")

	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("cannot read configuration file %s: %s", configFileName, err))
	}

	err = euronext.EnsureDirectoryExists(cfg.Repository)
	if err != nil {
		panic(fmt.Sprintf("cannot create directory %s: %s", cfg.Repository, err))
	}

	fmt.Println("xml file: " + cfg.XmlInstrumnts)
	instruments, err := readInstruments(cfg.XmlInstrumnts)
	if err != nil {
		panic(fmt.Sprintf("cannot read instruments: %s", err))
	}

	l := len(instruments)
	for i, ins := range instruments {
		err := ins.imp(cfg, i, l)
		if err != nil {
			fmt.Printf("\n%s\n", err)
		}
	}

	fmt.Println("\nfinished " + time.Now().Format("2006-01-02 15-04-05"))
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

	if !strings.HasSuffix(conf.Downloads, "/") {
		conf.Downloads += "/"
	}

	return &conf, nil
}

func readInstruments(fileName string) ([]instrument, error) {
	instruments := []instrument{}
	instrs, err := euronext.ReadXmlInstrumentsFile(fileName)
	if err != nil {
		return instruments, fmt.Errorf("cannot read instruments xml file '%s': %w", fileName, err)
	}

	fmt.Println(len(instrs.Instrument), "instruments read from", fileName)
	for _, inst := range instrs.Instrument {
		ins := instrument{
			Mnemonic: strings.ToLower(inst.Symbol),
			Mep:      strings.ToLower(inst.Mep),
			Mic:      strings.ToLower(inst.Mic),
			Isin:     strings.ToLower(inst.Isin),
			Type:     strings.ToLower(inst.Type),
		}
		instruments = append(instruments, ins)
	}

	return instruments, nil
}

func (s *instrument) safeMnemonic() string {
	mnemonic := s.Mnemonic
	if mnemonic == "prn" || mnemonic == "com" || mnemonic == "lpt" || mnemonic == "aux" || mnemonic == "com5" {
		mnemonic += "_"
	}

	return mnemonic
}

func (s *instrument) filePrefix() string {
	return fmt.Sprintf("%s_%s_%s_", s.Mic, s.Mnemonic, s.Isin)
}

func (s *instrument) filePathZip(dateSuffix string) string {
	return fmt.Sprintf("%s/%s/%s/endofday/%s_%s_%s_%s.zip",
		s.Mic, s.Type, s.safeMnemonic(), s.Mic, s.Mnemonic, s.Isin, dateSuffix)
}

func (s *instrument) filePathCsvGz() string {
	return fmt.Sprintf("%s/%s/%s/%s_%s_%s.1d.csv.gz",
		s.Mic, s.Type, s.safeMnemonic(), s.Mic, s.Mnemonic, s.Isin)
}

func (s *instrument) imp(cfg *config, el, elen int) error {
	dateSuffix := "2013-12-28"
	insZip := cfg.Downloads + s.filePathZip(dateSuffix)
	insCsvGz := cfg.Repository + s.filePathCsvGz()
	fmt.Printf("(%d of %d) %s ... ", el+1, elen, s.filePrefix())

	if _, err := os.Stat(insZip); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("no zip found")
			return nil
		} else {
			return fmt.Errorf("cannot check zip file '%s': %w", insZip, err)
		}
	}

	z, err := os.Open(insZip)
	if err != nil {
		return fmt.Errorf("cannot open zip file '%s': %w", insZip, err)
	}
	defer z.Close()

	stat, err := z.Stat()
	if err != nil {
		return fmt.Errorf("cannot stat zip file '%s': %w", insZip, err)
	}

	zipReader, err := zip.NewReader(z, stat.Size())
	if err != nil {
		return fmt.Errorf("cannot create zip reader for '%s': %w", insZip, err)
	}

	histOld, _, err := readCombinedDailyHistoryCsv(insCsvGz)
	if err != nil {
		return fmt.Errorf("cannot read old history file '%s': %w", insCsvGz, err)
	}

	sta := statis{}
	lenBefore := len(histOld)
	cntEntries := 0
	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("cannot open zip entry '%s': %w", file.Name, err)
		}
		defer rc.Close()

		// Read the contents of the file
		content, err := io.ReadAll(rc)
		if err != nil {
			return fmt.Errorf("cannot read zip entry '%s': %w", file.Name, err)
		}

		histNew, err := convertToCombinedDailyHistory(content, &sta)
		if err != nil {
			return fmt.Errorf("cannot convert zip entry '%s': %w", file.Name, err)
		}

		var messages []string
		histOld, messages = mergeCombinedDailyHistory(histOld, histNew, &sta)
		cntEntries += 1
		if len(messages) > 0 {
			fmt.Printf("\n[%d] %s messages:\n", cntEntries, file.Name)
			for i, message := range messages {
				fmt.Print(message)
				if i < len(messages)-1 {
					fmt.Print("\n")
				}
			}
		}
	}

	err = writeCombinedDailyHistoryCsv(insCsvGz, histOld, false)
	if err != nil {
		fmt.Printf("\ncannot write history file '%s': %s", insCsvGz, err)
	}

	lenAfter := len(histOld)
	fmt.Printf("\n(lines before %d -> after %d) input files %d [marked input %d, merged old %d merged replace (same %d, diff %d), merged new %d]\n", lenBefore, lenAfter, cntEntries, sta.markedInput, sta.mergedOld, sta.mergedSame, sta.mergedDiff, sta.mergedNew)
	return nil
}

func readCombinedDailyHistoryCsv(file string) ([]euronext.CombinedDailyHistory, string, error) {
	if _, err := os.Stat(file); err == nil {
		return euronext.ReadCombinedDailyHistoryCsv(file)
	} else if os.IsNotExist(err) {
		return make([]euronext.CombinedDailyHistory, 0), "", nil
	} else {
		es := fmt.Sprintf("error checking if file '%s' exists: ", file)
		return make([]euronext.CombinedDailyHistory, 0), es, fmt.Errorf("%s", es)
	}
}

func writeCombinedDailyHistoryCsv(file string, content []euronext.CombinedDailyHistory, backup bool) error {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("cannot create directory '%s': %w", dir, err)
	}

	if _, err := os.Stat(file); err == nil {
		if !backup {
			err = os.Remove(file)
			if err != nil {
				return fmt.Errorf("cannot remove file '%s': %w", file, err)
			}
		} else {
			file1 := file + ".bak"
			if _, err := os.Stat(file); err == nil {
				err = os.Remove(file)
				if err != nil {
					return fmt.Errorf("cannot remove file '%s': %w", file, err)
				}
			}
			err = os.Rename(file, file1)
			if err != nil {
				return fmt.Errorf("cannot backup file '%s': %w", file, err)
			}
		}
		_, err = euronext.WriteCombinedDailyHistoryCsv(file, content)
		return err
	} else if os.IsNotExist(err) {
		_, err = euronext.WriteCombinedDailyHistoryCsv(file, content)
		return err
	} else {
		return fmt.Errorf("error checking if file '%s' exists: ", file)
	}
}

func convertToCombinedDailyHistory(bs []byte, sta *statis) ([]euronext.CombinedDailyHistory, error) {
	combinedHist := []euronext.CombinedDailyHistory{}
	if len(bs) < 8 {
		return combinedHist, nil
	}

	jsonEod := euronext.JsonEod{}
	err := json.Unmarshal(bs, &jsonEod)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	for i, data := range jsonEod.Data {
		s := strings.TrimSpace(data.Date)
		s = strings.ReplaceAll(s, "\\/", "/")
		time, err := time.Parse("02/01/2006", s)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse date '%s' in item '%v': %w", i, data.Date, data, err)
		}

		s, open, marking, err := parseFloat(data.Open, false)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse open price '%s' in item '%v': %w", i, s, data, err)
		}

		s, high, marking, err := parseFloat(data.High, marking)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse high price '%s' in item '%v': %w", i, s, data, err)
		}

		s, low, marking, err := parseFloat(data.Low, marking)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse low price '%s' in item '%v': %w", i, s, data, err)
		}

		s, close, marking, err := parseFloat(data.Close, marking)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse close price '%s' in item '%v': %w", i, s, data, err)
		}

		s, numShares, marking, err := parseFloat(data.Nymberofshares, marking)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse number of shares '%s' in item '%v': %w", i, s, data, err)
		}

		s, numTrades, marking, err := parseFloat(data.Numoftrades, marking)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse number of trades '%s' in item '%v': %w", i, s, data, err)
		}

		s, turnover, marking, err := parseFloat(data.Turnover, marking)
		if err != nil {
			return combinedHist, fmt.Errorf("item %d: cannot parse turnover '%s' in item '%v': %w", i, s, data, err)
		}

		if open == 0 && close != 0 {
			open = close
			marking = true
		}
		if high == 0 && close != 0 {
			high = close
			marking = true
		}
		if low == 0 && close != 0 {
			low = close
			marking = true
		}

		if close == 0 && open != 0 {
			close = open
			marking = true
		}
		if high == 0 && open != 0 {
			high = open
			marking = true
		}
		if low == 0 && open != 0 {
			low = open
			marking = true
		}

		if close == 0 && high != 0 {
			close = high
			marking = true
		}
		if open == 0 && high != 0 {
			open = high
			marking = true
		}
		if low == 0 && high != 0 {
			low = high
			marking = true
		}

		if close == 0 && low != 0 {
			close = low
			marking = true
		}
		if open == 0 && low != 0 {
			open = low
			marking = true
		}
		if high == 0 && low != 0 {
			high = low
			marking = true
		}

		if marking {
			sta.markedInput += 1
		}

		if open != 0 || high != 0 || low != 0 || close != 0 {
			if open == 0 {
				fmt.Printf("\nitem %d: zero open value in '%v'", i, data)
			}
			if high == 0 {
				fmt.Printf("\nitem %d: zero open value in '%v'", i, data)
			}
			if low == 0 {
				fmt.Printf("\nitem %d: zero open value in '%v'", i, data)
			}
			if close == 0 {
				fmt.Printf("\nitem %d: zero open value in '%v'", i, data)
			}
		} else {
			fmt.Printf("\nitem %d: zero price values in '%v'", i, data)
		}

		entry := euronext.CombinedDailyHistory{
			Date:                   time,
			Open:                   open,
			High:                   high,
			Low:                    low,
			Last:                   close,
			Close:                  close,
			NumberOfShares:         numShares,
			NumberOfTrades:         numTrades,
			Turnover:               turnover,
			Vwap:                   0,
			OpenAdjusted:           open,
			HighAdjusted:           high,
			LowAdjusted:            low,
			LastAdjusted:           close,
			CloseAdjusted:          close,
			NumberOfSharesAdjusted: numShares,
			NumberOfTradesAdjusted: numTrades,
			TurnoverAdjusted:       turnover,
			VwapAdjusted:           0,
			AdjustmentFactor:       1,
			HasMarking:             marking,
			HasMarkingAdjusted:     marking,
		}
		combinedHist = append(combinedHist, entry)
	}

	return combinedHist, nil
}

func parseFloat(s string, marking bool) (string, float64, bool, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 || s[0] == '-' {
		return s, 0, true, nil
	}
	if s == "0" || s == "0.0" {
		return s, 0, marking, nil
	}
	s = strings.ReplaceAll(s, ",", "")
	v, err := strconv.ParseFloat(s, 64)
	return s, v, marking, err
}

func mergeCombinedDailyHistory(histOld, histNew []euronext.CombinedDailyHistory, sta *statis) ([]euronext.CombinedDailyHistory, []string) {
	messages := []string{}
	histMapOld := make(map[time.Time]euronext.CombinedDailyHistory)
	histMapNew := make(map[time.Time]euronext.CombinedDailyHistory)

	for _, entry := range histOld {
		histMapOld[entry.Date] = entry
	}
	for _, entry := range histNew {
		histMapNew[entry.Date] = entry
	}

	// Create a set of all dates
	dateSet := make(map[time.Time]struct{})
	for date := range histMapOld {
		dateSet[date] = struct{}{}
	}
	for date := range histMapNew {
		dateSet[date] = struct{}{}
	}

	// Collect the dates into a slice
	var dates []time.Time
	for date := range dateSet {
		dates = append(dates, date)
	}

	// Sort the slice in descending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].After(dates[j])
	})

	// Iterate through the sorted slice in descending order
	var mergedHistory []euronext.CombinedDailyHistory
	for _, date := range dates {
		entryOld, existsOld := histMapOld[date]
		entryNew, existsNew := histMapNew[date]
		notEqual := []string{}

		if existsNew && !existsOld {
			mergedHistory = append(mergedHistory, entryNew)
			sta.mergedNew += 1
			if entryNew.High < entryNew.Low || entryNew.High < entryNew.Open || entryNew.High < entryNew.Close || entryNew.High < entryNew.Last ||
				entryNew.Low > entryNew.Open || entryNew.Low > entryNew.Close || entryNew.Open > entryNew.High || entryNew.Low > entryNew.Last {
				notEqual = append(notEqual, fmt.Sprintf("(new) malformed price bar: %g, %g, %g, %g, %g", entryNew.Open, entryNew.High, entryNew.Low, entryNew.Close, entryNew.Last))
			}
			if len(notEqual) > 0 {
				sta.mergedDiff += 1
				messages = append(messages,
					fmt.Sprintf("Date %s:  %s",
						date.Format("2006-01-02"), strings.Join(notEqual, ", ")))
			}

		} else if existsNew && existsOld {
			if entryOld.Open != entryNew.Open {
				notEqual = append(notEqual, fmt.Sprintf("diff open: %g -> %g", entryOld.Open, entryNew.Open))
			}
			if entryOld.High != entryNew.High {
				notEqual = append(notEqual, fmt.Sprintf("diff high: %g -> %g", entryOld.High, entryNew.High))
			}
			if entryOld.Low != entryNew.Low {
				notEqual = append(notEqual, fmt.Sprintf("diff low: %g -> %g", entryOld.Low, entryNew.Low))
			}
			if entryOld.Last != entryNew.Last {
				notEqual = append(notEqual, fmt.Sprintf("diff last: %g -> %g", entryOld.Last, entryNew.Last))
			}
			if entryOld.Close != entryNew.Close {
				notEqual = append(notEqual, fmt.Sprintf("diff close: %g -> %g", entryOld.Close, entryNew.Close))
			}
			if entryOld.NumberOfShares != entryNew.NumberOfShares {
				notEqual = append(notEqual, fmt.Sprintf("diff num shares: %g -> %g", entryOld.NumberOfShares, entryNew.NumberOfShares))
			}
			if entryOld.NumberOfTrades != entryNew.NumberOfTrades {
				notEqual = append(notEqual, fmt.Sprintf("diff num trades: %g -> %g", entryOld.NumberOfTrades, entryNew.NumberOfTrades))
			}
			if entryOld.Turnover != entryNew.Turnover {
				notEqual = append(notEqual, fmt.Sprintf("diff turnover: %g -> %g", entryOld.Turnover, entryNew.Turnover))
			}
			if entryOld.Vwap != entryNew.Vwap {
				notEqual = append(notEqual, fmt.Sprintf("diff vwap: %g -> %g", entryOld.Vwap, entryNew.Vwap))
			}

			useOld := false
			if entryNew.Open == 0 && entryOld.Open != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace open: %g -> %g", entryOld.Open, entryNew.Open))
				useOld = true
			}
			if entryNew.High == 0 && entryOld.High != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace high: %g -> %g", entryOld.High, entryNew.High))
				useOld = true
			}
			if entryNew.Low == 0 && entryOld.Low != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace low: %g -> %g", entryOld.Low, entryNew.Low))
				useOld = true
			}
			if entryNew.Close == 0 && entryOld.Close != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace close: %g -> %g", entryOld.Low, entryNew.Low))
				useOld = true
			}
			if entryNew.Last == 0 && entryOld.Last != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace last: %g -> %g", entryOld.Last, entryNew.Last))
				useOld = true
			}
			if entryNew.NumberOfShares == 0 && entryOld.NumberOfShares != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace shares: %g -> %g", entryOld.NumberOfShares, entryNew.NumberOfShares))
				useOld = true
			}
			if entryNew.NumberOfTrades == 0 && entryOld.NumberOfTrades != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace trades: %g -> %g", entryOld.NumberOfTrades, entryNew.NumberOfTrades))
				useOld = true
			}
			if entryNew.Turnover == 0 && entryOld.Turnover != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace turnover: %g -> %g", entryOld.Turnover, entryNew.Turnover))
				useOld = true
			}
			if entryNew.Vwap == 0 && entryOld.Vwap != 0 {
				notEqual = append(notEqual, fmt.Sprintf("wont replace vwap: %g -> %g", entryOld.Vwap, entryNew.Vwap))
				useOld = true
			}

			if entryOld.HasMarking || entryNew.HasMarking {
				entryNew.HasMarking = true
			}
			if entryOld.HasMarkingAdjusted || entryNew.HasMarkingAdjusted {
				entryNew.HasMarkingAdjusted = true
			}

			if !useOld {
				if entryNew.High < entryNew.Low || entryNew.High < entryNew.Open || entryNew.High < entryNew.Close || entryNew.High < entryNew.Last ||
					entryNew.Low > entryNew.Open || entryNew.Low > entryNew.Close || entryNew.Open > entryNew.High || entryNew.Low > entryNew.Last {
					notEqual = append(notEqual, fmt.Sprintf("(repl) malformed price bar: %g, %g, %g, %g, %g", entryNew.Open, entryNew.High, entryNew.Low, entryNew.Close, entryNew.Last))
				}
			}

			if len(notEqual) > 0 {
				sta.mergedDiff += 1
				messages = append(messages,
					fmt.Sprintf("Date %s: %s",
						date.Format("2006-01-02"), strings.Join(notEqual, ", ")))
			} else {
				sta.mergedSame += 1
			}

			if useOld {
				mergedHistory = append(mergedHistory, entryOld)
			} else {
				mergedHistory = append(mergedHistory, entryNew)
			}
		} else { // if !existsNew && existsOld
			sta.mergedOld += 1
			mergedHistory = append(mergedHistory, entryOld)
		}
	}

	return euronext.SortCombinedDailyHistory(mergedHistory), messages
}
