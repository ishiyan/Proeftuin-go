package enrichment

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"euronext/euronext"
)

// downloadTextString downloads the content from the given URL with retries and timeout.
// The referer and userAgent are set in the request headers.
func downloadTextString(
	label string,
	url string,
	retries int,
	timeout time.Duration,
	pauseBeforeRetry time.Duration,
	referer string,
	verbose bool,
	userAgent string,
) string {
	// Create HTTP client with timeout and proxy settings
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment, // Uses system proxy settings
	}
	client := http.Client{Timeout: timeout, Transport: transport}
	if verbose {
		log.Println(url)
	}

	var lastErr error
	for attempt := 1; attempt <= retries; attempt++ {
		if attempt > 1 && attempt <= retries {
			time.Sleep(pauseBeforeRetry)
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			lastErr = err
			log.Printf("[%s] attempt %d: failed to create request: %v", label, attempt, err)
			continue
		}
		req.Header.Set("Referer", referer)
		req.Header.Set("User-Agent", userAgent)

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			log.Printf("[%s] attempt %d: request failed: %v", label, attempt, err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			log.Printf("[%s] attempt %d: reading body failed: %v", label, attempt, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("unexpected status: %s", resp.Status)
			log.Printf("[%s] attempt %d: bad status: %s", label, attempt, resp.Status)
			continue
		}

		return string(body)
	}

	log.Printf("[%s] all attempts failed: %v", label, lastErr)
	return ""
}

// extract returns the substring between prefix and suffix, or "" if not found.
func extract(text, prefix, suffix string) string {
	i := strings.Index(text, prefix)
	if i >= 0 {
		s := text[i+len(prefix):]
		j := strings.Index(s, suffix)
		if j > 0 {
			s = strings.TrimSpace(s[:j])
			if s == "-" {
				s = ""
			}
			return s
		}
	}
	return ""
}

// extract2 returns the substring between prefix1, prefix2, and suffix2, or "" if not found.
func extract2(text, prefix1, prefix2, suffix2 string) string {
	i := strings.Index(text, prefix1)
	if i >= 0 {
		s := text[i+len(prefix1):]
		j := strings.Index(s, prefix2)
		if j > 0 {
			s = s[j+len(prefix2):]
			k := strings.Index(s, suffix2)
			if k > 0 {
				s = strings.TrimSpace(s[:k])
				if s == "-" {
					s = ""
				}
				return s
			}
		}
	}
	return ""
}

// EnrichStockInstrument enriches the stock element of the given instrument with additional data
// such as CFI, ICB, and trading information by downloading the relevant blocks from Euronext.
func EnrichStockInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	verbose bool,
	userAgent string,
) {
	// <instrument vendor="Euronext" mep="AMS" isin="NL0000336543" symbol="BALNE" name="BALLAST NEDAM" type="stock" mic="XAMS" file="euronext/ams/stocks/eurls/loc/BALNE.xml" description="Ballast Nedam specializes in the ... sector.">
	//   <stock cfi="ES" compartment="B" tradingMode="continuous" currency="EUR" shares="1,431,522,482">
	//     <icb icb1="2000" icb2="2300" icb3="2350" icb4="2357"/>
	//   </stock>
	// </instrument>
	if instrument == nil {
		return
	}

	if instrument.Stock == nil {
		instrument.Stock = &euronext.XmlStock{
			Cfi:         "",
			Compartment: "",
			Currency:    "EUR",
			Shares:      "",
			TradingMode: "",
			Icb: &euronext.XmlIcb{
				Icb1: "",
				Icb2: "",
				Icb3: "",
				Icb4: "",
			},
		}
	}

	if instrument.Stock.Icb == nil {
		instrument.Stock.Icb = &euronext.XmlIcb{
			Icb1: "",
			Icb2: "",
			Icb3: "",
			Icb4: "",
		}
	}

	// https://live.euronext.com/en/product/equities/NL0000336543-XAMS/market-information
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/NL0000336543-XAMS/fs_cfi_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/NL0000336543-XAMS/fs_icb_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/NL0000336543-XAMS/fs_tradinginfo_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/NL0000336543-XAMS/fs_tradinginfo_pea_block
	isin := instrument.Isin
	mic := instrument.Mic
	referer := fmt.Sprintf("https://live.euronext.com/en/product/equities/%s-%s", isin, mic)
	timeout := time.Duration(timeoutSec) * time.Second
	pauseBeforeRetry := time.Duration(pauseBeforeRetrySec) * time.Second

	// CFI block
	uri := fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/%s-%s/fs_cfi_block", isin, mic)
	str := downloadTextString("cfi_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no CFI block fetched")
	} else {
		// <tr><td>CFI:CI</td></tr>
		value := extract(str, "<tr><td>CFI:", "</td></tr>")
		if value != "" {
			instrument.Stock.Cfi = strings.ToUpper(value)
		}
	}

	// ICB block
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/%s-%s/fs_icb_block", isin, mic)
	str = downloadTextString("icb_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no ICB block fetched")
	} else {
		// <td>Industry</td>
		// <td><strong>2000, Industrials</strong></td>
		value := extract2(str, "<td>Industry</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			if i := strings.Index(value, ","); i > 0 {
				value = strings.TrimSpace(value[:i])
			}
			if value != "-" && value != "" {
				instrument.Stock.Icb.Icb1 = value
			}
		}

		// <td>SuperSector</td>
		// <td><strong>2300, Construction & Materials</strong></td>
		value = extract2(str, "<td>SuperSector</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			if i := strings.Index(value, ","); i > 0 {
				value = strings.TrimSpace(value[:i])
			}
			if value != "-" && value != "" {
				instrument.Stock.Icb.Icb2 = value
			}
		}

		// <td>Sector</td>
		// <td><strong>2350, Construction & Materials</strong></td>
		value = extract2(str, "<td>Sector</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			if i := strings.Index(value, ","); i > 0 {
				value = strings.TrimSpace(value[:i])
			}
			if value != "-" && value != "" {
				instrument.Stock.Icb.Icb3 = value
			}
		}

		// <td>Subsector</td>
		// <td><strong>2357, Heavy Construction</strong></td>
		value = extract2(str, "<td>Subsector</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			if i := strings.Index(value, ","); i > 0 {
				value = strings.TrimSpace(value[:i])
			}
			if value != "-" && value != "" {
				instrument.Stock.Icb.Icb4 = value
			}
		}
	}

	// Trading info block
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/%s-%s/fs_tradinginfo_block", isin, mic)
	str = downloadTextString("tradinginfo_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no tradinginfo block fetched")
	} else {
		// <td>Trading currency</td>
		// <td><strong>EUR</strong></td>
		value := extract2(str, "<td>Trading currency</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Stock.Currency = strings.ToUpper(value)
		}

		// <td>Trading type</td>
		// <td><strong>Continuous</strong></td>
		value = extract2(str, "<td>Trading type</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			value = strings.ToLower(value)
			if value == "continous" {
				value = "continuous"
			}
			instrument.Stock.TradingMode = value
		}

		// <td>Shares outstanding</td>
		// <td><strong>270,045,923</strong></td>
		//
		// <td>Admitted shares</td>
		// <td><strong>220,299,776</strong></td>
		value = extract2(str, "<td>Shares outstanding</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Stock.Shares = strings.ToLower(value)
		} else {
			value = extract2(str, "<td>Admitted shares</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Stock.Shares = strings.ToLower(value)
			}
		}
	}

	// Trading info pea block
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/%s-%s/fs_tradinginfo_pea_block", isin, mic)
	str = downloadTextString("tradinginfo_pea_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no tradinginfo_pea block fetched")
	} else {
		// <strong>Compartment A (Large Cap)</strong>
		value := ""
		if strings.Contains(str, "<strong>Compartment A ") {
			value = "A"
		} else if strings.Contains(str, "<strong>Compartment B ") {
			value = "B"
		} else if strings.Contains(str, "<strong>Compartment C ") {
			value = "C"
		}
		if value != "" {
			instrument.Stock.Compartment = value
		}
	}

	// Name (if missing)
	if instrument.Name == "" {
		uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getDetailedQuote/%s-%s", isin, mic)
		str = downloadTextString("detailed quote", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
		if str == "" {
			log.Println("no detailed quote fetched")
		} else {
			// <strong>BALLAST NEDAM</strong>
			value := extract(str, "<strong>", "</strong>")
			if value != "" {
				instrument.Name = value
			}
		}
	}
}

// EnrichEtfInstrument enriches the ETV element of the given instrument with additional data
// such as CFI, INAV, and underlying information by downloading the relevant blocks from Euronext.
func EnrichEtvInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	verbose bool,
	userAgent string,
) {
	// <instrument vendor="Euronext" mep="PAR" mic="XPAR" isin="GB00B15KXP72" symbol="COFFP" name="ETFS COFFEE" type="etv" file="etf/COFFP.xml" description="">
	//   <etv cfi="DTZSPR" tradingMode="continuous" allInFees="0,49%" expenseRatio="" dividendFrequency="yearly" currency="EUR" issuer="ETFS COMMODITY SECURITIES LTD" shares="944,000">
	// </instrument>
	if instrument == nil {
		return
	}

	if instrument.Etv == nil {
		instrument.Etv = &euronext.XmlEtv{
			Cfi:               "",
			Currency:          "EUR",
			AllInFees:         "",
			ExpenseRatio:      "",
			DividendFrequency: "",
			LaunchDate:        "",
			Issuer:            "",
			Shares:            "",
			TradingMode:       "",
		}
	}

	// https://live.euronext.com/en/product/etfs/XS2792094604-ETFP/market-information
	// https://live.euronext.com/en/product/etfs/DE000A28M8D0-XPAR/market-information
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/XS2792094604-ETFP/fs_cfi_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/XS2792094604-ETFP/fs_generalinfo_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/XS2792094604-ETFP/fs_tradinginfo_etfs_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/XS2792094604-ETFP/fs_feessegmentation_block
	// https://live.euronext.com/en/ajax/getDetailedQuote/DE000A3G3ZL3-XPAR
	isin := instrument.Isin
	mic := instrument.Mic
	referer := fmt.Sprintf("https://live.euronext.com/en/product/etvs/%s-%s", isin, mic)
	timeout := time.Duration(timeoutSec) * time.Second
	pauseBeforeRetry := time.Duration(pauseBeforeRetrySec) * time.Second

	// --- CFI block ---
	uri := fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_cfi_block", isin, mic)
	str := downloadTextString("cfi_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no CFI block fetched")
	} else {
		// <tr><td>CFI:CI</td></tr>
		value := extract(str, "<tr><td>CFI:", "</td></tr>")
		if value != "" {
			instrument.Etv.Cfi = strings.ToUpper(value)
		}
	}

	// --- General Info block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_generalinfo_block", isin, mic)
	str = downloadTextString("generalinfo_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no generalinfo block fetched")
	} else {
		// <td>ETF Legal Name</td>
		// <td><strong>AMUNDI ETF GOVT BOND EURO BROAD INVESTMENT GRADE 1-3 UCITS ETF</strong></td>
		value := extract2(str, "<td>ETF Legal Name</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			upperValue := strings.ToUpper(value)
			instrument.Description = &upperValue
		}

		// Issuer: try several possible fields
		// <td>Issuer Name</td>
		// <td><strong>Amundi Asset Management</strong></td>
		//
		// <td>Nom de l'émetteur</td>
		// <td><strong>HSBC GLOBAL FUNDS ICAV</strong></td>
		//
		// <td>Fund Manager</td>
		// <td><strong>Amundi Investment Solutions</strong></td>
		value = extract2(str, "<td>Nom de l'émetteur</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.Issuer = strings.ToUpper(value)
		} else {
			value = extract2(str, "<td>Issuer Name</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Etv.Issuer = strings.ToUpper(value)
			} else {
				value = extract2(str, "<td>Fund Manager</td>", "<td><strong>", "</strong></td>")
				if value != "" {
					instrument.Etv.Issuer = strings.ToUpper(value)
				}
			}
		}

		// <td>Launch Date</td>
		// <td><strong>26/06/2009</strong></td>
		value = extract2(str, "<td>Launch Date</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.LaunchDate = strings.ToUpper(value)
		}

		// <td>Dividend frequency</td>
		// <td><strong>Annually</strong></td>
		value = extract2(str, "<td>Dividend frequency</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.DividendFrequency = strings.ToLower(value)
		}
	}

	// --- Trading Info ETFs block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_tradinginfo_etfs_block", isin, mic)
	str = downloadTextString("tradinginfo_etfs_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no tradinginfo_etfs block fetched")
	} else {
		// <td>Trading currency</td>
		// <td><strong>EUR</strong></td>
		value := extract2(str, "<td>Trading currency</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.Currency = strings.ToUpper(value)
		}

		// <td>Trading type</td>
		// <td><strong>Continuous</strong></td>
		value = extract2(str, "<td>Trading type</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			value = strings.ToLower(value)
			if value == "continous" {
				value = "continuous"
			}
			instrument.Etv.TradingMode = value
		}

		// <td>Exposition type</td>
		// <td><strong>Synthetic</strong></td>
		value = extract2(str, "<td>Exposition type</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.ExpositionType = strings.ToLower(value)
		}

		// <td>Shares outstanding</td>
		// <td><strong>270,045,923</strong></td>
		//
		// <td>Admitted shares</td>
		// <td><strong>220,299,776</strong></td>
		value = extract2(str, "<td>Shares outstanding</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.Shares = strings.ToLower(value)
		} else {
			value = extract2(str, "<td>Admitted shares</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Etv.Shares = strings.ToLower(value)
			}
		}
	}

	// --- Fees Segmentation block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_feessegmentation_block", isin, mic)
	str = downloadTextString("feessegmentation_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no feessegmentation block fetched")
	} else {
		// <td>All In Fees</td>
		// <td><strong>0,49%</strong></td>
		value := extract2(str, "<td>All In Fees</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.AllInFees = value
		}

		// <td>Expense Ratio</td>
		// <td><strong>0,49%</strong></td>
		value = extract2(str, "<td>Expense Ratio</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.ExpenseRatio = strings.ToLower(value)
		} else {
			// <td>TER</td>
			// <td><strong>0.14%</strong></td>
			value = extract2(str, "<td>TER</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Etv.ExpenseRatio = strings.ToUpper(value)
			}
		}
	}

	// --- Name (if missing) ---
	if instrument.Name == "" {
		uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getDetailedQuote/%s-%s", isin, mic)
		str = downloadTextString("detailed quote", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
		if str == "" {
			log.Println("no detailed quote fetched")
		} else {
			// <strong>Bitwise MSCI Digital Assets Select 20 ETP</strong>
			value := extract(str, "<strong>", "</strong>")
			if value != "" {
				instrument.Name = value
			}
		}
	}
}

// EnrichEtfInstrument enriches the ETF element of the given instrument with additional data
// such as CFI, INAV, and underlying information by downloading the relevant blocks from Euronext.
func EnrichEtfInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	verbose bool,
	userAgent string,
) {
	// <instrument vendor="Euronext" mep="PAR" mic="XPAR" isin="FR0010754135" symbol="C13" name="AMUNDI ETF EMTS1-3" type="etf" file="etf/C13.xml" description="Amundi ETF Govt Bond EuroMTS Broad 1-3">
	//   <etf cfi="EUOM" ter="0.14" tradingMode="continuous" launchDate="20100316" currency="EUR" issuer="AMUNDI" fraction="1" dividendFrequency="Annually" indexFamily="EuroMTS" expositionType="synthetic">
	//     <inav vendor="Euronext" mep="PAR" mic="XPAR" isin="QS0011161377" symbol="INC13" name="AMUNDI C13 INAV"/>
	//     <underlying vendor="Euronext" mep="PAR" mic="XPAR" isin="QS0011052618" symbol="EMTSAR" name="EuroMTS Eurozone Government Broad 1-3"/>
	//   </etf>
	// </instrument>
	if instrument == nil {
		return
	}

	if instrument.Etf == nil {
		instrument.Etf = &euronext.XmlEtf{
			Cfi:               "",
			Currency:          "EUR",
			Ter:               "",
			Mer:               "",
			DividendFrequency: "",
			ExpositionType:    "",
			Fraction:          "",
			LaunchDate:        "",
			Issuer:            "",
			IndexFamily:       "",
			TradingMode:       "",
			Shares:            "",
			Inav: euronext.XmlInav{
				Currency: "",
				Isin:     "",
				Mep:      "",
				Mic:      "",
				Name:     "",
				Symbol:   "",
				Vendor:   "",
				/*Target: []XmlTarget{
					{
						Isin:   "",
						Mep:    "",
						Mic:    "",
						Name:   "",
						Symbol: "",
						Vendor: "",
					},
				},*/
			},
			Underlying: euronext.XmlUnderlying{
				Isin:   "",
				Mep:    "",
				Mic:    "",
				Name:   "",
				Symbol: "",
				Vendor: "",
			},
		}
	}

	// https://live.euronext.com/en/product/etfs/IE0000KA1ZX3-ETFP/market-information
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/IE0000KA1ZX3-ETFP/fs_cfi_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/IE0000KA1ZX3-ETFP/fs_generalinfo_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/IE0000KA1ZX3-ETFP/fs_tradinginfo_etfs_block
	// https://live.euronext.com/en/ajax/getDetailedQuote/IE0000KA1ZX3-ETFP
	//
	// https://live.euronext.com/en/product/etfs/FR0010754135-XPAR/market-information
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/FR0010754135-XPAR/fs_cfi_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/FR0010754135-XPAR/fs_generalinfo_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/FR0010754135-XPAR/fs_tradinginfo_etfs_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/FR0010754135-XPAR/fs_feessegmentation_block

	isin := instrument.Isin
	mic := instrument.Mic
	referer := fmt.Sprintf("https://live.euronext.com/en/product/equities/%s-%s", isin, mic)
	timeout := time.Duration(timeoutSec) * time.Second
	pauseBeforeRetry := time.Duration(pauseBeforeRetrySec) * time.Second

	// --- CFI block ---
	uri := fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_cfi_block", isin, mic)
	str := downloadTextString("cfi_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no CFI block fetched")
	} else {
		// <tr><td>CFI:CI</td></tr>
		value := extract(str, "<tr><td>CFI:", "</td></tr>")
		if value != "" {
			instrument.Etf.Cfi = strings.ToUpper(value)
		}
	}

	// --- General Info block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_generalinfo_block", isin, mic)
	str = downloadTextString("generalinfo_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no generalinfo block fetched")
	} else {
		// <td>ETF Legal Name</td>
		// <td><strong>AMUNDI ETF GOVT BOND EURO BROAD INVESTMENT GRADE 1-3 UCITS ETF</strong></td>
		value := extract2(str, "<td>ETF Legal Name</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			upperValue := strings.ToUpper(value)
			instrument.Description = &upperValue
		}

		// Issuer: try several possible fields
		//
		// <td>Issuer Name</td>
		// <td><strong>Amundi Asset Management</strong></td>
		//
		// <td>Nom de l'émetteur</td>
		// <td><strong>HSBC GLOBAL FUNDS ICAV</strong></td>
		//
		// <td>Fund Manager</td>
		// <td><strong>Amundi Investment Solutions</strong></td>
		value = extract2(str, "<td>Nom de l'émetteur</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Issuer = strings.ToUpper(value)
		} else {
			value = extract2(str, "<td>Issuer Name</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Etf.Issuer = strings.ToUpper(value)
			} else {
				value = extract2(str, "<td>Fund Manager</td>", "<td><strong>", "</strong></td>")
				if value != "" {
					instrument.Etf.Issuer = strings.ToUpper(value)
				}
			}
		}

		// <td>Launch Date</td>
		// <td><strong>26/06/2009</strong></td>
		value = extract2(str, "<td>Launch Date</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.LaunchDate = strings.ToUpper(value)
		}

		// <td>Dividend frequency</td>
		// <td><strong>Annually</strong></td>
		value = extract2(str, "<td>Dividend frequency</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.DividendFrequency = strings.ToLower(value)
		}
	}

	// --- Trading Info ETFs block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_tradinginfo_etfs_block", isin, mic)
	str = downloadTextString("tradinginfo_etfs_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no tradinginfo_etfs block fetched")
	} else {
		// <td>Trading currency</td>
		// <td><strong>EUR</strong></td>
		value := extract2(str, "<td>Trading currency</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Currency = strings.ToUpper(value)
		}

		// <td>Trading type</td>
		// <td><strong>Continuous</strong></td>
		value = extract2(str, "<td>Trading type</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			value = strings.ToLower(value)
			if value == "continous" {
				value = "continuous"
			}
			instrument.Etf.TradingMode = value
		}

		// <td>Exposition type</td>
		// <td><strong>Synthetic</strong></td>
		value = extract2(str, "<td>Exposition type</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.ExpositionType = strings.ToLower(value)
		}

		// <td>Shares outstanding</td>
		// <td><strong>270,045,923</strong></td>
		//
		// <td>Admitted shares</td>
		// <td><strong>220,299,776</strong></td>
		value = extract2(str, "<td>Shares outstanding</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Shares = strings.ToLower(value)
		} else {
			value = extract2(str, "<td>Admitted shares</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Etf.Shares = strings.ToLower(value)
			}
		}

		// <td>Ticker INAV (Euronext)</td>
		// <td><strong>INC13</strong></td>
		value = extract2(str, "<td>Ticker INAV (Euronext)</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Inav.Symbol = value
		}

		// <td>INAV Name</td>
		// <td><strong>AMUNDI C13 INAV</strong></td>
		value = extract2(str, "<td>INAV Name</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Inav.Name = value
		}

		// <td>INAV ISIN code</td>
		// <td><strong>QS0011161377</strong></td>
		value = extract2(str, "<td>INAV ISIN code</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Inav.Isin = value
		}

		// <td>Underlying index</td>
		// <td><strong>FTSE Eurozone Gvt Br IG 1-3Y</strong></td>
		value = extract2(str, "<td>Underlying index</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Underlying.Name = value
		}

		// <td>Index</td>
		// <td><strong>EMIGA5</strong></td>
		value = extract2(str, "<td>Index</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Underlying.Symbol = value
		}
	}

	// --- Fees Segmentation block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/TRACK/%s-%s/fs_feessegmentation_block", isin, mic)
	str = downloadTextString("feessegmentation_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no feessegmentation block fetched")
	} else {
		// <td>TER</td>
		//  <td><strong>0.14%</strong></td>
		value := extract2(str, "<td>TER</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etf.Ter = strings.ToUpper(value)
		}
	}

	// --- Name (if missing) ---
	if instrument.Name == "" {
		uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getDetailedQuote/%s-%s", isin, mic)
		str = downloadTextString("detailed quote", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
		if str == "" {
			log.Println("no detailed quote fetched")
		} else {
			value := extract(str, "<strong>", "</strong>")
			if value != "" {
				instrument.Name = value
			}
		}
	}
}

// EnrichFundInstrument enriches the Fund element of the given instrument with additional data
// such as CFI, INAV, and underlying information by downloading the relevant blocks from Euronext.
func EnrichFundInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	verbose bool,
	userAgent string,
) {
	// <instrument vendor="Euronext" mep="AMS" mic="XAMS" isin="NL0006259996" symbol="AWAF" name="ACH WERELD AANDFD3" type="fund" file="fund/AWAF.xml" description="">
	//   <fund cfi="EUOISB" tradingmode="fixing" currency="EUR" issuer="ACHMEA BELEGGINGSFONDSEN" shares="860,248">
	// </instrument>
	if instrument == nil {
		return
	}

	if instrument.Fund == nil {
		instrument.Fund = &euronext.XmlFund{
			Cfi:         "",
			Currency:    "EUR",
			Issuer:      "",
			Shares:      "",
			TradingMode: "",
		}
	}

	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/LU2264552998-ATFX/fs_cfi_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/LU2264552998-ATFX/fs_issuerinfo_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/LU2264552998-ATFX/fs_tradinginfo_funds_block
	// https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/LU2264552998-ATFX/fs_info_block
	isin := instrument.Isin
	mic := instrument.Mic
	referer := fmt.Sprintf("https://live.euronext.com/en/product/funds/%s-%s", isin, mic)
	timeout := time.Duration(timeoutSec) * time.Second
	pauseBeforeRetry := time.Duration(pauseBeforeRetrySec) * time.Second

	// --- CFI block ---
	uri := fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/%s-%s/fs_cfi_block", isin, mic)
	str := downloadTextString("cfi_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no CFI block fetched")
	} else {
		// <tr><td>CFI:CI</td></tr>
		value := extract(str, "<tr><td>CFI:", "</td></tr>")
		if value != "" {
			instrument.Fund.Cfi = strings.ToUpper(value)
		}
	}

	// --- General Info block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/%s-%s/fs_issuerinfo_block", isin, mic)
	str = downloadTextString("issuerinfo_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no issuerinfo block fetched")
	} else {
		// >Issuer name : </span> <span class="issuerName-column-right"><strong>VARENNE UCITS</strong>
		value := extract(str, ">Issuer name : </span> <span class=\"issuerName-column-right\"><strong>", "</strong>")
		if value != "" {
			instrument.Fund.Issuer = value
		}

		// Issuer: try several possible fields
		// <td>Issuer Name</td>
		// <td><strong>Amundi Asset Management</strong></td>
		//
		// <td>Nom de l'émetteur</td>
		// <td><strong>HSBC GLOBAL FUNDS ICAV</strong></td>
		//
		// <td>Fund Manager</td>
		// <td><strong>Amundi Investment Solutions</strong></td>
		value = extract2(str, "<td>Nom de l'émetteur</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Etv.Issuer = strings.ToUpper(value)
		} else {
			value = extract2(str, "<td>Issuer Name</td>", "<td><strong>", "</strong></td>")
			if value != "" {
				instrument.Etv.Issuer = strings.ToUpper(value)
			} else {
				value = extract2(str, "<td>Fund Manager</td>", "<td><strong>", "</strong></td>")
				if value != "" {
					instrument.Etv.Issuer = strings.ToUpper(value)
				}
			}
		}
	}

	// --- Trading Info ETFs block ---
	uri = fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/FUNDS/%s-%s/fs_tradinginfo_funds_block", isin, mic)
	str = downloadTextString("tradinginfo_funds_block", uri, retries, timeout, pauseBeforeRetry, referer, verbose, userAgent)
	if str == "" {
		log.Println("no tradinginfo_funds block fetched")
	} else {
		// <td>Trading currency</td>
		// <td><strong>EUR</strong></td>
		value := extract2(str, "<td>Trading currency</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Fund.Currency = strings.ToUpper(value)
		}

		// <td>Shares outstanding</td>
		// <td><strong>2,422,386</strong></td>
		value = extract2(str, "<td>Shares outstanding</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			instrument.Fund.Shares = value
		}

		// <td>Trading type</td>
		// <td><strong>Continuous</strong></td>
		value = extract2(str, "<td>Trading type</td>", "<td><strong>", "</strong></td>")
		if value != "" {
			value = strings.ToLower(value)
			if value == "continous" {
				value = "continuous"
			}
			instrument.Fund.TradingMode = value
		}
	}
}

// EnrichInavInstrument enriches the Inav element of the given instrument with additional data
// such as CFI, INAV, and underlying information by downloading the relevant blocks from Euronext.
func EnrichInavInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	userAgent string,
) {
	// <instrument vendor="Euronext" mep="PAR" isin="QS0011161385" symbol="INC33" name="AMUNDI C33 INAV" type="inav" file="etf/INC33.xml" description="iNav Amundi ETF Govt Bond EuroMTS Broad 3-5">
	//   <inav currency="EUR">
	//     <target vendor="Euronext" mep="PAR" mic="XPAR" isin="FR0010754168" symbol="C33" name="AMUNDI ETF GOV 3-5"/>
	//   </inav>
	// </instrument>
	if instrument == nil {
		return
	}

	if instrument.Inav == nil {
		instrument.Inav = &euronext.XmlInav{
			Currency: "",
			Isin:     "",
			Mep:      "",
			Mic:      "",
			Name:     "",
			Symbol:   "",
			Vendor:   "",
			Target: []euronext.XmlTarget{
				{
					Isin:   "",
					Mep:    "",
					Mic:    "",
					Name:   "",
					Symbol: "",
					Vendor: "",
				},
			},
		}
	}
}

// EnrichIndexInstrument enriches the Index element of the given instrument with additional data
// such as ICB, INAV, and underlying information by downloading the relevant blocks from Euronext.
func EnrichIndexInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	userAgent string,
) {
	// https://live.euronext.com/en/product/indices/FR0014002B31-XPAR/market-information
	if instrument == nil {
		return
	}

	if instrument.Index == nil {
		instrument.Index = &euronext.XmlIndex{
			BaseDate:  "",
			BaseLevel: "",
			Currency:  "EUR",
			Family:    "",
			Kind:      "",
			Weighting: "",
			CapFactor: "",
			CalcFreq:  "",
			Icb: &euronext.XmlIcb{
				Icb1: "",
				Icb2: "",
				Icb3: "",
				Icb4: "",
			},
		}
	}
}

func EnrichInstrument(
	instrument *euronext.XmlInstrument,
	retries int,
	timeoutSec int,
	pauseBeforeRetrySec int,
	verbose bool,
	userAgent string,
) {
	if instrument == nil {
		return
	}

	switch instrument.Type {
	case "stock":
		EnrichStockInstrument(instrument, retries, timeoutSec, pauseBeforeRetrySec, verbose, userAgent)
	case "etv":
		EnrichEtvInstrument(instrument, retries, timeoutSec, pauseBeforeRetrySec, verbose, userAgent)
	case "etf":
		EnrichEtfInstrument(instrument, retries, timeoutSec, pauseBeforeRetrySec, verbose, userAgent)
	case "fund":
		EnrichFundInstrument(instrument, retries, timeoutSec, pauseBeforeRetrySec, verbose, userAgent)
	case "inav":
		EnrichInavInstrument(instrument, retries, timeoutSec, pauseBeforeRetrySec, userAgent)
	case "index":
		EnrichIndexInstrument(instrument, retries, timeoutSec, pauseBeforeRetrySec, userAgent)
	}
}
