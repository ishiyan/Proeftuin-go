package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const header = `import { TimeGranularity } from 'projects/mb/src/public-api';
import { Series } from '../../series.interface';

export const tradeSeriesZzzAperiodic: Series = {
	mnemonic: 'zzz',
	description: 'zzz.',
	timeStart: new Date(1999, 0, 4),
	timeEnd: new Date(2021, 11, 31),
	timeGranularity: TimeGranularity.Aperiodic,
	data: [`

const footer = `  ],
};`

func main() {
	//tformatPtr := flag.String("tformat", "2006/02/01 15:04:05.9999999", "csv time format in go style")
	//tzonePtr := flag.String("tzone", "", "time zone of the data")
	//volumePtr := flag.Int("volume", 0, "volume value if not present")
	headerPtr := flag.Bool("header", true, "the very first CSV line is a header")

	flag.Parse()

	var fin *os.File
	var fout *os.File

	if filename := flag.Arg(0); filename == "" {
		fmt.Println("expecting CSV file name as the positional argument")
		os.Exit(1)
	} else {
		if !strings.HasSuffix(filename, ".csv") {
			fmt.Println("expecting CSV file name to end with '.csv': ", filename)
			os.Exit(1)
		}

		fin, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: ", err)
			os.Exit(1)
		}
		defer fin.Close()

		fout, err := os.Create(filename + ".ts")
		if err != nil {
			panic(fmt.Sprintf("error creating file: %s", err))
		}
		defer fout.Close()
	}

	writeString(fout, header)

	csvReader := csv.NewReader(fin)
	csvReader.Comment = '#'
	csvReader.Comma = ';'

	lineNo := 0

	for {
		rec, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(fmt.Sprintf("error reading line %d: %s", lineNo, err))
		}

		if lineNo == 0 && *headerPtr {
			continue
		}

		if len(rec) < 2 {
			fmt.Printf("line %d: expected at least 2 parts, got %d", lineNo, len(rec))
			os.Exit(1)
		}

		fmt.Fprintf(fout, "%+v\n", rec)
	}

	writeString(fout, footer)
}

func writeString(f *os.File, s string) {
	if _, err := f.WriteString(s); err != nil {
		panic(fmt.Sprintf("error writing file: %s", err))
	}
}
