package euronext

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	XmlInstrumentStockType = "stock"
	XmlInstrumentIndexType = "index"
	XmlInstrumentEtvType   = "etv"
	XmlInstrumentEtfType   = "etf"
	XmlInstrumentInavType  = "inav"
	XmlInstrumentFundType  = "fund"
)

type XmlIcb struct {
	Icb1 string `xml:"icb1,attr" json:"icb1"`
	Icb2 string `xml:"icb2,attr" json:"icb2"`
	Icb3 string `xml:"icb3,attr" json:"icb3"`
	Icb4 string `xml:"icb4,attr" json:"icb4"`
}

type XmlStock struct {
	Cfi         string  `xml:"cfi,attr" json:"cfi"`
	Compartment string  `xml:"compartment,attr" json:"compartment"`
	Currency    string  `xml:"currency,attr" json:"currency"`
	Shares      string  `xml:"shares,attr" json:"shares"`
	TradingMode string  `xml:"tradingMode,attr" json:"tradingMode"`
	Icb         *XmlIcb `xml:"icb" json:"icb"`
}

type XmlIndex struct {
	BaseCap         *string `xml:"baseCap,attr" json:"baseCap"`
	BaseCapCurrency *string `xml:"baseCapCurrency,attr" json:"baseCapCurrency"`
	BaseDate        string  `xml:"baseDate,attr" json:"baseDate"`
	BaseLevel       string  `xml:"baseLevel,attr" json:"baseLevel"`
	CalcFreq        string  `xml:"calcFreq,attr" json:"calcFreq"`
	CapFactor       string  `xml:"capFactor,attr" json:"capFactor"`
	Currency        string  `xml:"currency,attr" json:"currency"`
	Family          string  `xml:"family,attr" json:"family"`
	Kind            string  `xml:"kind,attr" json:"kind"`
	Weighting       string  `xml:"weighting,attr" json:"weighting"`
	Icb             *XmlIcb `xml:"icb" json:"icb"`
}

type XmlTarget struct {
	Isin   string `xml:"isin,attr" json:"isin"`
	Mep    string `xml:"mep,attr" json:"mep"`
	Mic    string `xml:"mic,attr" json:"mic"`
	Name   string `xml:"name,attr" json:"name"`
	Symbol string `xml:"symbol,attr" json:"symbol"`
	Vendor string `xml:"vendor,attr" json:"vendor"`
}

type XmlInav struct {
	Currency string      `xml:"currency,attr" json:"currency"`
	Isin     string      `xml:"isin,attr" json:"isin"`
	Mep      string      `xml:"mep,attr" json:"mep"`
	Mic      string      `xml:"mic,attr" json:"mic"`
	Name     string      `xml:"name,attr" json:"name"`
	Symbol   string      `xml:"symbol,attr" json:"symbol"`
	Vendor   string      `xml:"vendor,attr" json:"vendor"`
	Target   []XmlTarget `xml:"target" json:"target"`
}

type XmlFund struct {
	Cfi         string `xml:"cfi,attr" json:"cfi"`
	Currency    string `xml:"currency,attr" json:"currency"`
	Issuer      string `xml:"issuer,attr" json:"issuer"`
	Shares      string `xml:"shares,attr" json:"shares"`
	TradingMode string `xml:"tradingMode,attr" json:"tradingMode"`
}

type XmlEtv struct {
	AllInFees         string `xml:"allInFees,attr" json:"allInFees"`
	Cfi               string `xml:"cfi,attr" json:"cfi"`
	Currency          string `xml:"currency,attr" json:"currency"`
	DividendFrequency string `xml:"dividendFrequency,attr" json:"dividendFrequency"`
	ExpenseRatio      string `xml:"expenseRatio,attr" json:"expenseRatio"`
	LaunchDate        string `xml:"launchDate,attr" json:"launchDate"`
	Issuer            string `xml:"issuer,attr" json:"issuer"`
	Shares            string `xml:"shares,attr" json:"shares"`
	TradingMode       string `xml:"tradingMode,attr" json:"tradingMode"`
}

type XmlUnderlying struct {
	Isin   string `xml:"isin,attr" json:"isin"`
	Mep    string `xml:"mep,attr" json:"mep"`
	Mic    string `xml:"mic,attr" json:"mic"`
	Name   string `xml:"name,attr" json:"name"`
	Symbol string `xml:"symbol,attr" json:"symbol"`
	Vendor string `xml:"vendor,attr" json:"vendor"`
}

type XmlEtf struct {
	Cfi               string        `xml:"cfi,attr" json:"cfi"`
	Currency          string        `xml:"currency,attr" json:"currency"`
	DividendFrequency string        `xml:"dividendFrequency,attr" json:"dividendFrequency"`
	ExpositionType    string        `xml:"expositionType,attr" json:"expositionType"`
	Fraction          string        `xml:"fraction,attr" json:"fraction"`
	IndexFamily       string        `xml:"indexFamily,attr" json:"indexFamily"`
	Issuer            string        `xml:"issuer,attr" json:"issuer"`
	LaunchDate        string        `xml:"launchDate,attr" json:"launchDate"`
	Mer               string        `xml:"mer,attr" json:"mer"`
	Ter               string        `xml:"ter,attr" json:"ter"`
	TradingMode       string        `xml:"tradingMode,attr" json:"tradingMode"`
	Shares            string        `xml:"shares,attr" json:"shares"`
	Inav              XmlInav       `xml:"inav" json:"inav"`
	Underlying        XmlUnderlying `xml:"underlying" json:"underlying"`
}

type XmlInstrument struct {
	Cfi           *string   `xml:"cfi,attr" json:"cfi,omitempty"`
	Description   *string   `xml:"description,attr" json:"description,omitempty"`
	File          string    `xml:"file,attr" json:"file"`
	FoundInSearch string    `xml:"foundInSearch,attr" json:"foundInSearch"`
	Isin          string    `xml:"isin,attr" json:"isin"`
	Mep           string    `xml:"mep,attr" json:"mep"`
	Mic           string    `xml:"mic,attr" json:"mic"`
	Name          string    `xml:"name,attr" json:"name"`
	Notes         *string   `xml:"notes,attr" json:"notes,omitempty"`
	Symbol        string    `xml:"symbol,attr" json:"symbol"`
	Tradingmode   *string   `xml:"tradingmode,attr" json:"tradingmode,omitempty"`
	Type          string    `xml:"type,attr" json:"type"`
	Vendor        *string   `xml:"vendor,attr" json:"vendor,omitempty"`
	Etf           *XmlEtf   `xml:"etf" json:"etf,omitempty"`
	Etv           *XmlEtv   `xml:"etv" json:"etv,omitempty"`
	Fund          *XmlFund  `xml:"fund" json:"fund,omitempty"`
	Inav          *XmlInav  `xml:"inav" json:"inav,omitempty"`
	Index         *XmlIndex `xml:"index" json:"index,omitempty"`
	Stock         *XmlStock `xml:"stock" json:"stock,omitempty"`
}

type XmlInstruments struct {
	Instrument []XmlInstrument `xml:"instrument" json:"instrument"`
}

func h5(mic, isin, mnemonic, typ string) string {
	// xams/index/ASCX.h5:/XAMS_ASCX_NL0000249142
	return fmt.Sprintf("%s/%s/%s.h5:/%s_%s_%s",
		strings.ToLower(mic),
		strings.ToLower(typ),
		strings.ToUpper(mnemonic),
		strings.ToUpper(mic),
		strings.ToUpper(mnemonic),
		strings.ToUpper(isin))
}

func strPtr(s string) *string {
	return &s
}

func NewXmlStock(mic, isin, mnemonic, mep string) *XmlInstrument {
	return &XmlInstrument{
		Type:        XmlInstrumentStockType,
		Mic:         strings.ToUpper(mic),
		Isin:        strings.ToUpper(isin),
		Symbol:      strings.ToUpper(mnemonic),
		Mep:         strings.ToUpper(mep),
		File:        h5(mic, isin, mnemonic, XmlInstrumentStockType),
		Description: strPtr(""),
		Vendor:      strPtr("Euronext"),
		Stock: &XmlStock{
			Cfi:         "",
			Compartment: "",
			Currency:    "EUR",
			Shares:      "",
			TradingMode: "",
			Icb: &XmlIcb{
				Icb1: "",
				Icb2: "",
				Icb3: "",
				Icb4: "",
			},
		},
	}
}

func NewXmlIndex(mic, isin, mnemonic, mep string) *XmlInstrument {
	return &XmlInstrument{
		Type:        XmlInstrumentIndexType,
		Mic:         strings.ToUpper(mic),
		Isin:        strings.ToUpper(isin),
		Symbol:      strings.ToUpper(mnemonic),
		Mep:         strings.ToUpper(mep),
		File:        h5(mic, isin, mnemonic, XmlInstrumentIndexType),
		Description: strPtr(""),
		Vendor:      strPtr("Euronext"),
		Index: &XmlIndex{
			BaseDate:  "",
			BaseLevel: "",
			Currency:  "EUR",
			Family:    "",
			Kind:      "",
			Weighting: "",
			CapFactor: "",
			CalcFreq:  "",
			// BaseCap: strPtr(""),
			// BaseCapCurrency: strPtr(""),
			Icb: &XmlIcb{
				Icb1: "",
				Icb2: "",
				Icb3: "",
				Icb4: "",
			},
		},
	}
}

func NewXmlEtv(mic, isin, mnemonic, mep string) *XmlInstrument {
	return &XmlInstrument{
		Type:        XmlInstrumentEtvType,
		Mic:         strings.ToUpper(mic),
		Isin:        strings.ToUpper(isin),
		Symbol:      strings.ToUpper(mnemonic),
		Mep:         strings.ToUpper(mep),
		File:        h5(mic, isin, mnemonic, XmlInstrumentEtvType),
		Description: strPtr(""),
		Vendor:      strPtr("Euronext"),
		Etv: &XmlEtv{
			Cfi:               "",
			Currency:          "EUR",
			AllInFees:         "",
			ExpenseRatio:      "",
			DividendFrequency: "",
			LaunchDate:        "",
			Issuer:            "",
			Shares:            "",
			TradingMode:       "",
		},
	}
}

func NewXmlEtf(mic, isin, mnemonic, mep string) *XmlInstrument {
	return &XmlInstrument{
		Type:        XmlInstrumentEtfType,
		Mic:         strings.ToUpper(mic),
		Isin:        strings.ToUpper(isin),
		Symbol:      strings.ToUpper(mnemonic),
		Mep:         strings.ToUpper(mep),
		File:        h5(mic, isin, mnemonic, XmlInstrumentEtfType),
		Description: strPtr(""),
		Vendor:      strPtr("Euronext"),
		Etf: &XmlEtf{
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
			Inav: XmlInav{
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
			Underlying: XmlUnderlying{
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

func NewXmlInav(mic, isin, mnemonic, mep string) *XmlInstrument {
	return &XmlInstrument{
		Type:        XmlInstrumentInavType,
		Mic:         strings.ToUpper(mic),
		Isin:        strings.ToUpper(isin),
		Symbol:      strings.ToUpper(mnemonic),
		Mep:         strings.ToUpper(mep),
		File:        h5(mic, isin, mnemonic, XmlInstrumentInavType),
		Description: strPtr(""),
		Vendor:      strPtr("Euronext"),
		Inav: &XmlInav{
			Currency: "",
			Isin:     "",
			Mep:      "",
			Mic:      "",
			Name:     "",
			Symbol:   "",
			Vendor:   "",
			Target: []XmlTarget{
				{
					Isin:   "",
					Mep:    "",
					Mic:    "",
					Name:   "",
					Symbol: "",
					Vendor: "",
				},
			},
		},
	}
}

func NewXmlFund(mic, isin, mnemonic, mep string) *XmlInstrument {
	return &XmlInstrument{
		Type:        XmlInstrumentFundType,
		Mic:         strings.ToUpper(mic),
		Isin:        strings.ToUpper(isin),
		Symbol:      strings.ToUpper(mnemonic),
		Mep:         strings.ToUpper(mep),
		File:        h5(mic, isin, mnemonic, XmlInstrumentFundType),
		Description: strPtr(""),
		Vendor:      strPtr("Euronext"),
		Fund: &XmlFund{
			Cfi:         "",
			Currency:    "EUR",
			Issuer:      "",
			Shares:      "",
			TradingMode: "",
		},
	}
}

func NewXmlInstrument(instrumentType, mic, isin, mnemonic, mep string) *XmlInstrument {
	switch instrumentType {
	case XmlInstrumentStockType:
		return NewXmlStock(mic, isin, mnemonic, mep)
	case XmlInstrumentIndexType:
		return NewXmlIndex(mic, isin, mnemonic, mep)
	case XmlInstrumentEtvType:
		return NewXmlEtv(mic, isin, mnemonic, mep)
	case XmlInstrumentEtfType:
		return NewXmlEtf(mic, isin, mnemonic, mep)
	case XmlInstrumentInavType:
		return NewXmlInav(mic, isin, mnemonic, mep)
	case XmlInstrumentFundType:
		return NewXmlFund(mic, isin, mnemonic, mep)
	}

	return nil
}

func xmlInstrumentToXmlString(ins *XmlInstrument) string {
	// <instrument foundInSearch="true" mic="XPAR" isin="FR0013426004" symbol="CLA" name="CLARANOVA" type="stock" file="xpar/stock/CLA.h5:/XPAR_CLA_FR0013426004" description="" mep="PAR" vendor="Euronext">
	if ins.FoundInSearch == "" {
		return fmt.Sprintf(
			"  <instrument mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" type=\"%s\" file=\"%s\" description=\"%s\" mep=\"%s\" vendor=\"%s\">\n",
			ins.Mic, ins.Isin, ins.Symbol, ins.Name, ins.Type, ins.File, *ins.Description, ins.Mep, *ins.Vendor)
	} else {
		return fmt.Sprintf(
			"  <instrument foundInSearch=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" type=\"%s\" file=\"%s\" description=\"%s\" mep=\"%s\" vendor=\"%s\">\n",
			ins.FoundInSearch, ins.Mic, ins.Isin, ins.Symbol, ins.Name, ins.Type, ins.File, *ins.Description, ins.Mep, *ins.Vendor)
	}
}

const xmlInstrumentClosingTag = "  </instrument>\n"

func XmlStockToXmlString(ins *XmlInstrument) string {
	// <instrument foundInSearch="true" mic="XPAR" isin="FR0013426004" symbol="CLA" name="CLARANOVA" type="stock" file="xpar/stock/CLA.h5:/XPAR_CLA_FR0013426004" description="" mep="PAR" vendor="Euronext">
	//   <stock cfi="ESVUFN" compartment="B" tradingMode="continuous" currency="EUR" shares="39,442,878">
	//     <icb icb1="9000" icb2="9500" icb3="9530" icb4="9537" />
	//   </stock>
	// </instrument>
	s := xmlInstrumentToXmlString(ins)
	st := ins.Stock
	if st != nil {
		s += fmt.Sprintf(
			"    <stock cfi=\"%s\" compartment=\"%s\" tradingMode=\"%s\" currency=\"%s\" shares=\"%s\">\n",
			st.Cfi, st.Compartment, st.TradingMode, st.Currency, st.Shares)

		if st.Icb != nil {
			s += fmt.Sprintf(
				"      <icb icb1=\"%s\" icb2=\"%s\" icb3=\"%s\" icb4=\"%s\" />\n",
				st.Icb.Icb1, st.Icb.Icb2, st.Icb.Icb3, st.Icb.Icb4)
		}
		s += "    </stock>\n"
	}
	s += xmlInstrumentClosingTag
	return s
}

func XmlIndexToXmlString(ins *XmlInstrument) string {
	//  <instrument foundInSearch="true" mic="XPAR" isin="FR0003500008" symbol="PX1" name="CAC 40" type="index" file="xpar/index/PX1.h5:/XPAR_PX1_FR0003500008" description="The CAC 40 index is an index weighted by free-float market capitalization that measures the performance of 40 stocks selected among the top 100 market capitalisation and the most active stocks listed on Euronext Paris. Base value: 1000 at 31/12/1987." mep="PAR" vendor="Euronext">
	//    <index kind="price" family="CAC 40" calcFreq="15s" baseDate="1987-12-31" baseLevel="1000" baseCap="370437433957.70" baseCapCurrency="EUR" weighting="float market cap" capFactor="0.15" currency="EUR" />
	//  </instrument>
	//  <instrument foundInSearch="false" mic="XAMS" isin="QS0011017223" symbol="NLMED" name="AEX MEDIA" type="index" file="xams/index/icb/NLMED.h5:/XAMS_NLMED_QS0011017223" description="AEX ICB sector (level 3) index 5550 Media (Price)." mep="AMS" vendor="Euronext">
	//    <index kind="price" family="ICB" calcFreq="15s" baseDate="1998-12-31" baseLevel="1000" weighting="full market cap" capFactor="no" currency="EUR">
	//      <icb icb1="5000" icb2="5500" icb3="5550" />
	//    </index>
	//  </instrument>
	s := xmlInstrumentToXmlString(ins)
	id := ins.Index
	if id != nil {
		opt := ""
		if id.BaseCap != nil {
			opt += fmt.Sprintf(" baseCap=\"%s\"", *id.BaseCap)
		}
		if id.BaseCapCurrency != nil {
			opt += fmt.Sprintf(" baseCapCurrency=\"%s\"", *id.BaseCapCurrency)
		}
		if id.Icb != nil {
			s += fmt.Sprintf(
				"    <index kind=\"%s\" family=\"%s\" calcFreq=\"%s\" baseDate=\"%s\" baseLevel=\"%s\" weighting=\"%s\" capFactor=\"%s\" currency=\"%s\"%s>\n",
				id.Kind, id.Family, id.CalcFreq, id.BaseDate, id.BaseLevel, id.Weighting, id.CapFactor, id.Currency, opt)
			s += fmt.Sprintf(
				"      <icb icb1=\"%s\" icb2=\"%s\" icb3=\"%s\" icb4=\"%s\" />\n",
				id.Icb.Icb1, id.Icb.Icb2, id.Icb.Icb3, id.Icb.Icb4)
			s += "    </index>\n"
		} else {
			s += fmt.Sprintf(
				"    <index kind=\"%s\" family=\"%s\" calcFreq=\"%s\" baseDate=\"%s\" baseLevel=\"%s\" weighting=\"%s\" capFactor=\"%s\" currency=\"%s\"%s />\n",
				id.Kind, id.Family, id.CalcFreq, id.BaseDate, id.BaseLevel, id.Weighting, id.CapFactor, id.Currency, opt)
		}
	}
	s += xmlInstrumentClosingTag
	return s
}

func XmlEtvToXmlString(ins *XmlInstrument) string {
	//  <instrument foundInSearch="-typ()" mic="XAMS" isin="JE00B588CD74" symbol="SGBS" name="ETFS PHYS SW GLD" type="etv" file="xams/etv/SGBS.h5:/XAMS_SGBS_JE00B588CD74" description="" mep="AMS" vendor="Euronext">
	//     <etv cfi="DEMMR" tradingMode="continuous" allInFees="" expenseRatio="" dividendFrequency="" currency="EUR" issuer="ETFS METAL SECURITIES LTD" shares="4,338,963" />
	//  </instrument>
	s := xmlInstrumentToXmlString(ins)
	etv := ins.Etv
	if etv != nil {
		s += fmt.Sprintf(
			"    <etv cfi=\"%s\" tradingMode=\"%s\" allInFees=\"%s\" expenseRatio=\"%s\" dividendFrequency=\"%s\" currency=\"%s\" issuer=\"%s\" shares=\"%s\" />\n",
			etv.Cfi, etv.TradingMode, etv.AllInFees, etv.ExpenseRatio, etv.DividendFrequency, etv.Currency, etv.Issuer, etv.Shares)
	}
	s += xmlInstrumentClosingTag
	return s
}

func XmlEtfToXmlString(ins *XmlInstrument) string {
	//  <instrument foundInSearch="-typ()" mic="XAMS" isin="IE00BYPC1H27" symbol="CNYB" name="IS CHN BND USD ACC" type="etf" file="xams/etf/CNYB.h5:/XAMS_CNYB_IE00BYPC1H27" description="ISHARES CHINA CNY BOND UCITS ETF - (USD) DIST" mep="AMS" vendor="Euronext">
	//    <etf cfi="CEOMS" tradingMode="continuous" ter="0.35%" launchDate="29/07/2019" issuer="iShares IV plc." fraction="" dividendFrequency="distribution" indexFamily="" expositionType="physical" currency="USD">
	//       <inav vendor="Euronext" mep="" mic="" isin="NSCFR0ICNYB4" symbol="ICNYB" name="ISHARES CNYB INAV" />
	//       <underlying vendor="Euronext" mep="" mic="" isin="" symbol="" name="BBG Barclays CH Treasury + PBI" />
	//    </etf>
	//  </instrument>
	s := xmlInstrumentToXmlString(ins)
	etf := ins.Etf
	if etf != nil {
		s += fmt.Sprintf(
			"    <etf cfi=\"%s\" tradingMode=\"%s\" ter=\"%s\" launchDate=\"%s\" issuer=\"%s\" fraction=\"%s\" dividendFrequency=\"%s\" indexFamily=\"%s\" expositionType=\"%s\" currency=\"%s\">\n",
			etf.Cfi, etf.TradingMode, etf.Ter, etf.LaunchDate, etf.Issuer, etf.Fraction, etf.DividendFrequency, etf.IndexFamily, etf.ExpositionType, etf.Currency)

		if len(etf.Inav.Target) > 0 {
			s += fmt.Sprintf(
				"      <inav vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" currency=\"%s\">\n",
				etf.Inav.Vendor, etf.Inav.Mep, etf.Inav.Mic, etf.Inav.Isin, etf.Inav.Symbol, etf.Inav.Name, etf.Inav.Currency)
			for _, target := range etf.Inav.Target {
				s += fmt.Sprintf(
					"        <target vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" />\n",
					target.Vendor, target.Mep, target.Mic, target.Isin, target.Symbol, target.Name)
			}
			s += "      </inav>\n"
		} else {
			s += fmt.Sprintf(
				"      <inav vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" currency=\"%s\" />\n",
				etf.Inav.Vendor, etf.Inav.Mep, etf.Inav.Mic, etf.Inav.Isin, etf.Inav.Symbol, etf.Inav.Name, etf.Inav.Currency)
		}

		s += fmt.Sprintf(
			"      <underlying vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" />\n",
			etf.Underlying.Vendor, etf.Underlying.Mep, etf.Underlying.Mic, etf.Underlying.Isin, etf.Underlying.Symbol, etf.Underlying.Name)
		s += "    </etf>\n"
	}
	s += xmlInstrumentClosingTag
	return s
}

func XmlFundToXmlString(ins *XmlInstrument) string {
	//  <instrument foundInSearch="false" mic="XMLI" isin="FR0012767044" symbol="U1" name="ANAX BO EM 2020 U1" type="fund" file="xmli/fund/U1.h5:/XMLI_U1_FR0012767044" description="" mep="PAR" vendor="Euronext">
	//    <fund cfi="CIOS" tradingMode="" currency="EUR" issuer="" shares="159,203" />
	//  </instrument>
	s := xmlInstrumentToXmlString(ins)
	fund := ins.Fund
	if fund != nil {
		s += fmt.Sprintf(
			"    <fund cfi=\"%s\" tradingMode=\"%s\" currency=\"%s\" issuer=\"%s\" shares=\"%s\" />\n",
			fund.Cfi, fund.TradingMode, fund.Currency, fund.Issuer, fund.Shares)
	}
	s += xmlInstrumentClosingTag
	return s
}

func XmlInavToXmlString(ins *XmlInstrument) string {
	s := xmlInstrumentToXmlString(ins)
	inv := ins.Inav
	if inv != nil {
		if len(inv.Target) > 0 {
			s += fmt.Sprintf(
				"    <inav vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" currency=\"%s\">\n",
				inv.Vendor, inv.Mep, inv.Mic, inv.Isin, inv.Symbol, inv.Name, inv.Currency)
			for _, target := range inv.Target {
				s += fmt.Sprintf(
					"        <target vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" />\n",
					target.Vendor, target.Mep, target.Mic, target.Isin, target.Symbol, target.Name)
			}
			s += "      </inav>\n"
		} else {
			s += fmt.Sprintf(
				"      <inav vendor=\"%s\" mep=\"%s\" mic=\"%s\" isin=\"%s\" symbol=\"%s\" name=\"%s\" currency=\"%s\" />\n",
				inv.Vendor, inv.Mep, inv.Mic, inv.Isin, inv.Symbol, inv.Name, inv.Currency)
		}
		s += "    </inav>\n"
	}
	s += xmlInstrumentClosingTag
	return s
}

func XmlInstrumentToXmlString(ins *XmlInstrument) string {
	switch ins.Type {
	case XmlInstrumentStockType:
		return XmlStockToXmlString(ins)
	case XmlInstrumentIndexType:
		return XmlIndexToXmlString(ins)
	case XmlInstrumentEtvType:
		return XmlEtvToXmlString(ins)
	case XmlInstrumentEtfType:
		return XmlEtfToXmlString(ins)
	case XmlInstrumentInavType:
		return XmlInavToXmlString(ins)
	case XmlInstrumentFundType:
		return XmlFundToXmlString(ins)
	}
	return ""
}

func ReadXmlInstrumentsFile(fileName string) (*XmlInstruments, error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open xml file: %w", err)
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	xmlInstruments := XmlInstruments{}
	err = xml.Unmarshal(byteValue, &xmlInstruments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal xml file: %w", err)
	}

	return &xmlInstruments, nil
}

func ReadJsonInstrumentsFile(fileName string) (*XmlInstruments, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open json file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read json file: %w", err)
	}

	xmlInstruments := XmlInstruments{}
	err = json.Unmarshal(byteValue, &xmlInstruments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	return &xmlInstruments, nil
}

func WriteXmlInstrumentsFile(fileName string, xmlInstruments *XmlInstruments) error {
	xmlData, err := xml.MarshalIndent(xmlInstruments, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal xml: %w", err)
	}

	xmlFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer xmlFile.Close()

	_, err = xmlFile.Write(xmlData)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil
}

func WriteJsonInstrumentsFile(fileName string, xmlInstruments *XmlInstruments) error {
	jsonData, err := json.MarshalIndent(xmlInstruments, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal json: %w", err)
	}

	jsonFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil
}

/*
type XmlEuronextInstruments struct {
	XmlEuronextInstrument []struct {
		Cfi           *string `xml:"cfi,attr"`
		Description   *string `xml:"description,attr"`
		File          string  `xml:"file,attr"`
		FoundInSearch string  `xml:"foundInSearch,attr"`
		Isin          string  `xml:"isin,attr"`
		Mep           string  `xml:"mep,attr"`
		Mic           string  `xml:"mic,attr"`
		Name          string  `xml:"name,attr"`
		Notes         *string `xml:"notes,attr"`
		Symbol        string  `xml:"symbol,attr"`
		Tradingmode   *string `xml:"tradingmode,attr"`
		Type          string  `xml:"type,attr"`
		Vendor        *string `xml:"vendor,attr"`
		Etf           *struct {
			Cfi               string   `xml:"cfi,attr"`
			Currency          string   `xml:"currency,attr"`
			DividendFrequency string   `xml:"dividendFrequency,attr"`
			ExpositionType    string   `xml:"expositionType,attr"`
			Fraction          string   `xml:"fraction,attr"`
			IndexFamily       string   `xml:"indexFamily,attr"`
			Issuer            string   `xml:"issuer,attr"`
			LaunchDate        string   `xml:"launchDate,attr"`
			Mer               *float64 `xml:"mer,attr"`
			Ter               string   `xml:"ter,attr"`
			TradingMode       string   `xml:"tradingMode,attr"`
			Inav              struct {
				Isin   string `xml:"isin,attr"`
				Mep    string `xml:"mep,attr"`
				Mic    string `xml:"mic,attr"`
				Name   string `xml:"name,attr"`
				Symbol string `xml:"symbol,attr"`
				Vendor string `xml:"vendor,attr"`
			} `xml:"inav"`
			Underlying struct {
				Isin   string `xml:"isin,attr"`
				Mep    string `xml:"mep,attr"`
				Mic    string `xml:"mic,attr"`
				Name   string `xml:"name,attr"`
				Symbol string `xml:"symbol,attr"`
				Vendor string `xml:"vendor,attr"`
			} `xml:"underlying"`
		} `xml:"etf"`
		Etv *struct {
			AllInFees         string `xml:"allInFees,attr"`
			Cfi               string `xml:"cfi,attr"`
			Currency          string `xml:"currency,attr"`
			DividendFrequency string `xml:"dividendFrequency,attr"`
			ExpenseRatio      string `xml:"expenseRatio,attr"`
			Issuer            string `xml:"issuer,attr"`
			Shares            string `xml:"shares,attr"`
			TradingMode       string `xml:"tradingMode,attr"`
		} `xml:"etv"`
		Fund *struct {
			Cfi         string `xml:"cfi,attr"`
			Currency    string `xml:"currency,attr"`
			Issuer      string `xml:"issuer,attr"`
			Shares      string `xml:"shares,attr"`
			TradingMode string `xml:"tradingMode,attr"`
		} `xml:"fund"`
		Inav *struct {
			Currency string `xml:"currency,attr"`
			Target   []struct {
				Isin   string `xml:"isin,attr"`
				Mep    string `xml:"mep,attr"`
				Mic    string `xml:"mic,attr"`
				Name   string `xml:"name,attr"`
				Symbol string `xml:"symbol,attr"`
				Vendor string `xml:"vendor,attr"`
			} `xml:"target"`
		} `xml:"inav"`
		Index *struct {
			BaseCap         *string `xml:"baseCap,attr"`
			BaseCapCurrency *string `xml:"baseCapCurrency,attr"`
			BaseDate        string  `xml:"baseDate,attr"`
			BaseLevel       string  `xml:"baseLevel,attr"`
			CalcFreq        string  `xml:"calcFreq,attr"`
			CapFactor       string  `xml:"capFactor,attr"`
			Currency        string  `xml:"currency,attr"`
			Family          string  `xml:"family,attr"`
			Kind            string  `xml:"kind,attr"`
			Weighting       string  `xml:"weighting,attr"`
			Icb             *struct {
				Icb1 int  `xml:"icb1,attr"`
				Icb2 *int `xml:"icb2,attr"`
				Icb3 *int `xml:"icb3,attr"`
			} `xml:"icb"`
		} `xml:"index"`
		Stock *struct {
			Cfi         string `xml:"cfi,attr"`
			Compartment string `xml:"compartment,attr"`
			Currency    string `xml:"currency,attr"`
			Shares      string `xml:"shares,attr"`
			TradingMode string `xml:"tradingMode,attr"`
			Icb         struct {
				Icb1 string `xml:"icb1,attr"`
				Icb2 string `xml:"icb2,attr"`
				Icb3 string `xml:"icb3,attr"`
				Icb4 string `xml:"icb4,attr"`
			} `xml:"icb"`
		} `xml:"stock"`
	} `xml:"instrument"`
}

func ReadXmlInstrumentsFile(fileName string) (*XmlEuronextInstruments, error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open xml file: %w", err)
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	euronextInstruments := XmlEuronextInstruments{}
	err = xml.Unmarshal(byteValue, &euronextInstruments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal xml file: %w", err)
	}

	return &euronextInstruments, nil
}
*/
