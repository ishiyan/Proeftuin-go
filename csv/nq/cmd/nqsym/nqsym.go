package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"nq/nasdaq"
)

type symbols struct {
	Updated string                `json:"updated"`
	Symbols []nasdaq.NasdaqSymbol `json:"symbols"`
}

func main() {
	t := time.Now().Format("2006-01-02_15-04-05")
	fmt.Println("=======================================")
	fmt.Println(t)
	fmt.Println("=======================================")

	categoryPtr := flag.String("category", "stock", "category: [stock, stock-nasdaq, stock-nyse, stock-amex, etf]")
	flag.Parse()

	symbolsFileName := "nasdaq-stock.json"
	category := nasdaq.ExchangeAll
	switch *categoryPtr {
	case "stock-nasdaq":
		symbolsFileName = "nasdaq-stock-nasdaq.json"
		category = nasdaq.ExchangeNasdaq
	case "stock-nyse":
		symbolsFileName = "nasdaq-stock-nyse.json"
		category = nasdaq.ExchangeNyse
	case "stock-amex":
		symbolsFileName = "nasdaq-stock-amex.json"
		category = nasdaq.ExchangeAmex
	case "etf":
		symbolsFileName = "nasdaq-etf.json"
		category = ""
	}

	fmt.Println("reading " + symbolsFileName)
	symPrev, err := readSymbols(symbolsFileName)
	if err != nil {
		panic(fmt.Sprintf("cannot read symbols: %s", err))
	}

	var syms []nasdaq.NasdaqSymbol
	if *categoryPtr == "etf" {
		syms, err = nasdaq.RetrieveSymbolsEtf()
		if err != nil {
			panic(fmt.Sprintf("Cannot retrieve etf symbols: %s", err))
		}

	} else {
		syms, err = nasdaq.RetrieveSymbolsStock(category)
		if err != nil {
			panic(fmt.Sprintf("Cannot retrieve stock symbols: %s", err))
		}
	}

	symDict := map[string]nasdaq.NasdaqSymbol{}
	for _, s := range symPrev.Symbols {
		symDict[s.Mnemonic] = s
	}

	for _, s := range syms {
		symDict[s.Mnemonic] = s
	}

	suffix := "." + t
	err = os.Rename(symbolsFileName, symbolsFileName+suffix)
	if err != nil {
		panic(fmt.Sprintf("Cannot rename original JSON file: %s", err))
	}

	f, err := os.OpenFile(symbolsFileName, os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("cannot create new JSON file: %s", err))
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	sumNew := symbols{
		Updated: t,
		Symbols: syms,
	}

	err = enc.Encode(sumNew)
	if err != nil {
		panic(fmt.Sprintf("cannot encode new symbols: %s", err))
	}
	fmt.Println("finished: " + time.Now().Format("2006-01-02_15-04-05"))
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
