package main

import (
	"archive/zip"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Config holds the JSON configuration.
type Config struct {
	InputBaseFolder  string `json:"input_base_folder"`
	RepositoryFolder string `json:"repository_folder"`
}

// Record represents a single OHLCV row.
type Record struct {
	DateTime string // "YYYY-MM-DD" or "YYYY-MM-DD hh:mm:ss"
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
}

// timeframeInfo describes how to map a zip file to output suffixes and datetime formats.
type timeframeInfo struct {
	ZipName      string   // e.g. "d_us_txt.zip"
	PathPrefix   string   // e.g. "data/daily/us/"
	Country      string   // e.g. "us"
	FileSuffixes []string // e.g. [".us"] — parts between ticker and ".txt"; empty string means just ".txt"
	Suffix       string   // e.g. "_1d.csv.gz"
	Daily        bool     // true => date-only format
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <country> <date>\nPlease specify a country (e.g. us, uk, jp, hk, hu) and a date (e.g. 2025-12-27)\n", os.Args[0])
		os.Exit(1)
	}
	country := strings.ToLower(os.Args[1])
	date := os.Args[2]

	cfgPath := "stooqmerge.json"
	if len(os.Args) > 3 {
		cfgPath = os.Args[3]
	}

	cfg, err := loadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	timeframes := []timeframeInfo{
		{ZipName: "d_us_txt.zip", PathPrefix: "data/daily/us/", Country: "us", FileSuffixes: []string{".us"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "h_us_txt.zip", PathPrefix: "data/hourly/us/", Country: "us", FileSuffixes: []string{".us"}, Suffix: "_1h.csv.gz", Daily: false},
		{ZipName: "5_us_txt.zip", PathPrefix: "data/5 min/us/", Country: "us", FileSuffixes: []string{".us"}, Suffix: "_5m.csv.gz", Daily: false},
		{ZipName: "d_uk_txt.zip", PathPrefix: "data/daily/uk/", Country: "uk", FileSuffixes: []string{".uk"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "h_uk_txt.zip", PathPrefix: "data/hourly/uk/", Country: "uk", FileSuffixes: []string{".uk"}, Suffix: "_1h.csv.gz", Daily: false},
		{ZipName: "5_uk_txt.zip", PathPrefix: "data/5 min/uk/", Country: "uk", FileSuffixes: []string{".uk"}, Suffix: "_5m.csv.gz", Daily: false},
		{ZipName: "d_jp_txt.zip", PathPrefix: "data/daily/jp/", Country: "jp", FileSuffixes: []string{".jp"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "d_hk_txt.zip", PathPrefix: "data/daily/hk/", Country: "hk", FileSuffixes: []string{".hk"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "h_hk_txt.zip", PathPrefix: "data/hourly/hk/", Country: "hk", FileSuffixes: []string{".hk"}, Suffix: "_1h.csv.gz", Daily: false},
		{ZipName: "5_hk_txt.zip", PathPrefix: "data/5 min/hk/", Country: "hk", FileSuffixes: []string{".hk"}, Suffix: "_5m.csv.gz", Daily: false},
		{ZipName: "d_hu_txt.zip", PathPrefix: "data/daily/hu/", Country: "hu", FileSuffixes: []string{".hu"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "h_hu_txt.zip", PathPrefix: "data/hourly/hu/", Country: "hu", FileSuffixes: []string{".hu"}, Suffix: "_1h.csv.gz", Daily: false},
		{ZipName: "5_hu_txt.zip", PathPrefix: "data/5 min/hu/", Country: "hu", FileSuffixes: []string{".hu"}, Suffix: "_5m.csv.gz", Daily: false},
		{ZipName: "d_world_txt.zip", PathPrefix: "data/daily/world/", Country: "world", FileSuffixes: []string{"", ".b", ".v"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "h_world_txt.zip", PathPrefix: "data/hourly/world/", Country: "world", FileSuffixes: []string{"", ".b", ".v"}, Suffix: "_1h.csv.gz", Daily: false},
		{ZipName: "5_world_txt.zip", PathPrefix: "data/5 min/world/", Country: "world", FileSuffixes: []string{"", ".b", ".v"}, Suffix: "_5m.csv.gz", Daily: false},
		{ZipName: "d_pl_txt.zip", PathPrefix: "data/daily/pl/", Country: "pl", FileSuffixes: []string{"", ".hu", ".n", ".pl"}, Suffix: "_1d.csv.gz", Daily: true},
		{ZipName: "h_pl_txt.zip", PathPrefix: "data/hourly/pl/", Country: "pl", FileSuffixes: []string{"", ".hu", ".n", ".pl"}, Suffix: "_1h.csv.gz", Daily: false},
		{ZipName: "5_pl_txt.zip", PathPrefix: "data/5 min/pl/", Country: "pl", FileSuffixes: []string{"", ".hu", ".n", ".pl"}, Suffix: "_5m.csv.gz", Daily: false},
		{ZipName: "d_macro_txt.zip", PathPrefix: "data/daily/macro/", Country: "macro", FileSuffixes: []string{".m"}, Suffix: "_1d.csv.gz", Daily: true},
	}

	found := false
	for _, tf := range timeframes {
		if tf.Country != country {
			continue
		}
		found = true
		zipPath := filepath.Join(cfg.InputBaseFolder, date, tf.ZipName)
		if _, err := os.Stat(zipPath); os.IsNotExist(err) {
			log.Printf("Zip file not found, skipping: %s", zipPath)
			continue
		}
		log.Printf("Start processing %s", zipPath)
		if err := processZip(zipPath, tf, cfg.RepositoryFolder); err != nil {
			log.Printf("Error processing %s: %v", zipPath, err)
		}
		log.Printf("Finished processing %s", zipPath)
	}
	if !found {
		fmt.Fprintf(os.Stderr, "Unknown country: %s\nSupported countries: us, uk, jp, hk, hu, world\n", country)
		os.Exit(1)
	}
}

// loadConfig reads and parses the JSON configuration file.
func loadConfig(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}
	if cfg.InputBaseFolder == "" || cfg.RepositoryFolder == "" {
		return cfg, fmt.Errorf("input_base_folder and repository_folder must be set in config")
	}
	return cfg, nil
}

// processZip opens a zip archive and processes each *.txt file inside it.
func processZip(zipPath string, tf timeframeInfo, repoFolder string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("open zip %s: %w", zipPath, err)
	}
	defer r.Close()

	for _, f := range r.File {
		name := f.Name
		// Only process .txt files
		if f.FileInfo().IsDir() || !strings.HasSuffix(strings.ToLower(name), ".txt") {
			continue
		}
		if err := processZipEntry(f, tf, repoFolder); err != nil {
			log.Printf("Error processing %s in %s: %v", name, zipPath, err)
		}
	}
	return nil
}

// processZipEntry reads a single *.txt from the zip, converts it, and merges with existing data.
func processZipEntry(f *zip.File, tf timeframeInfo, repoFolder string) error {
	// Determine the output path.
	// Input: "data/daily/us/nasdaq stocks/1/goog.us.txt"
	// We strip the path prefix to get: "nasdaq stocks/1/goog.us.txt"
	// Then strip numeric leaf subfolder and .us.txt to build:
	//   "{repoFolder}/us/nasdaq stocks/goog.1d.csv"
	name := f.Name
	// Normalise forward slashes
	name = strings.ReplaceAll(name, "\\", "/")

	// Strip the timeframe path prefix (case-insensitive match)
	rel := ""
	if idx := strings.Index(strings.ToLower(name), strings.ToLower(tf.PathPrefix)); idx >= 0 {
		rel = name[idx+len(tf.PathPrefix):]
	} else {
		return fmt.Errorf("unexpected path structure: %s", name)
	}

	// rel is e.g. "nasdaq stocks/1/goog.us.txt" or "nasdaq etfs/goog.us.txt"
	// Split into directory and filename
	dir := filepath.Dir(rel)
	base := filepath.Base(rel)

	// Remove file suffix (e.g. ".us.txt", ".b.txt", or just ".txt")
	ticker := base
	lower := strings.ToLower(ticker)
	stripped := false
	for _, fs := range tf.FileSuffixes {
		suffix := strings.ToLower(fs) + ".txt"
		if strings.HasSuffix(lower, suffix) {
			ticker = ticker[:len(ticker)-len(suffix)]
			stripped = true
			break
		}
	}
	if !stripped && strings.HasSuffix(lower, ".txt") {
		ticker = ticker[:len(ticker)-len(".txt")]
	}

	// If the immediate parent is a numeric subfolder (1, 2, 3, ...), strip it
	dirParts := strings.Split(filepath.ToSlash(dir), "/")
	if len(dirParts) > 0 {
		last := dirParts[len(dirParts)-1]
		if isNumeric(last) {
			dirParts = dirParts[:len(dirParts)-1]
		}
	}
	leafDir := strings.Join(dirParts, string(filepath.Separator))

	outFileName := strings.ToLower(ticker) + tf.Suffix
	outPath := filepath.Join(repoFolder, tf.Country, leafDir, outFileName)

	// Read input records from zip entry
	inputRecords, err := readZipCSV(f, tf.Daily)
	if err != nil {
		return fmt.Errorf("read csv %s: %w", f.Name, err)
	}
	if len(inputRecords) == 0 {
		return nil
	}

	// Sort input records by datetime
	sort.Slice(inputRecords, func(i, j int) bool {
		return inputRecords[i].DateTime < inputRecords[j].DateTime
	})

	// Read existing merged file if it exists
	var merged []Record
	if _, err := os.Stat(outPath); err == nil {
		merged, err = readMergedFile(outPath)
		if err != nil {
			return fmt.Errorf("read merged file %s: %w", outPath, err)
		}
	}

	if len(merged) == 0 {
		// No existing data, just write input
		merged = inputRecords
	} else {
		merged = mergeRecords(merged, inputRecords, outPath)
	}

	// Write the result
	if err := writeMergedFile(outPath, merged); err != nil {
		return fmt.Errorf("write merged file %s: %w", outPath, err)
	}
	log.Printf("Updated: %s (%d records)", outPath, len(merged))
	return nil
}

// readZipCSV reads records from a zip file entry and converts them.
func readZipCSV(f *zip.File, daily bool) ([]Record, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var records []Record

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Skip header
		if strings.HasPrefix(line, "<") {
			continue
		}
		// Parse: TICKER,PER,DATE,TIME,OPEN,HIGH,LOW,CLOSE,VOL,OPENINT
		fields := strings.Split(line, ",")
		if len(fields) < 10 {
			continue
		}

		dateStr := fields[2] // YYYYMMDD
		timeStr := fields[3] // HHMMSS
		open := fields[4]
		high := fields[5]
		low := fields[6]
		close_ := fields[7]
		vol := fields[8]

		var dt string
		if daily {
			// YYYY-MM-DD
			if len(dateStr) != 8 {
				continue
			}
			dt = dateStr[:4] + "-" + dateStr[4:6] + "-" + dateStr[6:8]
		} else {
			// YYYY-MM-DD hh:mm:ss
			if len(dateStr) != 8 || len(timeStr) < 6 {
				continue
			}
			dt = dateStr[:4] + "-" + dateStr[4:6] + "-" + dateStr[6:8] + " " +
				timeStr[:2] + ":" + timeStr[2:4] + ":" + timeStr[4:6]
		}

		records = append(records, Record{
			DateTime: dt,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close_,
			Volume:   vol,
		})
	}
	return records, nil
}

// readMergedFile reads an existing gzip-compressed merged CSV file.
func readMergedFile(path string) ([]Record, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("gzip reader: %w", err)
	}
	defer gr.Close()

	data, err := io.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var records []Record
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Format: DATETIME;OPEN;HIGH;LOW;CLOSE;VOLUME
		fields := strings.Split(line, ";")
		if len(fields) < 6 {
			continue
		}
		records = append(records, Record{
			DateTime: fields[0],
			Open:     fields[1],
			High:     fields[2],
			Low:      fields[3],
			Close:    fields[4],
			Volume:   fields[5],
		})
	}
	return records, nil
}

// mergeRecords merges input records into existing merged records.
// - Prepends records with datetime before the first merged record.
// - Appends records with datetime after the last merged record.
// - For overlapping datetimes with differences:
//   - Same volume: replace merged with input (correction), log ", corrected"
//   - Different volume: replace merged with input (split adjustment), log ", updated"
//     then adjust all merged records before the earliest split date.
func mergeRecords(merged, input []Record, filePath string) []Record {
	if len(merged) == 0 {
		return input
	}

	firstDT := merged[0].DateTime
	lastDT := merged[len(merged)-1].DateTime

	// Build a map of existing datetimes for quick lookup
	existingMap := make(map[string]int, len(merged))
	for i, r := range merged {
		existingMap[r.DateTime] = i
	}

	var prepend []Record
	var append_ []Record

	// Track the earliest datetime where a volume-different (split) diff occurs
	earliestSplitDT := ""
	var adjustmentRatio float64

	for _, r := range input {
		if r.DateTime < firstDT {
			prepend = append(prepend, r)
		} else if r.DateTime > lastDT {
			append_ = append(append_, r)
		} else {
			// Overlapping range — check for differences
			if idx, ok := existingMap[r.DateTime]; ok {
				existing := merged[idx]
				if existing.Open != r.Open || existing.High != r.High ||
					existing.Low != r.Low || existing.Close != r.Close ||
					existing.Volume != r.Volume {

					if existing.Volume == r.Volume {
						// Same volume => price correction
						log.Printf("DIFF %s at %s: merged=[%s;%s;%s;%s;%s] input=[%s;%s;%s;%s;%s], corrected",
							filePath, r.DateTime,
							existing.Open, existing.High, existing.Low, existing.Close, existing.Volume,
							r.Open, r.High, r.Low, r.Close, r.Volume)
						merged[idx] = r
					} else if existing.Open == r.Open || existing.High == r.High || existing.Low == r.Low || existing.Close == r.Close {
						// Same price (at least one of open, high, low or close ), volume differs => price and volume correction
						log.Printf("DIFF %s at %s: merged=[%s;%s;%s;%s;%s] input=[%s;%s;%s;%s;%s], corrected",
							filePath, r.DateTime,
							existing.Open, existing.High, existing.Low, existing.Close, existing.Volume,
							r.Open, r.High, r.Low, r.Close, r.Volume)
						merged[idx] = r
					} else {
						// Different volume and all prices => stock split / adjustment
						log.Printf("DIFF %s at %s: merged=[%s;%s;%s;%s;%s] input=[%s;%s;%s;%s;%s], updated",
							filePath, r.DateTime,
							existing.Open, existing.High, existing.Low, existing.Close, existing.Volume,
							r.Open, r.High, r.Low, r.Close, r.Volume)

						// Track earliest split datetime and compute ratio there
						if earliestSplitDT == "" || r.DateTime < earliestSplitDT {
							earliestSplitDT = r.DateTime
							volInput, err1 := strconv.ParseFloat(r.Volume, 64)
							volMerged, err2 := strconv.ParseFloat(existing.Volume, 64)
							if err1 == nil && err2 == nil && volMerged != 0 {
								adjustmentRatio = volInput / volMerged
							}
						}

						merged[idx] = r
					}
				}
			}
		}
	}

	// If a split was detected, adjust all merged records before the earliest split date
	if earliestSplitDT != "" && adjustmentRatio != 0 {
		adjusted := 0
		for i := range merged {
			if merged[i].DateTime >= earliestSplitDT {
				break
			}
			open, e1 := strconv.ParseFloat(merged[i].Open, 64)
			high, e2 := strconv.ParseFloat(merged[i].High, 64)
			low, e3 := strconv.ParseFloat(merged[i].Low, 64)
			cl, e4 := strconv.ParseFloat(merged[i].Close, 64)
			vol, e5 := strconv.ParseFloat(merged[i].Volume, 64)
			if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil {
				continue
			}
			merged[i].Open = formatFloat(open / adjustmentRatio)
			merged[i].High = formatFloat(high / adjustmentRatio)
			merged[i].Low = formatFloat(low / adjustmentRatio)
			merged[i].Close = formatFloat(cl / adjustmentRatio)
			merged[i].Volume = formatFloat(vol * adjustmentRatio)
			adjusted++
		}
		if adjusted > 0 {
			log.Printf("DIFF %s: adjusted %d records before %s with ratio %s",
				filePath, adjusted, earliestSplitDT, formatFloat(adjustmentRatio))
		}
	}

	// Sort prepend and append sections
	sort.Slice(prepend, func(i, j int) bool {
		return prepend[i].DateTime < prepend[j].DateTime
	})
	sort.Slice(append_, func(i, j int) bool {
		return append_[i].DateTime < append_[j].DateTime
	})

	// Build final result: prepend + merged + append
	result := make([]Record, 0, len(prepend)+len(merged)+len(append_))
	result = append(result, prepend...)
	result = append(result, merged...)
	result = append(result, append_...)
	return result
}

// formatFloat formats a float64 with minimal trailing zeros.
func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// writeMergedFile writes records to disk as a gzip-compressed CSV file.
func writeMergedFile(path string, records []Record) error {
	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create directory %s: %w", dir, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %s: %w", path, err)
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	defer gw.Close()

	for _, r := range records {
		if _, err := fmt.Fprintf(gw, "%s;%s;%s;%s;%s;%s\n",
			r.DateTime, r.Open, r.High, r.Low, r.Close, r.Volume); err != nil {
			return fmt.Errorf("write record: %w", err)
		}
	}

	return gw.Close()
}

// isNumeric returns true if the string consists only of digits.
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
