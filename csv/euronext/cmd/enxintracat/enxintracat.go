package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"euronext/euronext"
	"euronext/euronext/intraday"
)

const configFileName = "enxintracat.json"

type config struct {
	DownloadsFolder string `json:"downloadsFolder"`
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

	cfg, err := readConfig(configFileName)
	if err != nil {
		log.Panicf("cannot read configuration file %s: %s\n", configFileName, err)
	}

	err = euronext.EnsureDirectoryExists(cfg.DownloadsFolder)
	if err != nil {
		log.Panicf("cannot create downloads directory %s: %s\n", cfg.DownloadsFolder, err)
	}

	mp := intraday.NewTradeLabelMap()

	err = filepath.WalkDir(cfg.DownloadsFolder, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil // skip directories, only add files
		}

		jsInd, err := intraday.ReadJsonIntradayFile(path)
		if jsInd == nil {
			e := fmt.Errorf("error: cannot read json file %s: %w", path, err)
			log.Println(e)
			return nil
		}

		if len(jsInd.TradeTypeList) > 0 {
			for _, tradeType := range jsInd.TradeTypeList {
				if _, ok := mp[strings.ToLower(tradeType.Label)]; !ok {
					log.Printf("** new trade type '%s', code: '%s', idNXT: '%s', %v\n", tradeType.Label, tradeType.Code, tradeType.IDNXT, tradeType)
					mp[strings.ToLower(tradeType.Label)] = ""
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Panicf("error walking through directory %s: %s\n", cfg.DownloadsFolder, err)
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

	if !strings.HasSuffix(conf.DownloadsFolder, "/") {
		conf.DownloadsFolder += "/"
	}

	return &conf, nil
}
