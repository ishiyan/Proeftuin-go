package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"euronext/euronext"
	"euronext/euronext/intraday"
)

const configFileName = "enxintr.json"

type config struct {
	DownloadsFolder             string `json:"downloadsFolder"`
	RepositoryFolder            string `json:"repositoryFolder"`
	XmlInstrumntsFile           string `json:"xmlInstrumentsFile"`
	RepositoryGzipped           bool   `json:"repositoryGzipped"`
	DownloadRetryDelaySeconds   []int  `json:"downloadRetryDelaySeconds"`
	DownloadTimeoautSeconds     int    `json:"downloadTimeoutSeconds"`
	Concurrency                 int    `json:"concurrency"`
	UserAgent                   string `json:"userAgent"`
	DownloadRetryDelayDurations []time.Duration
	DownloadTimeoautDuration    time.Duration
}

type instrument struct {
	Mnemonic string `json:"mnemonic"`
	Mep      string `json:"mep"`
	Mic      string `json:"mic"`
	Isin     string `json:"isin"`
	Type     string `json:"type"`
}

type combi struct {
	DownloadError string
	Index         int
	Length        int
	Instrument    instrument
	Raw           []byte
}

type statistics struct {
	DownloadErrors []string
	MergeErrors    []string
	MergeMessages  []string
	ZeroLines      []string
	NoHistoryLines []string
}

func main() {
	now := time.Now()
	t := now.Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("enxintr_%s.log", t)
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Panicf("cannot create log file '%s': %s\n", logFileName, err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	sessionDate, err := euronext.SessionDate()
	if err != nil {
		log.Panicf("cannot get session date: %s\n", err)
	}
	log.Println("trading session date: " + sessionDate.Format("2006-01-02"))

	cfg, err := readConfig(configFileName)
	if err != nil {
		log.Panicf("cannot read configuration file %s: %s\n", configFileName, err)
	}

	err = euronext.EnsureDirectoryExists(cfg.RepositoryFolder)
	if err != nil {
		log.Panicf("cannot create repository directory %s: %s\n", cfg.RepositoryFolder, err)
	}

	err = euronext.EnsureDirectoryExists(cfg.DownloadsFolder)
	if err != nil {
		log.Panicf("cannot create downloads directory %s: %s\n", cfg.DownloadsFolder, err)
	}

	downloadName := fmt.Sprintf("%s", sessionDate.Format("20060102"))
	downloadPath := cfg.DownloadsFolder + "intradayday/" +
		fmt.Sprintf("%s/", sessionDate.Format("2006")) + downloadName + "/"

	err = euronext.EnsureDirectoryExists(downloadPath)
	if err != nil {
		log.Panicf("cannot create downloads directory %s: %s\n", downloadPath, err)
	}

	log.Println("xml instruments file: " + cfg.XmlInstrumntsFile)
	instruments, err := readInstruments(cfg.XmlInstrumntsFile)
	if err != nil {
		log.Panicf("cannot read instruments: %s\n", err)
	}

	stati := statistics{
		DownloadErrors: []string{"date;mep;mic;type;mnemonic;isin;error"},
		MergeErrors:    []string{"date;mep;mic;type;mnemonic;isin;error"},
		MergeMessages:  []string{"date;mep;mic;type;mnemonic;isin;message"},
		ZeroLines:      []string{"date;mep;mic;type;mnemonic;isin;lines"},
		NoHistoryLines: []string{"date;mep;mic;type;mnemonic;isin;lines"},
	}

	log.Println("=======================================")

	l := len(instruments)
	if cfg.Concurrency < 2 {
		for i, ins := range instruments {
			combi, err := ins.download(sessionDate, cfg, i, l, downloadPath, downloadName)
			if err == nil {
				combi.Instrument.merge(sessionDate, cfg, combi, &stati)
			}
		}
	} else {
		var wg sync.WaitGroup
		sem := make(chan struct{}, cfg.Concurrency)

		// Channel to pass data to the second (merging) pipeline
		combiChan := make(chan combi, l)

		// WaitGroup for the second pipeline
		var mergeWg sync.WaitGroup

		// Start the second pipeline to merge data from the channel
		mergeWg.Add(1)
		go func() {
			defer mergeWg.Done()
			for combi := range combiChan {
				combi.Instrument.merge(sessionDate, cfg, combi, &stati)
			}
		}()

		for i, ins := range instruments {
			wg.Add(1)
			sem <- struct{}{} // Acquire a slot

			go func(i, l int, ins instrument) {
				defer wg.Done()
				defer func() { <-sem }() // Release the slot

				combi, err := ins.download(sessionDate, cfg, i, l, downloadPath, downloadName)
				if err == nil {
					combiChan <- combi
				}
			}(i, l, ins)
		}

		// Wait for the first pipeline to finish processing
		wg.Wait()

		// Close the channel after all goroutines have finished
		close(combiChan)

		// Wait for the second pipeline to finish processing
		mergeWg.Wait()
	}

	_ = zipDownloads(downloadPath)

	log.Println("\nprocessed " + time.Now().Format("2006-01-02 15-04-05"))

	log.Printf("\n\ninstruments with download errors: %d from %d\n", len(stati.DownloadErrors)-1, l)
	for _, z := range stati.DownloadErrors {
		log.Println(z)
	}

	log.Printf("\n\ninstruments with merge errors: %d from %d\n", len(stati.MergeErrors)-1, l)
	for _, z := range stati.MergeErrors {
		log.Println(z)
	}

	log.Printf("\n\ninstruments with merge messages: %d from %d\n", len(stati.MergeMessages)-1, l)
	for _, z := range stati.MergeMessages {
		log.Println(z)
	}

	log.Printf("\n\ninstruments with zero lines: %d from %d\n", len(stati.ZeroLines)-1, l)
	for _, z := range stati.ZeroLines {
		log.Println(z)
	}

	log.Printf("\n\ninstruments with valid header but no history: %d from %d\n", len(stati.NoHistoryLines)-1, l)
	for _, z := range stati.NoHistoryLines {
		log.Println(z)
	}

	log.Println("\nfinished " + time.Now().Format("2006-01-02 15-04-05"))
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

	if !strings.HasSuffix(conf.DownloadsFolder, "/") {
		conf.DownloadsFolder += "/"
	}

	if !strings.HasSuffix(conf.RepositoryFolder, "/") {
		conf.RepositoryFolder += "/"
	}

	if conf.Concurrency < 2 {
		conf.Concurrency = 0
	} else if conf.Concurrency > 8 {
		conf.Concurrency = 8
	}

	if conf.DownloadTimeoautSeconds < 1 {
		conf.DownloadTimeoautSeconds = 1
	}
	conf.DownloadTimeoautDuration = time.Duration(conf.DownloadTimeoautSeconds) * time.Second

	conf.DownloadRetryDelayDurations = make([]time.Duration, len(conf.DownloadRetryDelaySeconds))
	for i, delay := range conf.DownloadRetryDelaySeconds {
		if delay < 1 {
			delay = 1
		}

		conf.DownloadRetryDelayDurations[i] = time.Duration(delay) * time.Second
	}

	return &conf, nil
}

func readInstruments(fileName string) ([]instrument, error) {
	instruments := []instrument{}
	instrs, err := euronext.ReadXmlInstrumentsFile(fileName)
	if err != nil {
		return instruments, fmt.Errorf("cannot read instruments xml file '%s': %w", fileName, err)
	}

	log.Printf(" %d instruments read from %s\n", len(instrs.Instrument), fileName)
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
	if mnemonic == "prn" || mnemonic == "com" || mnemonic == "lpt" || mnemonic == "aux" {
		mnemonic += "_"
	}

	return mnemonic
}

func (s *instrument) fileFolder() string {
	return fmt.Sprintf("%s/%s/%s/", s.Mic, s.Type, s.safeMnemonic())
}

func (s *instrument) fileName() string {
	return fmt.Sprintf("%s_%s_%s", s.Mic, s.Mnemonic, s.Isin)
}

func (s *instrument) download(
	sessionDate time.Time,
	cfg *config,
	el int,
	elen int,
	downloadFolder string,
	downloadName string,
) (combi, error) {
	insName := s.fileName()
	prefix := fmt.Sprintf("(%d of %d) %s %s to %s ... ", el+1, elen, insName, s.Type, downloadName)
	combi := combi{
		DownloadError: "",
		Index:         el,
		Length:        elen,
		Instrument:    *s,
		Raw:           nil,
	}

	bs, err := intraday.GetIntradayData(s.Isin, s.Mic, s.Mnemonic, s.Type,
		cfg.DownloadTimeoautDuration, cfg.DownloadRetryDelayDurations, cfg.UserAgent)
	if err != nil {
		combi.DownloadError = err.Error()
		log.Println(prefix + "skipping: " + combi.DownloadError)
		return combi, err
	}

	combi.Raw = bs
	prefix += "downloaded ... "

	prefix = fmt.Sprintf("(%d of %d) %s writing %s to %s ... ", el+1, elen, insName, s.Type, downloadName)
	jsonFile := filepath.Join(downloadFolder, insName+".json")
	err = os.WriteFile(jsonFile, bs, 0644)
	if err != nil {
		combi.DownloadError = "failed to save: " + err.Error()
		log.Println(prefix + combi.DownloadError)
	} else {
		log.Println(prefix + "saved")
	}

	return combi, nil
}

// zipFolder zips the folder at srcDir (including the folder itself) into destZip.
func zipFolder(srcDir, destZip string) error {
	z, err := os.Create(destZip)
	if err != nil {
		return fmt.Errorf("cannot create zip file '%s': %w", destZip, err)
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()

	parent := filepath.Dir(srcDir)
	err = filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil // skip directories, only add files
		}
		relPath, err := filepath.Rel(parent, path)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath) // for zip standard

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		wr, err := w.Create(relPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(wr, f)
		return err
	})
	return err
}

func zipDownloads(downloadFolder string) error {
	file := fmt.Sprintf("%senx_eoi", downloadFolder)
	fz := file + ".zip"
	counter := 0
again:
	_, err := os.Stat(fz)
	if err == nil {
		counter++
		fz = fmt.Sprintf("%s.%d.zip", file, counter)
		goto again
	}

	log.Printf("archiving downloads from %s to %s ... ", downloadFolder, fz)

	if err := zipFolder(downloadFolder, fz); err != nil {
		log.Println("failed: %v", err)
		return err
	} else {
		log.Println("done")
		return nil
	}
}

func (s *instrument) merge(sessionDate time.Time, cfg *config, combi combi, stati *statistics) {
	sd := sessionDate.Format("2006-01-02")
	insFolder := s.fileFolder()
	insName := s.fileName()
	prefix := fmt.Sprintf("[%d of %d] %s to %s ... ", combi.Index+1, combi.Length, insName, insFolder)

	if len(combi.DownloadError) > 0 {
		stati.DownloadErrors = append(stati.DownloadErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, combi.DownloadError))
		if combi.Raw == nil {
			log.Println(prefix + "not merged due to download error")
			return
		}
	}

	if len(combi.Raw) < 10 {
		es := fmt.Sprintf("raw data is too short, not merged: [%s]" + string(combi.Raw))
		log.Println(prefix + es)
		return
	}

	jsonInd := intraday.JsonIntraday{}
	err := json.Unmarshal(combi.Raw, &jsonInd)
	if err != nil {
		es := fmt.Sprintf("cannot unmarshal json data: %s", err)
		stati.MergeErrors = append(stati.MergeErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
		log.Println(prefix + es)
		return
	}

	if jsonInd.Rows == nil || len(jsonInd.Rows) == 0 {
		es := "no trades found"
		stati.NoHistoryLines = append(stati.NoHistoryLines,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
		log.Println(prefix + es)
		return
	}

	/*histNew, err := convertToCombinedDailyHistory(combined)
	combined = nil
	if err != nil {
		es := "cannot convert to combined daily history: "
		stati.MergeErrors = append(stati.MergeErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es+err.Error()))
		fmt.Println(log + es + err.Error())
		return
	}*/

	// xpar/stock/gle/intraday/202506.csv.gz
	// xpar/stock/gle/intraday/2025.csv.gz
	// xpar/stock/gle/intraday/2025_xpar_gle_ie1234567890_1t.csv.gz

	path := cfg.Repository + insFolder
	err = euronext.EnsureDirectoryExists(path)
	if err != nil {
		es := fmt.Sprintf("cannot create instrument repository directory '%s': ", path)
		stati.MergeErrors = append(stati.MergeErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
		fmt.Println(log + es + err.Error())
		return
	}

	file := fmt.Sprintf("%s%s.1d.csv", path, insName)
	if cfg.RepositoryGzipped {
		file += ".gz"
	}

	if _, err := os.Stat(file); err == nil {
		histOld, es, err := euronext.ReadCombinedDailyHistoryCsv(file)
		if err != nil {
			stati.MergeErrors = append(stati.MergeErrors,
				fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
			fmt.Println(log + es + err.Error())
			return
		}
		histMerged, messages := euronext.MergeCombinedDailyHistory(histOld, histNew)
		if len(messages) > 0 {
			for _, m := range messages {
				stati.MergeMessages = append(stati.MergeMessages,
					fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, m))
			}
		}
		es, err = euronext.BackupFile(file)
		if err != nil {
			stati.MergeErrors = append(stati.MergeErrors,
				fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
			fmt.Println(log + es + err.Error())
			return
		}
		es, err = euronext.WriteCombinedDailyHistoryCsv(file, histMerged)
		if err != nil {
			stati.MergeErrors = append(stati.MergeErrors,
				fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
			fmt.Println(log + es + err.Error())
			return
		}
	} else if os.IsNotExist(err) {
		histNew = euronext.SortCombinedDailyHistory(histNew)
		es, err := euronext.WriteCombinedDailyHistoryCsv(file, histNew)
		if err != nil {
			stati.MergeErrors = append(stati.MergeErrors,
				fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
			fmt.Println(log + es + err.Error())
			return
		}
	} else {
		es := fmt.Sprintf("error checking if file '%s' exists: ", file)
		stati.MergeErrors = append(stati.MergeErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", sd, s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, es))
		fmt.Println(log + es + err.Error())
		return
	}

	fmt.Println(log + "merged")
}

func combine(bsRaw, bsAdj []byte) ([]string, int, int) {
	linesRaw := strings.Split(string(bsRaw), "\n")
	linesAdj := strings.Split(string(bsAdj), "\n")
	lenRaw := len(linesRaw)
	lenAdj := len(linesAdj)
	combined := []string{}
	if lenRaw == lenAdj {
		for i := 0; i < lenRaw; i++ {
			if len(linesRaw[i]) == 0 && len(linesAdj[i]) == 0 {
				continue
			}
			combined = append(combined, linesRaw[i]+";"+linesAdj[i])
		}
	} else {
		l := max(lenRaw, lenAdj)
		for i := 0; i < l; i++ {
			if i < lenRaw && i < lenAdj {
				if len(linesRaw[i]) == 0 && len(linesAdj[i]) == 0 {
					continue
				}
				combined = append(combined, linesRaw[i]+";"+linesAdj[i])
			} else if i < lenRaw {
				if len(linesRaw[i]) == 0 {
					continue
				}
				combined = append(combined, linesRaw[i]+";")
			} else {
				if len(linesAdj[i]) == 0 {
					continue
				}
				combined = append(combined, ";"+linesAdj[i])
			}
		}
	}

	return combined, lenRaw, lenAdj
}

func convertToCombinedDailyHistory(lines []string) ([]euronext.CombinedDailyHistory, error) {
	combinedHist := []euronext.CombinedDailyHistory{}
	expectedParts := 20
	for i, line := range lines {
		if i < 3 {
			continue
		}

		if i == 3 {
			if !strings.HasPrefix(line, "Date;Open;") {
				return combinedHist, fmt.Errorf("line 4: unexpected header line: %s", line)
			}

			continue
		}

		if i == 4 && len(line) < 10 {
			return combinedHist, nil // Empty history
		}

		parts := strings.Split(line, ";")
		if len(parts) != expectedParts {
			return combinedHist, fmt.Errorf("line %d: expected %d line parts, got %d: %s", i+1, expectedParts, len(parts), line)
		}

		s0, marking := cleanString(parts[0], false)
		time, err := time.Parse("02/01/2006", s0)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse date '%s' in line '%s': %w", i+1, s0, line, err)
		}

		s, marking := cleanString(parts[10], marking)
		if s0 != s {
			return combinedHist, fmt.Errorf("line %d: date '%s' does not match adjusted date '%s' in line '%s'", i+1, s0, s, line)
		}

		s, openRaw, marking, err := parseFloat(parts[1], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse open price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, openAdj, marking, err := parseFloat(parts[11], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse open adjusted price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, highRaw, marking, err := parseFloat(parts[2], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse high price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, highAdj, marking, err := parseFloat(parts[12], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse high adjusted price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, lowRaw, marking, err := parseFloat(parts[3], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse low price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, lowAdj, marking, err := parseFloat(parts[13], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse low adjusted price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, lastRaw, marking, err := parseFloat(parts[4], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse last price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, lastAdj, marking, err := parseFloat(parts[14], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse last adjusted price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, closeRaw, marking, err := parseFloat(parts[5], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse close price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, closeAdj, marking, err := parseFloat(parts[15], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse close adjusted price '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, sharesRaw, marking, err := parseFloat(parts[6], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse number of shares '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, sharesAdj, marking, err := parseFloat(parts[16], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse number of adjusted shares '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, tradesRaw, marking, err := parseFloat(parts[7], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse number of trades '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, tradesAdj, marking, err := parseFloat(parts[17], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse number of adjusted trades '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, turnoverRaw, marking, err := parseFloat(parts[8], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse turnover '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, turnoverAdj, marking, err := parseFloat(parts[18], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse adjusted turnover '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, vwapRaw, marking, err := parseFloat(parts[9], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse vwap '%s' in line '%s': %w", i+1, s, line, err)
		}

		s, vwapAdj, marking, err := parseFloat(parts[19], marking)
		if err != nil {
			return combinedHist, fmt.Errorf("line %d: cannot parse adjusted vwap '%s' in line '%s': %w", i+1, s, line, err)
		}

		factor := 1.
		if closeRaw != closeAdj && closeRaw != 0 {
			factor = closeAdj / closeRaw
		}

		entry := euronext.CombinedDailyHistory{
			Date:                   time,
			Open:                   openRaw,
			High:                   highRaw,
			Low:                    lowRaw,
			Last:                   lastRaw,
			Close:                  closeRaw,
			NumberOfShares:         sharesRaw,
			NumberOfTrades:         tradesRaw,
			Turnover:               turnoverRaw,
			Vwap:                   vwapRaw,
			OpenAdjusted:           openAdj,
			HighAdjusted:           highAdj,
			LowAdjusted:            lowAdj,
			LastAdjusted:           lastAdj,
			CloseAdjusted:          closeAdj,
			NumberOfSharesAdjusted: sharesAdj,
			NumberOfTradesAdjusted: tradesAdj,
			TurnoverAdjusted:       turnoverAdj,
			VwapAdjusted:           vwapAdj,
			AdjustmentFactor:       factor,
			HasMarking:             marking,
		}
		combinedHist = append(combinedHist, entry)
	}

	return combinedHist, nil
}

func parseFloat(s string, marking bool) (string, float64, bool, error) {
	s, marking = cleanString(s, marking)
	if len(s) == 0 || s == "0" || s == "0.0" {
		return s, 0, marking, nil
	}
	v, err := strconv.ParseFloat(s, 64)
	return s, v, marking, err
}

func cleanString(s string, marking bool) (string, bool) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "'") {
		s = s[1:]
		marking = true
	}
	return s, marking
}
