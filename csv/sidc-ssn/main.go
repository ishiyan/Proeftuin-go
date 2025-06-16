package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	RtList = make(map[time.Time]float64)
	RsList = make(map[time.Time]float64)
	RnList = make(map[time.Time]float64)
)

func fetch(url string) {
	log.Printf("Downloading URL %s", url)
	client := &http.Client{
		Timeout: 180 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Request creation failed: %v", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091106 Shiretoko/3.5.5")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Download failed: %v, uri=%s", err, url)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}
		line = strings.TrimRight(line, "\r\n")
		if len(line) < 37 {
			log.Printf("illegal line [%s], length < 37", line)
			continue
		}
		s := line[0:10]
		if s[5] == ' ' {
			sb := []byte(s)
			sb[5] = '0'
			s = string(sb)
		}
		dt, err := time.Parse("2006 01 02", s)
		if err != nil {
			log.Printf("cannot parse date-time [%s] in line [%s], skipping the line", s, line)
			continue
		}
		s = strings.TrimSpace(line[20:25])
		rt, err := strconv.Atoi(s)
		if err != nil {
			log.Printf("cannot parse total sunspot number [%s] in line [%s], skipping the line", s, line)
			continue
		}
		if len(line) > 37 {
			s = strings.TrimSpace(line[24:29])
			rn, err := strconv.Atoi(s)
			if err != nil {
				log.Printf("cannot parse north hemisphere sunspot number [%s] in line [%s], skipping the line", s, line)
				continue
			}
			s = strings.TrimSpace(line[24:29])
			rs, err := strconv.Atoi(s)
			if err != nil {
				log.Printf("cannot parse south hemisphere sunspot number [%s] in line [%s], skipping the line", s, line)
				continue
			}
			if rn < 0 {
				RnList[dt] = math.NaN()
			} else {
				RnList[dt] = float64(rn)
			}
			if rs < 0 {
				RsList[dt] = math.NaN()
			} else {
				RsList[dt] = float64(rs)
			}
		}
		if rt < 0 {
			RtList[dt] = math.NaN()
		} else {
			RtList[dt] = float64(rt)
		}
	}
	log.Printf("Download complete")
}

func main() {
	log.Println("=======================================================================================")
	log.Printf("Started: %s", time.Now().Format(time.RFC3339))

	// Fetch data
	fetch("http://www.sidc.be/silso/DATA/EISN/EISN_current.txt")
	fetch("http://www.sidc.be/silso/DATA/SN_d_hem_V2.0.txt")

	// Prepare output directory
	//h5File := "repository/sidc.h5"
	//dir := filepath.Dir(h5File)
	//if _, err := os.Stat(dir); os.IsNotExist(err) {
	//	os.MkdirAll(dir, 0755)
	//}

	// Write RtList, RnList, RsList to HDF5 (pseudo-code)
	// Replace with actual HDF5 writing logic as needed
	log.Printf("Updating Rt: %s", time.Now().Format(time.RFC3339))
	for dt, v := range RtList {
		log.Println(dt.Format("2006 01 02"), v)
	}
	log.Printf("Updating Rn: %s", time.Now().Format(time.RFC3339))
	for dt, v := range RnList {
		log.Println(dt.Format("2006 01 02"), v)
	}
	log.Printf("Updating Rs: %s", time.Now().Format(time.RFC3339))
	for dt, v := range RsList {
		log.Println(dt.Format("2006 01 02"), v)
	}

	log.Printf("Finished: %s", time.Now().Format(time.RFC3339))
}
