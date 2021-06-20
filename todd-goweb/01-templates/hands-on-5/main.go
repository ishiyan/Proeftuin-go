package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   float64
}

func main() {
	// http.HandleFunc("/", foo)
	// http.ListenAndServe(":8080", nil)
	foo(nil, nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := prs("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute( /*res*/ os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)
		close, _ := strconv.ParseFloat(row[4], 64)
		adjClose, _ := strconv.ParseFloat(row[6], 64)
		volume, _ := strconv.ParseFloat(row[5], 64)

		records = append(records, Record{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			AdjClose: adjClose,
			Volume:   volume,
		})
	}

	return records

}
