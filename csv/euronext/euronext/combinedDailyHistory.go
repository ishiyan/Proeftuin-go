package euronext

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CombinedDailyHistory struct {
	Date                   time.Time `json:"date"`
	Open                   float64   `json:"open"`
	High                   float64   `json:"high"`
	Low                    float64   `json:"low"`
	Last                   float64   `json:"last"`
	Close                  float64   `json:"close"`
	NumberOfShares         float64   `json:"number of shares"`
	NumberOfTrades         float64   `json:"number of trades"`
	Turnover               float64   `json:"turnover"`
	Vwap                   float64   `json:"vwap"`
	OpenAdjusted           float64   `json:"open adjusted"`
	HighAdjusted           float64   `json:"high adjusted"`
	LowAdjusted            float64   `json:"low adjusted"`
	LastAdjusted           float64   `json:"last adjusted"`
	CloseAdjusted          float64   `json:"close adjusted"`
	NumberOfSharesAdjusted float64   `json:"number of shares adjusted"`
	NumberOfTradesAdjusted float64   `json:"number of trades adjusted"`
	TurnoverAdjusted       float64   `json:"turnover adjusted"`
	VwapAdjusted           float64   `json:"vwap adjusted"`
	AdjustmentFactor       float64   `json:"adjustment factor"`
	HasMarking             bool      `json:"has marking"`
	HasMarkingAdjusted     bool      `json:"has marking adjusted"`
}

const CombinedDailyHistoryDateFormat = "2006-01-02"

func CombinedDailyHistoryHeaders() []string {
	return []string{
		"date", "open", "high", "low", "last", "close", "number of shares", "number of trades",
		"turnover", "vwap", "open adjusted", "high adjusted", "low adjusted", "last adjusted",
		"close adjusted", "number of shares adjusted", "number of trades adjusted", "turnover adjusted",
		"vwap adjusted", "adjustment factor", "has marking", "has marking adjusted",
	}
}

func WriteCombinedDailyHistoryCsv(fileName string, history []CombinedDailyHistory) (string, error) {
	gz := strings.HasSuffix(fileName, ".gz")

	file, err := os.Create(fileName)
	if err != nil {
		es := fmt.Sprintf("cannot create csv file %s: ", fileName)
		return es, fmt.Errorf("%s%w", es, err)
	}
	defer file.Close()

	var w *csv.Writer
	if gz {
		gzipWriter := gzip.NewWriter(file)
		defer gzipWriter.Close()
		w = csv.NewWriter(gzipWriter)
	} else {
		w = csv.NewWriter(file)
	}
	defer w.Flush()

	if err := w.Write(CombinedDailyHistoryHeaders()); err != nil {
		es := fmt.Sprintf("cannot write header to csv file %s: ", fileName)
		return es, fmt.Errorf("%s%w", es, err)
	}

	for _, rec := range history {
		row := []string{
			rec.Date.Format(CombinedDailyHistoryDateFormat),
			strconv.FormatFloat(rec.Open, 'f', -1, 64),
			strconv.FormatFloat(rec.High, 'f', -1, 64),
			strconv.FormatFloat(rec.Low, 'f', -1, 64),
			strconv.FormatFloat(rec.Last, 'f', -1, 64),
			strconv.FormatFloat(rec.Close, 'f', -1, 64),
			strconv.FormatFloat(rec.NumberOfShares, 'f', -1, 64),
			strconv.FormatFloat(rec.NumberOfTrades, 'f', -1, 64),
			strconv.FormatFloat(rec.Turnover, 'f', -1, 64),
			strconv.FormatFloat(rec.Vwap, 'f', -1, 64),
			strconv.FormatFloat(rec.OpenAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.HighAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.LowAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.LastAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.CloseAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.NumberOfSharesAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.NumberOfTradesAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.TurnoverAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.VwapAdjusted, 'f', -1, 64),
			strconv.FormatFloat(rec.AdjustmentFactor, 'f', -1, 64),
			strconv.FormatBool(rec.HasMarking),
			strconv.FormatBool(rec.HasMarkingAdjusted),
		}

		if err := w.Write(row); err != nil {
			es := fmt.Sprintf("cannot write row to csv file %s: ", fileName)
			return es, fmt.Errorf("%s%w", es, err)
		}
	}

	return "", nil
}

func ReadCombinedDailyHistoryCsv(fileName string) ([]CombinedDailyHistory, string, error) {
	gz := strings.HasSuffix(fileName, ".gz")
	history := []CombinedDailyHistory{}

	file, err := os.Open(fileName)
	if err != nil {
		es := fmt.Sprintf("cannot open csv file %s: ", fileName)
		return history, es, fmt.Errorf("%s%w", es, err)
	}
	defer file.Close()

	var r *csv.Reader
	if gz {
		gzipReader, err := gzip.NewReader(file)
		if err != nil {
			es := fmt.Sprintf("cannot create gzip reader for file %s: ", fileName)
			return history, es, fmt.Errorf("%s%w", es, err)
		}
		defer gzipReader.Close()
		r = csv.NewReader(gzipReader)
	} else {
		r = csv.NewReader(file)
	}

	if _, err := r.Read(); err != nil {
		es := fmt.Sprintf("cannot read header from csv file %s: ", fileName)
		return history, es, fmt.Errorf("%s%w", es, err)
	}

	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			es := fmt.Sprintf("cannot read row from csv file %s: ", fileName)
			return history, es, fmt.Errorf("%s%w", es, err)
		}

		date, err := time.Parse(CombinedDailyHistoryDateFormat, record[0])
		if err != nil {
			es := fmt.Sprintf("cannot parse date from csv file %s: ", fileName)
			return history, es, fmt.Errorf("%s%w", es, err)
		}

		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		last, _ := strconv.ParseFloat(record[4], 64)
		close, _ := strconv.ParseFloat(record[5], 64)
		numberOfShares, _ := strconv.ParseFloat(record[6], 64)
		numberOfTrades, _ := strconv.ParseFloat(record[7], 64)
		turnover, _ := strconv.ParseFloat(record[8], 64)
		vwap, _ := strconv.ParseFloat(record[9], 64)
		openAdjusted, _ := strconv.ParseFloat(record[10], 64)
		highAdjusted, _ := strconv.ParseFloat(record[11], 64)
		lowAdjusted, _ := strconv.ParseFloat(record[12], 64)
		lastAdjusted, _ := strconv.ParseFloat(record[13], 64)
		closeAdjusted, _ := strconv.ParseFloat(record[14], 64)
		numberOfSharesAdjusted, _ := strconv.ParseFloat(record[15], 64)
		numberOfTradesAdjusted, _ := strconv.ParseFloat(record[16], 64)
		turnoverAdjusted, _ := strconv.ParseFloat(record[17], 64)
		vwapAdjusted, _ := strconv.ParseFloat(record[18], 64)
		adjustmentFactor, _ := strconv.ParseFloat(record[19], 64)
		hasMarking, _ := strconv.ParseBool(record[20])
		HasMarkingAdjusted, _ := strconv.ParseBool(record[21])

		history = append(history, CombinedDailyHistory{
			Date:                   date,
			Open:                   open,
			High:                   high,
			Low:                    low,
			Last:                   last,
			Close:                  close,
			NumberOfShares:         numberOfShares,
			NumberOfTrades:         numberOfTrades,
			Turnover:               turnover,
			Vwap:                   vwap,
			OpenAdjusted:           openAdjusted,
			HighAdjusted:           highAdjusted,
			LowAdjusted:            lowAdjusted,
			LastAdjusted:           lastAdjusted,
			CloseAdjusted:          closeAdjusted,
			NumberOfSharesAdjusted: numberOfSharesAdjusted,
			NumberOfTradesAdjusted: numberOfTradesAdjusted,
			TurnoverAdjusted:       turnoverAdjusted,
			VwapAdjusted:           vwapAdjusted,
			AdjustmentFactor:       adjustmentFactor,
			HasMarking:             hasMarking,
			HasMarkingAdjusted:     HasMarkingAdjusted,
		})
	}

	return history, "", nil
}

func SortCombinedDailyHistory(history []CombinedDailyHistory) []CombinedDailyHistory {
	sortHistory := make([]CombinedDailyHistory, len(history))
	copy(sortHistory, history)
	for i := 0; i < len(sortHistory); i++ {
		for j := i + 1; j < len(sortHistory); j++ {
			if sortHistory[i].Date.After(sortHistory[j].Date) {
				sortHistory[i], sortHistory[j] = sortHistory[j], sortHistory[i]
			}
		}
	}

	return sortHistory
}

func MergeCombinedDailyHistory(histOld, histNew []CombinedDailyHistory) ([]CombinedDailyHistory, []string) {
	messages := []string{}
	histMapOld := make(map[time.Time]CombinedDailyHistory)
	histMapNew := make(map[time.Time]CombinedDailyHistory)

	for _, entry := range histOld {
		histMapOld[entry.Date] = entry
	}
	for _, entry := range histNew {
		histMapNew[entry.Date] = entry
	}

	// Create a set of all dates
	dateSet := make(map[time.Time]struct{})
	for date := range histMapOld {
		dateSet[date] = struct{}{}
	}
	for date := range histMapNew {
		dateSet[date] = struct{}{}
	}

	// Collect the dates into a slice
	var dates []time.Time
	for date := range dateSet {
		dates = append(dates, date)
	}

	// Sort the slice in descending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].After(dates[j])
	})

	// Iterate through the sorted slice in descending order
	var mergedHistory []CombinedDailyHistory
	phase := 0
	multiplier := 1.
	factorNew := 0.
	factorOld := 0.
	for _, date := range dates {
		entryOld, existsOld := histMapOld[date]
		entryNew, existsNew := histMapNew[date]

		// We only update newer downloaded files here,
		// so we should have 3 phases when looping through the dates.
		// Weare looping in the descending order.
		// 1. New enntry exists, old entry does not exist.
		//    Take the new entry and add it to the merged history.
		// 2. Both entries exist. If we enteredthis phase,
		//    we should never see phase 1 anymore.
		//    Take the new entry and add it to the merged history.
		//    Verify that not adjusted values are the same.
		//    Calculate multiplier for adjustment factors.
		// 3. Old entry exists, new entry does not exist.
		//    If we entered this phase, we should never see phase 1 or 2 anymore.
		//    Take the old entry and add it to the merged history.
		//    Multiply by the calculated multiplier if it is not 1.

		if existsNew && !existsOld {
			if phase == 0 {
				phase = 1
			} else if phase != 1 {
				messages = append(messages,
					fmt.Sprintf("Date %s (pase 1) expecting phase 1, got phase %d",
						date.Format("2006-01-02"), phase))
			}
			mergedHistory = append(mergedHistory, entryNew)
		} else if existsNew && existsOld {
			if phase == 1 {
				phase = 2
			} else if phase != 2 {
				messages = append(messages,
					fmt.Sprintf("Date %s (pase 2) expecting phase 2, got phase %d",
						date.Format("2006-01-02"), phase))
			}
			notEqual := []string{}
			if entryOld.Open != entryNew.Open {
				notEqual = append(notEqual, "open")
			}
			if entryOld.High != entryNew.High {
				notEqual = append(notEqual, "high")
			}
			if entryOld.Low != entryNew.Low {
				notEqual = append(notEqual, "low")
			}
			if entryOld.Last != entryNew.Last {
				notEqual = append(notEqual, "last")
			}
			if entryOld.Close != entryNew.Close {
				notEqual = append(notEqual, "close")
			}
			if entryOld.NumberOfShares != entryNew.NumberOfShares {
				notEqual = append(notEqual, "number of shares")
			}
			if entryOld.NumberOfTrades != entryNew.NumberOfTrades {
				notEqual = append(notEqual, "number of trades")
			}
			if entryOld.Turnover != entryNew.Turnover {
				notEqual = append(notEqual, "turnover")
			}
			if entryOld.Vwap != entryNew.Vwap {
				notEqual = append(notEqual, "vwap")
			}
			if entryOld.HasMarking != entryNew.HasMarking {
				notEqual = append(notEqual, "has marking")
			}
			if len(notEqual) > 0 {
				messages = append(messages,
					fmt.Sprintf("Date %s (pase 2) different values for %s",
						date.Format("2006-01-02"), strings.Join(notEqual, ", ")))
			}
			factorNew = entryNew.AdjustmentFactor
			factorOld = entryOld.AdjustmentFactor
			if factorOld != 0 {
				multiplier = factorNew / factorOld
			}
			mergedHistory = append(mergedHistory, entryNew)
		} else { // if !existsNew && existsOld
			if phase == 2 {
				phase = 3
			} else if phase != 3 {
				messages = append(messages,
					fmt.Sprintf("Date %s (pase 3) expecting phase 3, got phase %d",
						date.Format("2006-01-02"), phase))
			}

			if multiplier != 1.0 {
				entryOld.AdjustmentFactor *= multiplier
				entryOld.OpenAdjusted *= multiplier
				entryOld.HighAdjusted *= multiplier
				entryOld.LowAdjusted *= multiplier
				entryOld.LastAdjusted *= multiplier
				entryOld.CloseAdjusted *= multiplier
				entryOld.NumberOfSharesAdjusted /= multiplier
				entryOld.VwapAdjusted *= multiplier
			}
			mergedHistory = append(mergedHistory, entryOld)
		}
	}

	return SortCombinedDailyHistory(mergedHistory), messages
}
