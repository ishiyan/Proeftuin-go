package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type oecd struct {
	date      time.Time
	value     float64
	inflation float64
}

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	if len(fileName) < 1 {
		fmt.Print("expecting input OECD CSV file name as an argument")
		return
	}

	s := readOecd(fileName)

	fileName += ".converted"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Panic("Failed to create file '" + fileName + "'")
	}
	defer f.Close()

	var id int

	for _, e := range s {
		id++
		line := fmt.Sprintf("    (%v,'NL','%v',%v,%.4f),\n", id, e.date.Format("2006-01-02"), e.value, e.inflation)
		f.WriteString(line)
	}

	for _, e := range s {
		id++
		line := fmt.Sprintf("    (%v,'BE','%v',%v,%.4f),\n", id, e.date.Format("2006-01-02"), e.value, e.inflation)
		f.WriteString(line)
	}

	for _, e := range s {
		id++
		line := fmt.Sprintf("    (%v,'SG','%v',%v,%.4f),\n", id, e.date.Format("2006-01-02"), e.value, e.inflation)
		f.WriteString(line)
	}
}

func readOecd(filePath string) []oecd {
	split := readCsvFile(filePath, 3)
	s := make([]oecd, 0)

	// 2022-11	128.8686	0.17%
	for _, e := range split {
		if len(e) != 3 {
			log.Panic("Expecterd 3 splitted parts")
		}

		d, err := time.Parse("2006-01", e[0])
		if err != nil {
			log.Panic("Failed to parse year-month '" + e[0] + "'")
		}

		firstOfMonth := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.Local)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		value, err := strconv.ParseFloat(e[1], 64)
		if err != nil {
			log.Panic("Failed to parse value '" + e[1] + "'")
		}

		inflation, err := strconv.ParseFloat(strings.TrimSuffix(e[2], "%"), 64)
		if err != nil {
			log.Panic("Failed to parse trimmed '" + strings.TrimSuffix(e[2], "%") + "'")
		}

		v := oecd{
			date:      lastOfMonth,
			value:     value,
			inflation: inflation / 100,
		}

		s = append(s, v)
	}

	return s
}

func readCsvFile(filePath string, numFields int) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Panic("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'
	r.TrimLeadingSpace = true
	r.Comment = '#'
	r.FieldsPerRecord = numFields

	records, err := r.ReadAll()
	if err != nil {
		log.Panic("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
