package discovery

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type InstrumentInfo struct {
	Mic            string `json:"mic"`
	MicDescription string `json:"micDescription"`
	Mep            string `json:"mep"`
	Isin           string `json:"isin"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Key            string `json:"key"`
	Type           string `json:"type"`
	IsApproved     bool   `json:"isApproved"`
	IsDiscovered   bool   `json:"isDiscovered"`
}

var (
	KnownEuronextMics = createKnownEuronextMics()
	KnownOtherMics    = createKnownOtherMics()
	instrumentInfoMap = make(map[string]*InstrumentInfo)
	knownMicToMepMap  = createKnownMicToMepMapping()
	unknownMicMap     = make(map[string]string)
	categories        = createCategories()
	bodyMap           = map[string]string{
		"start":          "0",
		"length":         "10000",
		"iDisplayStart":  "0",
		"iDisplayLength": "10000",
	}
)

type CategoryInfo struct {
	Type     string
	Uri      string
	Referer  string
	FileName string
}

func createKnownMicToMepMapping() map[string]string {
	return map[string]string{
		"XAMS": "AMS", "ALXA": "AMS", "TNLA": "AMS", "XAMC": "AMS",
		"XBRU": "BRU", "ALXB": "BRU", "ENXB": "BRU", "MLXB": "BRU", "TNLB": "BRU",
		"XPAR": "PAR", "ALXP": "PAR", "XMLI": "PAR", "XPMC": "PAR",
		"XLIS": "LIS", "ALXL": "LIS", "ENXL": "LIS",
		"XDUB": "DUB", "XMSM": "DUB", "XESM": "DUB", "XACD": "DUB",
		"XOSL": "OSL", "XOAS": "OSL", "MERK": "OSL", "VPXB": "OSL",
		"MTAA": "MIL", "MTAH": "MIL", "EXGM": "MIL", "BGEM": "MIL", "ETLX": "MIL", "ETFP": "MIL", "ATFX": "MIL", "MIVX": "MIL",
		"XLDN": "LON", "XLON": "LON", "XLIF": "LON",
		"XHFT": "OTH",
		"XETR": "OTH", "XMCE": "OTH", "XVTX": "OTH", "FRAA": "OTH", "XCSE": "OTH", "XSTO": "OTH", "XHEL": "OTH", "XMAD": "OTH", "XIST": "OTH",
	}
}

func createKnownEuronextMics() []string {
	return []string{
		"XAMS", "ALXA", "TNLA", "XAMC",
		"XBRU", "ALXB", "ENXB", "MLXB", "TNLB",
		"XPAR", "ALXP", "XMLI", "XPMC",
		"XLIS", "ALXL", "ENXL",
		"XDUB", "XMSM", "XESM", "XACD",
		"XOSL", "XOAS", "MERK", "VPXB",
		"MTAA", "MTAH", "EXGM", "BGEM", "ETLX", "ETFP", "ATFX", "MIVX",
		"XHFT",
	}
}

func createKnownOtherMics() []string {
	return []string{
		"XLDN", "XLIF", "XLON",
		"XETR", "XMCE", "XVTX", "FRAA", "XCSE", "XSTO", "XHEL", "XMAD", "XIST",
	}
}

func createStockCategory(mic string) CategoryInfo {
	return CategoryInfo{
		Type:     "stock",
		Uri:      "https://live.euronext.com/en/pd/data/stocks?mics=" + mic,
		Referer:  "https://live.euronext.com/products/equities/list",
		FileName: "stocks_" + mic,
	}
}

func createIndexCategory(mic string) CategoryInfo {
	return CategoryInfo{
		Type:     "index",
		Uri:      "https://live.euronext.com/en/pd/data/index?mics=" + mic,
		Referer:  "https://live.euronext.com/products/indices/list",
		FileName: "indices_" + mic,
	}
}

func createEtvCategory(mic string) CategoryInfo {
	return CategoryInfo{
		Type:     "etv",
		Uri:      "https://live.euronext.com/en/pd/data/etv?mics=" + mic,
		Referer:  "https://live.euronext.com/products/etfs/list",
		FileName: "etvs_" + mic,
	}
}

func createEtfCategory(mic string) CategoryInfo {
	return CategoryInfo{
		Type:     "etf",
		Uri:      "https://live.euronext.com/en/pd/data/track?mics=" + mic,
		Referer:  "https://live.euronext.com/products/etfs/list",
		FileName: "etfs_" + mic,
	}
}

func createFundCategory(mic string) CategoryInfo {
	return CategoryInfo{
		Type:     "fund",
		Uri:      "https://live.euronext.com/en/pd/data/funds?mics=" + mic,
		Referer:  "https://live.euronext.com/products/funds/list",
		FileName: "funds_" + mic,
	}
}

func createCategories() []CategoryInfo {
	categories := []CategoryInfo{}
	mics := KnownEuronextMics
	mics = append(mics, KnownOtherMics...)

	for _, mic := range mics {
		categories = append(categories, createStockCategory(mic))
	}

	for _, mic := range mics {
		categories = append(categories, createIndexCategory(mic))
	}

	for _, mic := range mics {
		categories = append(categories, createEtvCategory(mic))
	}

	for _, mic := range mics {
		categories = append(categories, createEtfCategory(mic))
	}

	for _, mic := range mics {
		categories = append(categories, createFundCategory(mic))
	}

	return categories
}

// downloadPost performs an HTTP POST request to download a file.
// Returns true if the file was downloaded successfully and meets the minimal length.
func downloadPost(
	uri string,
	filePath string,
	minimalLength int64,
	overwrite bool,
	retries int,
	timeout time.Duration,
	pauseBeforeRetry time.Duration,
	bodyKeyValueMap map[string]string,
	referer string,
	verbose bool,
	userAgent string,
) bool {
	// Prepare POST data
	postData := url.Values{}
	for key, value := range bodyKeyValueMap {
		postData.Set(key, value)
	}

	firstTry := true
	for retries > 0 {
		// If this is not the first try, wait before retrying
		if firstTry {
			firstTry = false
		} else {
			time.Sleep(pauseBeforeRetry)
		}

		// Check if file exists and meets minimal length
		if !overwrite {
			if fi, err := os.Stat(filePath); err == nil {
				if fi.Size() > minimalLength {
					log.Printf("file %s already exists, skipping\n", filePath)
					return true
				}
				log.Printf("file %s already exists but length %d is smaller than the minimal length %d, overwriting\n", filePath, fi.Size(), minimalLength)
			}
		}

		pd := postData.Encode()
		if verbose {
			log.Printf("%s POST body: %s\n", uri, pd)
		}

		req, err := http.NewRequest("POST", uri, bytes.NewBufferString(pd))
		if err != nil {
			log.Printf("failed to create request: %v\n", err)
			retries--
			continue
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("Accept-Language", "en-us,en;q=0.5")
		req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("Referer", referer)
		req.Header.Set("Accept", "application/json, text/javascript, */*")

		// Create HTTP client with timeout and proxy settings
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment, // Uses system proxy settings
		}
		client := &http.Client{Timeout: timeout, Transport: transport}

		resp, err := client.Do(req)
		if err != nil {
			if retries > 1 {
				log.Printf("file %s: download failed [%v], retrying (%d)\n", filePath, err, retries)
			} else {
				log.Printf("file %s: download failed [%v], giving up (%d)\n", filePath, err, retries)
			}
			retries--
			continue
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("file %s: download failed, status code %d is not OK, retrying (%d)\n", filePath, resp.StatusCode, retries)
			retries--
			resp.Body.Close()
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("file %s: failed to read body: [%v], retrying (%d)\n", filePath, err, retries)
			retries--
			resp.Body.Close()
			continue
		}

		if int64(len(body)) <= minimalLength {
			log.Printf("file %s: downloaded length %d is smaller than the minimal length %d, retrying\n", filePath, len(body), minimalLength)
			retries--
			resp.Body.Close()
			continue
		}

		err = os.WriteFile(filePath, body, 0644) // owner: read/write, group/others: read
		if err != nil {
			log.Printf("file %s: failed to write: [%v], retrying (%d)\n", filePath, err, retries)
			retries--
			resp.Body.Close()
			continue
		}

		resp.Body.Close()
		return true
	}
	log.Printf("file %s: failed to download after retries\n", filePath)
	return false
}

func retrieveTotalRecords(filename string) int {
	// {"iTotalRecords":null,"iTotalDisplayRecords":null,"aaData":[]}
	// {"iTotalRecords":56,"iTotalDisplayRecords":56,"aaData":[[
	const (
		prefix  = `{"iTotalRecords":`
		prefix2 = `{"iTotalRecords":0,`
		prefix3 = `{"iTotalRecords":null,`
	)

	data, err := os.ReadFile(filename)
	if err != nil {
		return 0
	}
	text := string(data)
	if strings.HasPrefix(text, prefix2) || strings.HasPrefix(text, prefix3) {
		return 0
	} else if strings.HasPrefix(text, prefix) {
		// null,"iTotalDisplayRecords"  or  56,"iTotalDisplayRecords"
		start := len(prefix)
		end := strings.Index(text[start:], ",")
		if end == -1 {
			return 0
		}
		totalRecords, err := strconv.Atoi(text[start : start+end])
		if err != nil {
			return 0
		}
		return totalRecords
	}
	return 0
}

// parseJSON parses a JSON string and updates the instrumentInfoMap.
// Returns true if parsing was successful, false otherwise.
func parseJSON(s, typ string, verbose bool) bool {
	if strings.Contains(s, "\"aaData\": []") {
		return false
	}
	if strings.Contains(s, ",null,") {
		s = strings.ReplaceAll(s, ",null,", ",\"null\",")
	}

	const splitter = "\",\""

	splitted := strings.Split(s, splitter)
	if len(splitted) < 7 {
		log.Printf("splitted array has length %d instead of 7, skipping %s\n", len(splitted), s)
		return false
	}

	stripTrailingChars := func(str string) string {
		return strings.Trim(str, "\"")
	}

	ii := &InstrumentInfo{
		Isin:           stripTrailingChars(splitted[1]),
		Symbol:         stripTrailingChars(splitted[2]),
		Name:           "",
		MicDescription: strings.ReplaceAll(stripTrailingChars(splitted[3]), `\u00e9`, "é"),
		Type:           typ,
	}
	if ii.Isin == "null" {
		ii.Isin = ""
	}
	if strings.HasSuffix(ii.MicDescription, "Pari") {
		ii.MicDescription += "s"
	}
	z := "/" + ii.Isin + "-"
	i := strings.Index(splitted[0], z)
	s0 := splitted[0]
	if i >= 0 {
		s0 = splitted[0][i+len(z):]
	}
	i = strings.Index(s0, `\`)
	if i >= 0 {
		ii.Mic = s0[:i]
	}
	if mep, ok := knownMicToMepMap[ii.Mic]; ok {
		ii.Mep = mep
	} else {
		if _, exists := unknownMicMap[ii.Mic]; !exists {
			unknownMicMap[ii.Mic] = "OTH"
		}
		ii.Mep = "OTH"
	}
	pattern := `\u0027\u003E`
	i = strings.Index(splitted[0], pattern)
	if i > 0 {
		s1 := splitted[0][i+len(pattern):]
		j := strings.Index(s1, `\u003C\/a\u003E`)
		if j > 0 {
			ii.Name = s1[:j]
			ii.Name = strings.ReplaceAll(ii.Name, `\u00E9`, "é")
			ii.Name = strings.ReplaceAll(ii.Name, `\u0026`, "&")
			ii.Name = strings.ReplaceAll(ii.Name, "&#039;", "'")
			ii.Name = strings.ReplaceAll(ii.Name, "&amp;", "&")
		}
	}
	ii.Key = strings.ToUpper(fmt.Sprintf("%s_%s_%s", ii.Mic, ii.Symbol, ii.Isin))
	if v, exists := instrumentInfoMap[ii.Key]; exists {
		if verbose {
			log.Printf("duplicate isin, skipping (2):\n")
			log.Printf("(1)(%s) mep=%s, mic=%s, symbol=%s, isin=%s\n", v.Key, v.Mep, v.Mic, v.Symbol, v.Isin)
			log.Printf("(2)(%s) mep=%s, mic=%s, symbol=%s, isin=%s\n", ii.Key, ii.Mep, ii.Mic, ii.Symbol, ii.Isin)
		}
		if !strings.EqualFold(v.Symbol, ii.Symbol) || v.Key != ii.Key || !strings.EqualFold(v.Mep, ii.Mep) || !strings.EqualFold(v.Isin, ii.Isin) {
			instrumentInfoMap[ii.Key] = ii
		}
	} else {
		instrumentInfoMap[ii.Key] = ii
	}
	return true
}

// parseFile reads the file at filename, parses its content, and calls parseJSON for each record.
// Returns the number of records parsed.
func parseFile(filename, typ string, verbose bool) int {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("error reading file %s: %v", filename, err)
		return 0
	}
	content := string(contentBytes)
	totalRecords := 0

	i := strings.Index(content, "[[")
	if i < 0 {
		if verbose {
			log.Printf("no data section found in file %s\n", filename)
		}
		return 0
	}
	content = content[i+2:]
	for {
		i = strings.Index(content, "],[")
		if i >= 0 {
			if parseJSON(content[:i], typ, verbose) {
				totalRecords++
			}
			content = content[i+3:]
		} else {
			break
		}
	}
	i = strings.Index(content, "]]")
	if i >= 0 {
		if parseJSON(content[:i], typ, verbose) {
			totalRecords++
		}
	}
	return totalRecords
}

func downloadAndParse(
	category CategoryInfo,
	folderPath string,
	downloadRetries int,
	downloadTimeout time.Duration,
	downloadPauseBeforeRetry time.Duration,
	verbose bool,
	userAgent string,
) bool {
	bodyMap["start"] = "0"
	bodyMap["iDisplayStart"] = "0"
	filename := filepath.Join(folderPath, category.FileName+".json")

	if !downloadPost(category.Uri, filename, 0, true, downloadRetries,
		downloadTimeout, downloadPauseBeforeRetry, bodyMap, category.Referer,
		verbose, userAgent) {
		return false
	}

	totalRecords := retrieveTotalRecords(filename)
	parsedRecords := parseFile(filename, category.Type, verbose)
	totalParsed := parsedRecords
	log.Printf("%s: total records = %d, parsed records = %d, total parsed = %d\n", category.FileName, totalRecords, parsedRecords, totalParsed)
	page := 0
	for totalRecords > totalParsed {
		keyNew := strconv.Itoa(totalParsed + 1)
		bodyMap["start"] = keyNew
		bodyMap["iDisplayStart"] = keyNew
		filename = filepath.Join(folderPath, fmt.Sprintf("%s.%d.json", category.FileName, page+1))
		if !downloadPost(category.Uri, filename, 0, true, downloadRetries,
			downloadTimeout, downloadPauseBeforeRetry, bodyMap, category.Referer,
			verbose, userAgent) {
			return false
		}
		parsedRecords = parseFile(filename, category.Type, verbose)
		totalParsed += parsedRecords
		log.Printf("%s: total records = %d, parsed records = %d, total parsed = %d\n", category.FileName, totalRecords, parsedRecords, totalParsed)
		if parsedRecords == 0 {
			break
		}
		page++
	}
	return true
}

// zipFiles adds files with the given extension to the zip archive recursively.
func zipFiles(archive *zip.Writer, baseDir, ext string) error {
	return filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error walking path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if len(ext) > 0 && !strings.HasSuffix(strings.ToLower(info.Name()), ext) {
			log.Printf("skipping file %s, does not match extension %s\n", path, ext)
			return nil
		}
		relPath, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}
		zipPath := filepath.ToSlash(relPath)
		w, err := archive.Create(zipPath)
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	})
}

// zipHomogenousDirectory zips all files with the given extension in a directory recursively.
func zipHomogenousDirectory(zipFile, directory, ext string, delete bool) error {
	directory = strings.TrimRight(directory, string(os.PathSeparator))
	zipWriter, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer zipWriter.Close()
	archive := zip.NewWriter(zipWriter)
	defer archive.Close()

	err = zipFiles(archive, directory, ext)
	if err != nil {
		return err
	}
	if delete {
		return os.RemoveAll(directory)
	}
	return nil
}

// Fetch downloads instrument data from Euronext and returns a map of InstrumentInfo.
// It creates a directory structure based on the current date and year, downloads JSON files for each category,
// and optionally zips the downloaded files. It also handles retries and timeouts for downloads.
func Fetch(
	downloadPath string,
	now time.Time,
	downloadRetries int,
	downloadTimeoutSec int,
	downloadPauseBeforeRetrySec int,
	zipDownloadPath bool,
	deleteDownloadPath bool,
	verbose bool,
	userAgent string,
) map[string]*InstrumentInfo {
	sep := string(os.PathSeparator)
	if downloadPath != "" && !strings.HasSuffix(downloadPath, sep) {
		downloadPath += sep
	}

	downloadPath = fmt.Sprintf("%s%d%s", downloadPath, now.Year(), sep)
	folder := now.Format("20060102")
	folderPath := filepath.Join(downloadPath, folder) + sep
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			log.Printf("cannot create directory '%s' to file: %v\n", folderPath, err)
			return instrumentInfoMap
		}
	}

	log.Printf("downloading to %s\n", folderPath)
	downloadTimeout := time.Duration(downloadTimeoutSec) * time.Second
	downloadPauseBeforeRetry := time.Duration(downloadPauseBeforeRetrySec) * time.Second
	for _, category := range categories {
		if !downloadAndParse(category, folderPath, downloadRetries, downloadTimeout,
			downloadPauseBeforeRetry, verbose, userAgent) {
			continue
		}
	}

	if zipDownloadPath {
		zipName := filepath.Join(downloadPath, folder) + "edi.zip"
		log.Printf("zipping folder %s to %s\n", folderPath, zipName)
		if err := zipHomogenousDirectory(zipName, folderPath, "", false); err != nil {
			log.Printf("error zipping folder %s to %s: %v\n", folderPath, zipName, err)
			return instrumentInfoMap
		}
		log.Printf("zipped folder %s to %s\n", folderPath, zipName)
	}

	if deleteDownloadPath {
		if err := os.RemoveAll(folderPath); err != nil {
			log.Printf("error deleting folder %s: %v\n", folderPath, err)
			return instrumentInfoMap
		}
		log.Printf("deleted folder %s\n", folderPath)
	}
	return instrumentInfoMap
}
