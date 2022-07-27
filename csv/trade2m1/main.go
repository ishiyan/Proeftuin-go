package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	tformatPtr := flag.String("tformat", "2006/01/02 15:04:05.9999999", "csv time format in go style")
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

		fout, err = os.Create(filename + ".m1.csv")
		if err != nil {
			fail(fmt.Sprintf("error creating file: %s", err))
		}
		defer fout.Close()
	}

	csvReader := csv.NewReader(fin)
	csvReader.Comment = '#'
	csvReader.Comma = ';'

	t0 := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	lineNo, m := 0, -1
	co, ch, cl, cc, cv := 0., 0., 0., 0., 0.

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

		if len(rec) < 3 {
			fail(fmt.Sprintf("line %d: expected at least 3 parts, got %d", lineNo, len(rec)))
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

		v, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			fail(fmt.Sprintf("line %d: failed to parse volume part '%s': %s", lineNo, rec[2], err))
		}

		if m < 0 { // initial assignment
			m = t.Minute()
			co, ch, cl, cc = p, p, p, p
			cv = v
		} else {
			if m == t.Minute() { // the same minute, update bar
				ch = math.Max(ch, p)
				cl = math.Min(cl, p)
				cc = p
				cv += v
			} else { // save bar and initialize new minute
				writeString(fout, fmt.Sprintf("%s;%v;%v;%v;%v;%v\n", t.Format("2006/01/02 15:04"), co, ch, cl, cc, cv))
				m = t.Minute()
				co, ch, cl, cc = p, p, p, p
				cv = v
			}
		}

		lineNo++
	}

	writeString(fout, fmt.Sprintf("%s;%v;%v;%v;%v;%v\n", t0.Format("2006/01/02 15:04"), co, ch, cl, cc, cv))
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
