package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"euronext/euronext"
	"euronext/euronext/discovery"
	"euronext/euronext/enrichment"
)

const configFileName = "enxdisc.json"

type config struct {
	DownloadRetries             int    `json:"downloadRetries"`
	DownloadTimeoutSec          int    `json:"downloadTimeoutSec"`
	DownloadPauseBeforeRetrySec int    `json:"downloadPauseBeforeRetrySec"`
	UserAgent                   string `json:"userAgent"`
	ZipDownloadedFolder         bool   `json:"zipDownloadedFolder"`
	DeleteDownloadedFolder      bool   `json:"deleteDownloadedFolder"`
	EnrichDiscoveredInstruments bool   `json:"enrichDiscoveredInstruments"`
	GzipBackupXmlInstruments    bool   `json:"gzipBackupXmlInstruments"`
	DownloadsFolder             string `json:"downloadsFolder"`
	RepositoryFolder            string `json:"repositoryFolder"`
	XmlInstrumntsFile           string `json:"xmlInstrumentsFile"`
	XmlInstrumntsFileOther      string `json:"xmlInstrumentsFileOther"`
}

func main() {
	now := time.Now()
	t := now.Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("enxdisc_%s.log", t)
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Panicf("cannot create log file '%s': %s", logFileName, err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	cfg, err := readConfig(configFileName)
	if err != nil {
		log.Panicf("cannot read configuration file %s: %s", configFileName, err)
	}

	err = ensureDirectoryExists(cfg.RepositoryFolder)
	if err != nil {
		log.Panicf("cannot create repository directory %s: %s", cfg.RepositoryFolder, err)
	}

	err = ensureDirectoryExists(cfg.DownloadsFolder)
	if err != nil {
		log.Panicf("cannot create downloads directory %s: %s", cfg.DownloadsFolder, err)
	}

	log.Println("xml file: " + cfg.XmlInstrumntsFile)
	if err := backupXmlFile(cfg.XmlInstrumntsFile, now, cfg.GzipBackupXmlInstruments); err != nil {
		log.Panicf("cannot backup xml file: %s", err)
	}

	instruments, err := euronext.ReadXmlInstrumentsFile(cfg.XmlInstrumntsFile)
	if err != nil {
		log.Panicf("cannot read instruments: %s", err)
	}

	log.Println("xml file: " + cfg.XmlInstrumntsFileOther)
	if err := backupXmlFile(cfg.XmlInstrumntsFileOther, now, cfg.GzipBackupXmlInstruments); err != nil {
		log.Panicf("cannot backup xml file: %s", err)
	}

	instrumentsOther, err := euronext.ReadXmlInstrumentsFile(cfg.XmlInstrumntsFileOther)
	if err != nil {
		log.Panicf("cannot read instruments: %s", err)
	}

	log.Printf("found %d instruments in %s\n", len(instruments.Instrument), cfg.XmlInstrumntsFile)
	log.Printf("found %d instruments in %s\n", len(instrumentsOther.Instrument), cfg.XmlInstrumntsFileOther)
	instruments.Instrument = append(instruments.Instrument, instrumentsOther.Instrument...)
	log.Printf("total instruments: %d\n", len(instruments.Instrument))

	actualInstrumentsMap := discovery.Fetch(
		cfg.DownloadsFolder,
		now,
		cfg.DownloadRetries,
		cfg.DownloadTimeoutSec,
		cfg.DownloadPauseBeforeRetrySec,
		cfg.ZipDownloadedFolder,
		cfg.DeleteDownloadedFolder,
		cfg.UserAgent)

	log.Printf("fetched %d actual instruments\n", len(actualInstrumentsMap))

	newInstruments := make([]euronext.XmlInstrument, 0)
	newInstrumentsOther := make([]euronext.XmlInstrument, 0)

	for _, ai := range actualInstrumentsMap {
		matched := matchIsinMicMnemonic(instruments, ai)
		ai.IsApproved = len(matched) > 0
		if len(matched) > 1 {
			log.Printf("found %d duplicate matches for %s-%s-%s:\n", len(matched), ai.Mic, ai.Symbol, ai.Isin)
			for i, m := range matched {
				log.Printf("  %d: %s-%s-%s\n", i, m.Mic, m.Symbol, m.Isin)
			}
		}

		matchedOther := matchIsinMicMnemonic(instrumentsOther, ai)
		ai.IsDiscovered = len(matchedOther) > 0
		if len(matchedOther) > 1 {
			log.Printf("found %d duplicate matches for %s-%s-%s in other instruments:\n", len(matchedOther), ai.Mic, ai.Symbol, ai.Isin)
			for i, m := range matchedOther {
				log.Printf("  %d: %s-%s-%s\n", i, m.Mic, m.Symbol, m.Isin)
			}
		}

		if cfg.EnrichDiscoveredInstruments && !ai.IsApproved && !ai.IsDiscovered {
			log.Printf("discovered %s: \"%s\":", ai.Type, ai.Key)

			xmlIns := euronext.NewXmlInstrument(ai.Type, ai.Mic, ai.Isin, ai.Symbol, ai.Mep)
			if xmlIns == nil {
				log.Printf("invalid instrument type '%s' for %s-%s-%s, skipping\n", ai.Type, ai.Mic, ai.Symbol, ai.Isin)
				continue
			}

			enrichment.EnrichInstrument(xmlIns, cfg.DownloadRetries, cfg.DownloadTimeoutSec, cfg.DownloadPauseBeforeRetrySec, cfg.UserAgent)

			if contains(discovery.KnownEuronextMics, ai.Mic) {
				newInstruments = append(newInstruments, *xmlIns)
			} else {
				newInstrumentsOther = append(newInstrumentsOther, *xmlIns)
			}
		}
	}

	log.Printf("adding %d new instruments to %s\n", len(newInstruments), cfg.XmlInstrumntsFile)
	if err := appendToXmlInstrumentsFile(cfg.XmlInstrumntsFile, now, newInstruments); err != nil {
		log.Printf("cannot append new instruments to xml file %s: %s", cfg.XmlInstrumntsFile, err)
	}

	log.Printf("adding %d new instruments to %s\n", len(newInstrumentsOther), cfg.XmlInstrumntsFileOther)
	if err := appendToXmlInstrumentsFile(cfg.XmlInstrumntsFileOther, now, newInstrumentsOther); err != nil {
		log.Printf("cannot append new instruments to xml file %s: %s", cfg.XmlInstrumntsFileOther, err)
	}

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

	if !strings.HasSuffix(conf.RepositoryFolder, "/") {
		conf.RepositoryFolder += "/"
	}

	return &conf, nil
}

func ensureDirectoryExists(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = os.MkdirAll(directory, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory '%s': %w", directory, err)
		}
	}

	return nil
}

func backupXmlFile(filePath string, dateTime time.Time, doGzip bool) error {
	suffix := dateTime.Format("20060102_150405")
	filePathBackup := filePath + "." + suffix + ".xml"
	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if doGzip {
		filePathBackup += ".gz"
		var buf bytes.Buffer
		w := gzip.NewWriter(&buf)
		if _, err := w.Write(input); err != nil {
			return err
		}
		if err := w.Close(); err != nil {
			return err
		}
		input = buf.Bytes()
	}
	return os.WriteFile(filePathBackup, input, 0644) // owner: read/write, group/others: read
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if strings.EqualFold(v, item) {
			return true
		}
	}
	return false
}

func matchIsinMicMnemonic(existing *euronext.XmlInstruments, discovered *discovery.InstrumentInfo) []*discovery.InstrumentInfo {
	matched := make([]*discovery.InstrumentInfo, 0)
	for _, ei := range existing.Instrument {
		if strings.EqualFold(ei.Isin, discovered.Isin) && strings.EqualFold(ei.Mic, discovered.Mic) && strings.EqualFold(ei.Symbol, discovered.Symbol) {
			matched = append(matched, discovered)
		}
	}
	return matched
}

func appendToXmlInstrumentsFile(filePath string, now time.Time, instruments []euronext.XmlInstrument) error {
	lines := now.Format("  <!-- 20060102_150405 -->") + "\n"
	for _, ins := range instruments {
		lines += euronext.XmlInstrumentToXmlString(&ins)
	}
	lines += "</instruments>\n"

	// Read the entire file into memory
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("cannot read file '%s': %w", filePath, err)
	}

	closingTag := []byte("</instruments>")
	idx := bytes.LastIndex(content, closingTag)
	if idx == -1 {
		return fmt.Errorf("no </instruments> closing tag found in file '%s'", filePath)
	}

	// Prepare the new content
	var buf bytes.Buffer
	buf.Write(content[:idx]) // everything before the closing tag
	buf.WriteString(lines)

	// Write back to the file (overwrite)
	err = os.WriteFile(filePath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("cannot write file '%s': %w", filePath, err)
	}

	return nil
}
