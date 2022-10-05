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

export const tradeSeries_TODO_Aperiodic: Series = {
  mnemonic: '_TODO_',
  description: '_TODO_.',
  timeStart: new Date(_TODO_),
  timeEnd: new Date(_TODO_),
  timeGranularity: TimeGranularity.Aperiodic,
  data: [
`

const footer = `  ],
};`

func main() {
	tformatPtr := flag.String("tformat", "2006/01/02 15:04:05.9999999", "csv time format in go style")
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

		if len(rec) < 2 {
			fail(fmt.Sprintf("line %d: expected at least 2 parts, got %d", lineNo, len(rec)))
		}

		t, err := time.Parse(*tformatPtr, rec[0])
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse time part '%s' using format '%s': %s", lineNo, rec[0], *tformatPtr, err))
		}

		if t0.After(t) {
			fail(fmt.Sprintf("line %d: time part '%s' time '%v' is before previous line time '%v'", lineNo, rec[0], t, t0))
		}

		t0 = t

		p, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse price part '%s': %s", lineNo, rec[1], err))
		}

		var v float64 = float64(*volumePtr)

		if len(rec) > 2 {
			v, err = strconv.ParseFloat(rec[2], 64)
			if err != nil {
				fail(fmt.Sprintf("line %d: failed to parse volume part '%s': %s", lineNo, rec[2], err))
			}
		}

		m := t.Nanosecond() / 100
		if m > 999 {
			// m = 999
			fail(fmt.Sprintf("line %d: too many millisecond ticks in fractional time part: '%s', JavaScript does not support this", lineNo, rec[0]))
		}

		writeString(fout, fmt.Sprintf(
			"    { time: new Date(%d, %d, %d, %d, %d, %d, %d), price: %v, volume: %v },\n",
			t.Year(), t.Month()-1, t.Day(), t.Hour(), t.Minute(), t.Second(), m,
			p, v))
		lineNo++
	}

	writeString(fout, footer)
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
