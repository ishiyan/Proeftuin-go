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

	estr "estr/estr"
)

const configFileName = "estr.json"

type config struct {
	Actual                      bool   `json:"actual"`
	Pre                         bool   `json:"pre"`
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

func downloadSeries(what estr.What, startDate time.Time, cfg *config, downloadPath string) ([]estr.Point, error) {
	if pts, err := estr.Fetch(what, true, downloadPath, cfg.DownloadTimeoutDuration,
		cfg.DownloadRetryDelayDurations, cfg.UserAgent, cfg.VerboseDownload); err != nil {
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

func updateSeries(repository string, what estr.What, cfg *config, downloadPath string) error {
	s1, err := estr.ReadCSV(repository, what)
	if err != nil {
		return fmt.Errorf("cannot read csv file: %w", err)
	}

	date := time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})
	if len(s1) > 0 {
		date = s1[len(s1)-1].Date.Add(24 * time.Hour)
	}

	s2, err := downloadSeries(what, date, cfg, downloadPath)
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
	now := time.Now()
	t := now.Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("estr_%s.log", t)
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

	log.Println("actual:", cfg.Actual)
	log.Println("pre:", cfg.Pre)
	log.Println("repository folder:", cfg.RepositoryFolder)
	log.Println("download folder:", cfg.DownloadsFolder)
	log.Println("download retry delay seconds:", cfg.DownloadRetryDelaySeconds)
	log.Println("download timeout seconds:", cfg.DownloadTimeoutSeconds)
	log.Println("verbose download:", cfg.VerboseDownload)
	log.Println("zip download folder:", cfg.ZipDownloadedFolder)
	log.Println("delete download folder:", cfg.DeleteDownloadedFolder)
	log.Println("=======================================")

	if cfg.Pre {
		log.Println("Updating pre-series...")
		for _, w := range []estr.What{estr.EstrRatePre, estr.EstrVolumePre, estr.EstrTransactionsPre} {
			if err = updateSeries(cfg.RepositoryFolder, w, cfg, downloadPath); err != nil {
				log.Printf("%s: %s\n", estr.WhatMnemonic(w), err)
			}
		}
	}

	if cfg.Actual {
		log.Println("Updating actual series...")
		for _, w := range []estr.What{estr.EstrRateAct, estr.EstrVolumeAct, estr.EstrTransactionsAct} {
			if err = updateSeries(cfg.RepositoryFolder, w, cfg, downloadPath); err != nil {
				log.Printf("%s: %s\n", estr.WhatMnemonic(w), err)
			}
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
		file := fmt.Sprintf("%sestr", downloadFolder)
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
