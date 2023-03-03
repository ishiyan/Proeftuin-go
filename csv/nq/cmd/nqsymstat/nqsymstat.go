package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"nq/nasdaq"
)

type symbols struct {
	Updated string                `json:"updated"`
	Symbols []nasdaq.NasdaqSymbol `json:"symbols"`
}

type isinSedolSymbol struct {
	Mnemonic string `json:"mnemonic"`
	ISIN     string `json:"ISIN"`
	SEDOL    string `json:"SEDOL"`
}

func main() {
	t := time.Now().Format("2006-01-02_15-04-05")
	fmt.Println("=======================================")
	fmt.Println(t)
	fmt.Println("=======================================")

	symbolsPtr := flag.String("symbols", "nasdaq-all.json", "symbols filename.json")
	flag.Parse()

	syms, err := readSymbols(*symbolsPtr)
	if err != nil {
		panic(fmt.Sprintf("Cannot read symbols: %s", err))
	}

	micMap := map[string]int{}
	exchangeMap := map[string]int{}
	stockTypeMap := map[string]int{}
	assetClassMap := map[string]int{}
	industryMap := map[string]int{}
	sectorMap := map[string]int{}
	regionMap := map[string]int{}
	listedMap := map[bool]int{}
	ndx100Map := map[bool]int{}
	expenseRatioMap := map[string]int{}

	for _, s := range syms.Symbols {
		micMap[s.Mic] += 1
		exchangeMap[s.Exchange] += 1
		stockTypeMap[s.StockType] += 1
		assetClassMap[s.AssetClass] += 1
		industryMap[s.Industry] += 1
		sectorMap[s.Sector] += 1
		regionMap[s.Region] += 1
		listedMap[s.IsListed] += 1
		ndx100Map[s.IsNasdaq100] += 1
		expenseRatioMap[s.ExpenseRatio] += 1
	}

	fmt.Println("\nfile: " + *symbolsPtr)

	fmt.Println("\nmic:")
	for k, v := range micMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nexchange:")
	for k, v := range exchangeMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nstock type (stocks):")
	for k, v := range stockTypeMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nasset class:")
	for k, v := range assetClassMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nindustry (stocks):")
	for k, v := range industryMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nsector: (stocks)")
	for k, v := range sectorMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nregion (stocks):")
	for k, v := range regionMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nexpense ratio (etf):")
	for k, v := range expenseRatioMap {
		fmt.Printf("(%d) '%s'\n", v, k)
	}

	fmt.Println("\nis listed nasdaq (stocks):")
	for k, v := range listedMap {
		fmt.Printf("(%d) '%v'\n", v, k)
	}

	fmt.Println("\nis nasdaq 100 (stocks):")
	for k, v := range ndx100Map {
		fmt.Printf("(%d) '%v'\n", v, k)
	}

	isins, err := readIsinSedol(*symbolsPtr)
	if err != nil {
		fmt.Printf("Cannot read isin-sedol file: %s", err)
	}

	bothList := make([]isinSedolSymbol, 0)
	isinOnlyList := make([]isinSedolSymbol, 0)
	sedolOnlyList := make([]isinSedolSymbol, 0)
	bothEmptyList := make([]isinSedolSymbol, 0)

	const empty = ""
	for _, i := range isins {
		if i.ISIN == empty {
			if i.SEDOL == empty {
				bothEmptyList = append(bothEmptyList, i)
			} else {
				sedolOnlyList = append(sedolOnlyList, i)
			}
		} else {
			if i.SEDOL == empty {
				isinOnlyList = append(isinOnlyList, i)
			} else {
				bothList = append(bothList, i)
			}
		}
	}

	fmt.Printf("\n\nno ISIN, no SEDOL: %d\n", len(bothEmptyList))
	for _, v := range bothEmptyList {
		fmt.Printf("%s, ", v.Mnemonic)
	}

	fmt.Printf("\n\nno SEDOL: %d\n", len(isinOnlyList))
	for _, v := range isinOnlyList {
		fmt.Printf("%s, ", v.Mnemonic)
	}

	fmt.Printf("\n\nno ISIN: %d\n", len(sedolOnlyList))
	for _, v := range sedolOnlyList {
		fmt.Printf("%s, ", v.Mnemonic)
	}

	fmt.Printf("\n\nboth ISIN and SEDOL: %d\n", len(bothList))
}

func readSymbols(fileName string) (*symbols, error) {
	var s symbols

	f, err := os.Open(fileName)
	if err != nil {
		return &s, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&s)
	if err != nil {
		return &s, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	return &s, nil
}

func readIsinSedol(fileName string) ([]isinSedolSymbol, error) {
	s := make([]isinSedolSymbol, 0)

	fileName = strings.ReplaceAll(fileName, ".json", ".isin-sedol.json")
	f, err := os.Open(fileName)
	if err != nil {
		return s, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&s)
	if err != nil {
		return s, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	return s, nil
}
