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
// https://www.sidc.be/SILSO/home
// https://www.sidc.be/SILSO/datafiles
/*
DOI:	https://doi.org/10.24414/qnza-ac80
Title:	International Sunspot Number V2.0
Published:	2015
Publisher:	Royal Observatory of Belgium (ROB)
Data Availability:	https://www.sidc.be/SILSO/datafiles
Creators:	Clette Frédéric, Lefèvre Laure
Contributors:	SILSO team
*/
// https://www.astro.oma.be/doi/ROB-SIDC-SILSO_SunspotNumberV2.html
// this link contains the data description
/*
https://www.sidc.be/SILSO/infosndhem
Daily hemispheric sunspot number (1992-now):
Daily total and hemispheric sunspot numbers derived by the formula: R= Ns + 10 * Ng, with Ns the number of spots and Ng the number of groups counted either over the entire solar disk (total), the North hemisphere or South hemisphere (based on the sunspot group heliographic latitude).
Files: SN_d_hem_V2.0.txt (column format), SN_d_hem_V2.0.csv (CSV format)
Contents: year, month, day, decimal year, SNvalue(tot), SNvalue(N), SNvalue(S), SNerror(tot), SNerror(N), SNerror(S), Nb observations,indicator
https://www.sidc.be/SILSO/DATA/SN_d_hem_V2.0.csv
https://www.sidc.be/SILSO/DATA/SN_d_hem_V2.0.txt
1992 01 01 1992.001  186   0 186   14.3   1.0  14.3   19  -1  -1  
1992 01 02 1992.004  190  18 172    8.2   2.6   7.8   21  -1  -1  
1992 01 03 1992.007  234  26 208   18.3   6.1  17.2   21  -1  -1  
. . .
2025 06 26 2025.484  105  47  58   13.1   7.3   6.1   38  31  31 *
2025 06 27 2025.486  118  49  69   14.0   7.9   8.1   39  33  33 *
2025 06 28 2025.489  137  72  65   16.1   7.5  15.2   35  31  31 *
2025 06 29 2025.492  152  68  84   13.5   8.1   9.2   43  35  35 *
2025 06 30 2025.495  153  66  87   16.9  12.4  12.5   38  32  32 *
*/
/*
https://www.sidc.be/SILSO/infosndtot
Daily total sunspot number (1818-now):
Daily total sunspot number derived by the formula: R= Ns + 10 * Ng, with Ns the number of spots and Ng the number of groups counted over the entire solar disk.
Files:SN_d_tot_V2.0.txt (column format), SN_d_tot_V2.0.csv (CSV format)
Contents: year, month, day, decimal year, SNvalue , SNerror, Nb observations,indicator
https://www.sidc.be/SILSO/DATA/SN_d_tot_V2.0.csv
https://www.sidc.be/SILSO/DATA/SN_d_tot_V2.0.txt
1818  1 01 1818.001   -1  -1.0    0  
1818  1 02 1818.004   -1  -1.0    0  
1818  1 03 1818.007   -1  -1.0    0  
1818  1 04 1818.010   -1  -1.0    0  
. . .
2025 06 27 2025.486  118  14.0   39 *
2025 06 28 2025.489  137  16.1   35 *
2025 06 29 2025.492  152  13.5   43 *
2025 06 30 2025.495  153  16.9   38 *
*/
/*
https://www.sidc.be/SILSO/eisninfo
Estimated international sunspot number (EISN) :
The estimated international sunspot number (EISN) is a daily value obtained by a simple average over available sunspot counts from prompt stations in the SILSO network. The raw values from each station are scaled using their mean annual k personal coefficient over the last elapsed year.
File : EISN_current.txt
Content : year, month, day, decimal year, estimated SN , estimated error, nb stat used, nb stat available,indicator
https://www.sidc.be/SILSO/DATA/EISN/EISN_current.txt
2025 07 01 2025.497 155  12.5  35  43
2025 07 02 2025.500 147  10.7  37  43
2025 07 03 2025.503 123  11.7  38  42
2025 07 04 2025.505 108   9.6  37  44
2025 07 05 2025.508  77   7.6  27  33
2025 07 06 2025.511  87   9.0  22  25
2025 07 07 2025.514  94  10.1  24  28
2025 07 08 2025.516  88  11.3  24  26
2025 07 09 2025.519  88  10.9  33  36
2025 07 10 2025.522  80  10.0  37  41
2025 07 11 2025.525  97  12.9  33  40
2025 07 12 2025.527 115  12.8  28  30
2025 07 13 2025.530 129  12.4  24  27
2025 07 14 2025.533 142  13.0  18  23
*/
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
