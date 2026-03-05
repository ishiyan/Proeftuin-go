package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
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

	fileName := symbolsFileName[:len(symbolsFileName)-4] + t + ".old.json"
	err = os.Rename(symbolsFileName, fileName)
	if err != nil {
		panic(fmt.Sprintf("Cannot rename original JSON file: %s", err))
	}

	fileName = symbolsFileName[:len(symbolsFileName)-4] + t + ".new.json"
	if err := writeSymbolsFile(fileName, t, syms); err != nil {
		panic(fmt.Sprintf("cannot write new JSON file: %s", err))
	}

	if err := writeSymbolsFile(symbolsFileName, t, syms); err != nil {
		panic(fmt.Sprintf("cannot write new JSON file: %s", err))
	}

	sort.Slice(symPrev.Symbols, func(i, j int) bool {
		return symPrev.Symbols[i].Mnemonic < symPrev.Symbols[j].Mnemonic
	})

	sort.Slice(syms, func(i, j int) bool {
		return syms[i].Mnemonic < syms[j].Mnemonic
	})

	fileName = symbolsFileName[:len(symbolsFileName)-4] + t + ".sorted.old.json"
	if err := writeSymbolsFile(fileName, t, symPrev.Symbols); err != nil {
		panic(fmt.Sprintf("cannot write sorted original JSON file: %s", err))
	}

	fileName = symbolsFileName[:len(symbolsFileName)-4] + t + ".sorted.new.json"
	if err := writeSymbolsFile(fileName, t, syms); err != nil {
		panic(fmt.Sprintf("cannot write sorted new JSON file: %s", err))
	}

	// Build map of previous symbols for lookup
	symPrevDict := map[string]nasdaq.NasdaqSymbol{}
	for _, s := range symPrev.Symbols {
		symPrevDict[s.Mnemonic] = s
	}

	// Build map of new symbols for lookup
	symsDict := map[string]nasdaq.NasdaqSymbol{}
	for _, s := range syms {
		symsDict[s.Mnemonic] = s
	}

	// Find introduced symbols (in syms but not in symPrev)
	var introduced []nasdaq.NasdaqSymbol
	for _, s := range syms {
		if _, exists := symPrevDict[s.Mnemonic]; !exists {
			introduced = append(introduced, s)
		}
	}

	// Find removed symbols (in symPrev but not in syms)
	var removed []nasdaq.NasdaqSymbol
	for _, s := range symPrev.Symbols {
		if _, exists := symsDict[s.Mnemonic]; !exists {
			removed = append(removed, s)
		}
	}

	fileName = symbolsFileName[:len(symbolsFileName)-4] + t + ".sorted.added.json"
	if err := writeSymbolsFile(fileName, t, introduced); err != nil {
		panic(fmt.Sprintf("cannot write sorted added JSON file: %s", err))
	}

	fileName = symbolsFileName[:len(symbolsFileName)-4] + t + ".sorted.removed.json"
	if err := writeSymbolsFile(fileName, t, removed); err != nil {
		panic(fmt.Sprintf("cannot write sorted removed JSON file: %s", err))
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

func writeSymbolsFile(fileName string, updated string, syms []nasdaq.NasdaqSymbol) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot create file '%s': %w", fileName, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	data := symbols{
		Updated: updated,
		Symbols: syms,
	}

	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("cannot encode to '%s': %w", fileName, err)
	}

	return nil
}
