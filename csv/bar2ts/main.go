package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const header = `import { TimeGranularity } from 'projects/mb/src/public-api';
import { Series } from '../../series.interface';

export const barSeries_TODO_1d: Series = {
  mnemonic: '_TODO_',
  description: '_TODO_.',
  timeStart: new Date(_TODO_),
  timeEnd: new Date(_TODO_),
  timeGranularity: TimeGranularity.Day1,
  data: [
`

const footer = `  ],
};`

func main() {
	tformatPtr := flag.String("tformat", "2006/01/02 15:04:05.9999999", "csv time format in go style")
	tgranPtr := flag.String("tgran", "day", "time granularity of bars: {year, month, week, day, hour, min}")
	volumePtr := flag.Int("volume", 0, "volume value if not present")
	headerPtr := flag.Bool("header", true, "the very first CSV line is a header")

	flag.Parse()

	var fin *os.File
	var fout *os.File
	var err error

	if filename := flag.Arg(0); filename == "" {
		fail("expecting CSV file name as the positional argument")
	} else {
		if !strings.HasSuffix(filename, ".csv") {
			fail(fmt.Sprintf("expecting CSV file name to end with '.csv': %s", filename))
		}

		fin, err = os.Open(filename)
		if err != nil {
			fail(fmt.Sprintf("error opening file: %s", err))
		}
		defer fin.Close()

		fout, err = os.Create(filename + ".ts")
		if err != nil {
			fail(fmt.Sprintf("error creating file: %s", err))
		}
		defer fout.Close()
	}

	writeString(fout, header)

	csvReader := csv.NewReader(fin)
	csvReader.Comment = '#'
	csvReader.Comma = ';'

	t0 := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	lineNo := 0

	for {
		rec, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			fail(fmt.Sprintf("error reading line %d: %s", lineNo, err))
		}

		if lineNo == 0 && *headerPtr {
			lineNo++
			continue
		}

		if len(rec) < 5 {
			fail(fmt.Sprintf("line %d: expected at least 5 parts, got %d", lineNo, len(rec)))
		}

		t, err := time.Parse(*tformatPtr, rec[0])
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse time part '%s' using format '%s': %s", lineNo, rec[0], *tformatPtr, err))
		}

		if t0.After(t) {
			fail(fmt.Sprintf("line %d: time part '%s' time '%v' is before previous line time '%v'", lineNo, rec[0], t, t0))
		}

		t0 = t

		op, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse opening price part '%s': %s", lineNo, rec[1], err))
		}

		hp, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse highest price part '%s': %s", lineNo, rec[2], err))
		}

		lp, err := strconv.ParseFloat(rec[3], 64)
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse lowest price part '%s': %s", lineNo, rec[3], err))
		}

		cp, err := strconv.ParseFloat(rec[4], 64)
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse closing price part '%s': %s", lineNo, rec[4], err))
		}

		if op > hp || lp > hp || cp > hp {
			fail(fmt.Sprintf("line %d: high price '%v' is not the highest: %+v", lineNo, hp, rec))
		}

		if op < lp || hp < lp || cp < lp {
			fail(fmt.Sprintf("line %d: low price '%v' is not the lowest: %+v", lineNo, lp, rec))
		}

		if op <= 0 || lp <= 0 || hp <= 0 || cp <= 0 {
			fail(fmt.Sprintf("line %d: price should be positive: %+v", lineNo, rec))
		}

		var v float64 = float64(*volumePtr)

		if len(rec) > 5 {
			v, err = strconv.ParseFloat(rec[5], 64)
			if err != nil {
				fail(fmt.Sprintf("line %d: failed to parse volume part '%s': %s", lineNo, rec[5], err))
			}
		}

		writeString(fout, fmt.Sprintf(
			"    { time: %s, open: %v, high: %v, low: %v, close: %v, volume: %.f },\n",
			time2Ts(t, *tgranPtr), op, hp, lp, cp, v))
		lineNo++
	}

	writeString(fout, footer)
}

func time2Ts(t time.Time, g string) string {
	switch g {
	case "year", "month":
		return fmt.Sprintf("new Date(%d, %d)", t.Year(), t.Month()-1)
	case "week", "day":
		return fmt.Sprintf("new Date(%d, %d, %d)", t.Year(), t.Month()-1, t.Day())
	case "hour":
		return fmt.Sprintf("new Date(%d, %d, %d, %d)", t.Year(), t.Month()-1, t.Day(), t.Hour())
	case "min":
		return fmt.Sprintf("new Date(%d, %d, %d, %d, %d)", t.Year(), t.Month()-1, t.Day(), t.Hour(), t.Minute())
	default:
		fail(fmt.Sprintf("unknown time granularity '%s', expection one of {year, month, week, day, hour, min}", g))
		return ""
	}
}

func writeString(f *os.File, s string) {
	if _, err := f.WriteString(s); err != nil {
		fail(fmt.Sprintf("error writing file: %s", err))
	}
}

func fail(s string) {
	fmt.Print(s)
	os.Exit(1)
}
