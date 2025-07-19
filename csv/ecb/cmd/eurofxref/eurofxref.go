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
	"time"

	"ecb/internal/eurofxref"
)

const configFileName = "eurofxref.json"

type config struct {
	LastDay                     bool   `json:"lastDay"`
	Last90DayHistory            bool   `json:"last90DayHistory"`
	FullHistory                 bool   `json:"fullHistory"`
	RepositoryFolder            string `json:"repositoryFolder"`
	DownloadsFolder             string `json:"downloadsFolder"`
	ZipDownloadedFolder         bool   `json:"zipDownloadedFolder"`
	DeleteDownloadedFolder      bool   `json:"deleteDownloadedFolder"`
	VerboseDownload             bool   `json:"verboseDownload"`
	DownloadRetryDelaySeconds   []int  `json:"downloadRetryDelaySeconds"`
	DownloadTimeoutSeconds      int    `json:"downloadTimeoutSeconds"`
	UserAgent                   string `json:"userAgent"`
	DownloadRetryDelayDurations []time.Duration
	DownloadTimeoutDuration     time.Duration
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

func updateSeries(repository string, what eurofxref.What, cfg *config, downloadPath string) error {
	if psm, err := eurofxref.Fetch(what, true, downloadPath, cfg.DownloadTimeoutDuration,
		cfg.DownloadRetryDelayDurations, cfg.UserAgent, cfg.VerboseDownload); err != nil {
		return fmt.Errorf("cannot download: %w", err)
	} else {
		for currency, pts2 := range psm {
			if len(pts2) == 0 {
				continue // Skip empty series
			}

			log.Printf("Updating %s series for currency %s\n", eurofxref.WhatMnemonic(what), currency)

			pts1, err := eurofxref.ReadCSV(repository, currency)
			if err != nil {
				log.Println("cannot read csv file, skipping: %w", err)
				continue
			}

			date := time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})
			if len(pts1) > 0 {
				date = pts1[len(pts1)-1].Date.Add(24 * time.Hour)
			}

			for _, p := range pts2 {
				if p.Date.Before(date) {
					continue
				}
				pts1 = append(pts1, p)
			}

			if err = eurofxref.WriteCSV(repository, currency, pts1); err != nil {
				log.Println("cannot write csv file, skipping: %w", err)
				continue
			}
		}
		return nil
	}
}

func main() {
	now := time.Now()
	t := now.Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("eurofxref_%s.log", t)
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

	downloadName := now.Format("20060102")
	downloadPath := cfg.DownloadsFolder + now.Format("2006") + "/" + downloadName + "/"
	log.Println("downloading to " + downloadPath)

	log.Println("full history:", cfg.FullHistory)
	log.Println("last 90 day history:", cfg.Last90DayHistory)
	log.Println("last day:", cfg.LastDay)
	log.Println("repository folder:", cfg.RepositoryFolder)
	log.Println("download folder:", cfg.DownloadsFolder)
	log.Println("download retry delay seconds:", cfg.DownloadRetryDelaySeconds)
	log.Println("download timeout seconds:", cfg.DownloadTimeoutSeconds)
	log.Println("verbose download:", cfg.VerboseDownload)
	log.Println("zip download folder:", cfg.ZipDownloadedFolder)
	log.Println("delete download folder:", cfg.DeleteDownloadedFolder)
	log.Println("=======================================")

	if cfg.FullHistory {
		log.Println("Updating full history...")
		if err = updateSeries(cfg.RepositoryFolder, eurofxref.EurFxRefFull, cfg, downloadPath); err != nil {
			log.Printf("%s: %s\n", eurofxref.WhatMnemonic(eurofxref.EurFxRefFull), err)
		}
	}

	if cfg.Last90DayHistory {
		log.Println("Updating last 90 days history...")
		if err = updateSeries(cfg.RepositoryFolder, eurofxref.EurFxRef90, cfg, downloadPath); err != nil {
			log.Printf("%s: %s\n", eurofxref.WhatMnemonic(eurofxref.EurFxRef90), err)
		}
	}

	if cfg.LastDay {
		log.Println("Updating last day...")
		if err = updateSeries(cfg.RepositoryFolder, eurofxref.EurFxRefLast, cfg, downloadPath); err != nil {
			log.Printf("%s: %s\n", eurofxref.WhatMnemonic(eurofxref.EurFxRefLast), err)
		}
	}
	log.Println("processed")
	log.Println("=======================================")
	archive(downloadPath, cfg.ZipDownloadedFolder, cfg.DeleteDownloadedFolder)
	log.Println("finished")
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
		file := fmt.Sprintf("%seurofxref", downloadFolder)
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
