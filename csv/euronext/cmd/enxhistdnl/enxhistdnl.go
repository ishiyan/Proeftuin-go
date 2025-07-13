package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"euronext/euronext"
	"euronext/euronext/endofday"
)

const configFileName = "enxhistdnl.json"

type config struct {
	XmlInstrumntsFile           string `json:"xmlInstrumentsFile"`
	DownloadsFolder             string `json:"downloadsFolder"`
	ZipDownloadedFolder         bool   `json:"zipDownloadedFolder"`
	DeleteDownloadedFolder      bool   `json:"deleteDownloadedFolder"`
	VerboseDownload             bool   `json:"verboseDownload"`
	DownloadRetryDelaySeconds   []int  `json:"downloadRetryDelaySeconds"`
	DownloadTimeoutSeconds      int    `json:"downloadTimeoutSeconds"`
	Concurrency                 int    `json:"concurrency"`
	UserAgent                   string `json:"userAgent"`
	DownloadRetryDelayDurations []time.Duration
	DownloadTimeoutDuration     time.Duration
}

type instrument struct {
	Mnemonic string `json:"mnemonic"`
	Mep      string `json:"mep"`
	Mic      string `json:"mic"`
	Isin     string `json:"isin"`
	Type     string `json:"type"`
}

type statistics struct {
	DownloadErrors []string
	FileErrors     []string
}

func main() {
	now := time.Now()
	t := now.Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("enxhistdnl_%s.log", t)
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Panicf("cannot create log file '%s': %s\n", logFileName, err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	cfg, err := readConfig(configFileName)
	if err != nil {
		log.Panicf("cannot read configuration file %s: %s\n", configFileName, err)
	}

	log.Println("xml instruments file:", cfg.XmlInstrumntsFile)
	log.Println("download folder:", cfg.DownloadsFolder)
	log.Println("download retry delay seconds:", cfg.DownloadRetryDelaySeconds)
	log.Println("download timeout seconds:", cfg.DownloadTimeoutSeconds)
	log.Println("download concurrency:", cfg.Concurrency)
	log.Println("verbose download:", cfg.VerboseDownload)
	log.Println("zip download folder:", cfg.ZipDownloadedFolder)
	log.Println("delete download folder:", cfg.DeleteDownloadedFolder)

	sessionDate, err := euronext.SessionDate()
	if err != nil {
		log.Panicf("cannot get session date: %s\n", err)
	}
	log.Println("trading session date: " + sessionDate.Format("2006-01-02"))

	err = euronext.EnsureDirectoryExists(cfg.DownloadsFolder)
	if err != nil {
		log.Panicf("cannot create downloads directory %s: %s\n", cfg.DownloadsFolder, err)
	}

	downloadName := sessionDate.Format("20060102")
	downloadPath := cfg.DownloadsFolder + "endofday/" +
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
		DownloadErrors: []string{"mep;mic;type;mnemonic;isin;adj;error"},
		FileErrors:     []string{"mep;mic;type;mnemonic;isin;adj;error"},
	}

	log.Println("=======================================")
	log.Println("downloading to " + downloadPath)

	l := len(instruments)
	if cfg.Concurrency < 2 {
		for i, ins := range instruments {
			ins.download(cfg, i, l, downloadPath, downloadName, &stati, false)
			ins.download(cfg, i, l, downloadPath, downloadName, &stati, true)
		}
	} else {
		var wg sync.WaitGroup
		sem := make(chan struct{}, cfg.Concurrency)

		for i, ins := range instruments {
			wg.Add(1)
			sem <- struct{}{} // Acquire a slot

			go func(i, l int, ins instrument) {
				defer wg.Done()
				defer func() { <-sem }() // Release the slot

				ins.download(cfg, i, l, downloadPath, downloadName, &stati, false)
				ins.download(cfg, i, l, downloadPath, downloadName, &stati, true)
			}(i, l, ins)
		}

		// Wait for the first pipeline to finish processing
		wg.Wait()
	}

	log.Println("processed")
	log.Println("=======================================")
	log.Printf("instruments with download errors: %d from %d\n", len(stati.DownloadErrors)-1, l)
	for _, z := range stati.DownloadErrors {
		log.Println(z)
	}

	log.Println("=======================================")
	log.Printf("instruments with file save errors: %d from %d\n", len(stati.FileErrors)-1, l)
	for _, z := range stati.FileErrors {
		log.Println(z)
	}

	log.Println("=======================================")
	archive(downloadPath, cfg.ZipDownloadedFolder, cfg.DeleteDownloadedFolder)
	log.Println("finished")
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

	if conf.Concurrency < 2 {
		conf.Concurrency = 0
	} else if conf.Concurrency > 8 {
		conf.Concurrency = 8
	}

	if conf.DownloadTimeoutSeconds < 1 {
		conf.DownloadTimeoutSeconds = 1
	}
	conf.DownloadTimeoutDuration = time.Duration(conf.DownloadTimeoutSeconds) * time.Second

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

func (s *instrument) fileName() string {
	return fmt.Sprintf("%s_%s_%s", s.Mic, s.Mnemonic, s.Isin)
}

func (s *instrument) download(
	cfg *config,
	el int,
	elen int,
	downloadFolder string,
	downloadName string,
	sta *statistics,
	isAdjusted bool,
) {
	adj := "unadj"
	if isAdjusted {
		adj = "adj"
	}
	insName := s.fileName()
	prefix := fmt.Sprintf("(%d of %d) %s %s %s: downloading ... ", el+1, elen, insName, s.Type, adj)

	bs, err := endofday.FetchEndofdayData(s.Isin, s.Mic, s.Mnemonic, s.Type,
		cfg.DownloadTimeoutDuration, cfg.DownloadRetryDelayDurations, cfg.UserAgent,
		cfg.VerboseDownload, isAdjusted)
	if err != nil {
		e := err.Error()
		sta.DownloadErrors = append(sta.DownloadErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, adj, e))
		log.Println(prefix + "failed: " + e)
		return
	}

	prefix = fmt.Sprintf("%ssaving %d bytes ... ", prefix, len(bs))
	adj2 := "_unadjusted_"
	if isAdjusted {
		adj2 = "_adjusted_"
	}

	csvFile := filepath.Join(downloadFolder,
		strings.ToUpper(insName)+adj2+downloadName+".csv")
	err = os.WriteFile(csvFile, bs, 0644)
	if err != nil {
		e := err.Error()
		sta.FileErrors = append(sta.FileErrors,
			fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s", s.Mep, s.Mic, s.Type, s.Mnemonic, s.Isin, adj, e))
		log.Println(prefix + "failed: " + e)
	} else {
		log.Println(prefix + "done")
	}
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

func archive(downloadFolder string, zipDownloadedFolder, deleteDownloadedFolder bool) {
	downloadFolder = strings.TrimSuffix(downloadFolder, "/")

	if zipDownloadedFolder {
		file := fmt.Sprintf("%senx_eod", downloadFolder)
		fz := file + ".zip"
		counter := 0
	again:
		_, err := os.Stat(fz)
		if err == nil {
			counter++
			fz = fmt.Sprintf("%s.%d.zip", file, counter)
			goto again
		}
		prefix := fmt.Sprintf("archiving from %s to %s ... ", downloadFolder, fz)

		if err := zipFolder(downloadFolder, fz); err != nil {
			log.Println(prefix + "failed: " + err.Error())
			return
		} else {
			log.Println(prefix + "done")
		}
	}

	if deleteDownloadedFolder {
		prefix := fmt.Sprintf("deleting folder %s ... ", downloadFolder)

		if err := os.RemoveAll(downloadFolder); err != nil {
			log.Println(prefix + "failed: " + err.Error())
		} else {
			log.Println(prefix + "done")
		}
	}
}
