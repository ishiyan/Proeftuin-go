package eurofxref

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const csvTimeFormat = "2006/01/02"

func ensureDirectoryExists(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err = os.MkdirAll(directory, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory '%s': %w", directory, err)
		}
	}

	return nil
}

func WhatMnemonic(what What) string {
	switch what {
	case EurFxRefLast:
		return "last rates"
	case EurFxRef90:
		return "90d rates"
	case EurFxRefFull:
		return "full rates"
	default:
		return "unknown"
	}
}

func filePath(repository, currency string) string {
	return repository + "eurfxref." + currency + ".csv"
}

func ReadCSV(repository, currency string) ([]Point, error) {
	var f *os.File
	var err error

	if err = ensureDirectoryExists(repository); err != nil {
		return nil, err
	}

	series := make([]Point, 0)
	file := filePath(repository, currency)

	if _, err = os.Stat(file); os.IsNotExist(err) {
		if f, err = os.Create(file); err != nil {
			return nil, fmt.Errorf("cannot create file '%s': %w", file, err)
		} else {
			f.Close()
			return series, nil
		}
	}

	if f, err = os.Open(file); err != nil {
		return nil, fmt.Errorf("cannot open file '%s': %w", file, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comment = '#'
	csvReader.Comma = ';'
	csvReader.ReuseRecord = true

	t0 := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	lineNo := 0

	for {
		rec, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("error reading line %d: %w", lineNo, err)
		}

		if len(rec) < 2 {
			return nil, fmt.Errorf("line %d: expected at least 2 parts, got %d", lineNo, len(rec))
		}

		t, err := time.Parse(csvTimeFormat, rec[0])
		if err != nil {
			return nil, fmt.Errorf("line %d: failed to parse time part '%s' using format '%s': %w", lineNo, rec[0], csvTimeFormat, err)
		}

		if t0.After(t) {
			return nil, fmt.Errorf("line %d: time part '%s' time '%v' is before previous line time '%v'", lineNo, rec[0], t, t0)
		}

		t0 = t

		v, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return nil, fmt.Errorf("line %d: failed to parse value part '%s': %w", lineNo, rec[1], err)
		}

		lineNo++
		series = append(series, Point{
			Date:  t,
			Value: v,
		})
	}

	return series, nil
}

func WriteCSV(repository, currency string, points []Point) error {
	file := filePath(repository, currency)
	backPath := file + ".bak"

	if err := os.Rename(file, backPath); err != nil {
		return fmt.Errorf("cannot rename file '%s' to  '%s': %w", file, backPath, err)
	}

	if fout, err := os.Create(file); err != nil {
		return fmt.Errorf("cannot create file '%s': %w", file, err)
	} else {
		defer fout.Close()

		for _, p := range points {
			s := fmt.Sprintf("%s;%v\n", p.Date.Format(csvTimeFormat), p.Value)
			if _, err := fout.WriteString(s); err != nil {
				return fmt.Errorf("cannot write file: %w", err)
			}
		}
	}

	return nil
}
